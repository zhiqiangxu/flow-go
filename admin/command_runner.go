package admin

import (
	"context"
	"errors"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/onflow/flow-go/admin/admin"
	"github.com/rs/zerolog"
)

const (
	CommandRunnerMaxQueueLength = 128
	CommandRunnerNumWorkers     = 1
)

type CommandRunner struct {
	mu         sync.RWMutex
	handlers   map[string]CommandHandler
	validators map[string]CommandValidator
	commandQ   chan *CommandRequest
	address    string
	logger     zerolog.Logger
}

type CommandHandler func(ctx context.Context, data map[string]interface{}) error
type CommandValidator func(data map[string]interface{}) error

func NewCommandRunner(logger zerolog.Logger, address string) *CommandRunner {
	return &CommandRunner{
		handlers:   make(map[string]CommandHandler),
		validators: make(map[string]CommandValidator),
		commandQ:   make(chan *CommandRequest, CommandRunnerMaxQueueLength),
		address:    address,
		logger:     logger.With().Str("admin", "command_runner").Logger(),
	}
}

func (r *CommandRunner) RegisterHandler(command string, handler CommandHandler) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.handlers[command]; ok {
		return false
	}
	r.handlers[command] = handler
	return true
}

func (r *CommandRunner) UnregisterHandler(command string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.handlers, command)
}

func (r *CommandRunner) RegisterValidator(command string, validator CommandValidator) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.validators[command]; ok {
		return false
	}
	r.validators[command] = validator
	return true
}

func (r *CommandRunner) UnregisterValidator(command string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.validators, command)
}

func (r *CommandRunner) getHandler(command string) CommandHandler {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.handlers[command]
}

func (r *CommandRunner) getValidator(command string) CommandValidator {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.validators[command]
}

func (r *CommandRunner) Start(ctx context.Context) {
	for i := 0; i < CommandRunnerNumWorkers; i++ {
		go r.processLoop(ctx)
	}
	go r.runAdminServer(ctx)
}

func (r *CommandRunner) runAdminServer(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	default:
	}

	r.logger.Info().Msg("admin server starting up")

	listener, err := net.Listen("tcp", r.address)
	if err != nil {
		r.logger.Fatal().Err(err).Msg("failed to listen on admin server address")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, NewAdminServer(r.commandQ))
	go grpcServer.Serve(listener)

	<-ctx.Done()
	r.logger.Info().Msg("admin server shutting down")

	grpcServer.Stop()
	close(r.commandQ)
}

func (r *CommandRunner) processLoop(ctx context.Context) {
	defer func() {
		for command := range r.commandQ {
			close(command.responseChan)
		}
	}()

	for {
		select {
		case command := <-r.commandQ:
			r.logger.Info().Str("command", command.command).Msg("received new command")

			var err error

			if validator := r.getValidator(command.command); validator != nil {
				if validationErr := validator(command.data); validationErr != nil {
					err = status.Error(codes.InvalidArgument, validationErr.Error())
					goto sendResponse
				}
			}

			if handler := r.getHandler(command.command); handler != nil {
				// TODO: we can probably merge the command context with the worker context
				// using something like: https://github.com/teivah/onecontext
				if handleErr := handler(command.ctx, command.data); handleErr != nil {
					if errors.Is(handleErr, context.Canceled) {
						err = status.Error(codes.Canceled, "client canceled")
					} else if errors.Is(handleErr, context.DeadlineExceeded) {
						err = status.Error(codes.DeadlineExceeded, "request timed out")
					} else {
						err = status.Error(codes.Unknown, handleErr.Error())
					}
				}
			} else {
				err = status.Error(codes.Unimplemented, "invalid command")
			}

		sendResponse:
			command.responseChan <- &CommandResponse{err}
			close(command.responseChan)
		case <-ctx.Done():
			r.logger.Info().Msg("process loop shutting down")
			return
		}
	}

}
