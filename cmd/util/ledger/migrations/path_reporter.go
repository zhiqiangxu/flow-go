package migrations

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"os"
	"path"
	"runtime"
	"sync"
	"time"

	"github.com/rs/zerolog"

	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/interpreter"

	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/ledger"
)

// PathReporter iterates through registers getting the location and balance of all FlowVaults
type PathReporter struct {
	Log       zerolog.Logger
	OutputDir string
}

func (r *PathReporter) filename() string {
	return path.Join(r.OutputDir, fmt.Sprintf("path_report_%d.json", int32(time.Now().Unix())))
}

type pathDataPoint struct {
	// Path is the storage path
	Path string `json:"path"`
	// Address is the owner of the storage path
	Address string `json:"address"`
	// Type is the type at address
	Type string `json:"type"`
	// Size is the size of the register (without the key)
	Size int `json:"size"`
}

// Report creates a balance_report_*.json file that contains data on all FlowVaults in the state commitment.
// I recommend using gojq to browse through the data, because of the large uint64 numbers which jq won't be able to handle.
func (r *PathReporter) Report(payload []ledger.Payload) error {
	fn := r.filename()
	r.Log.Info().Msgf("Running FLOW path Reporter. Saving output to %s.", fn)

	f, err := os.Create(fn)
	if err != nil {
		return err
	}

	defer func() {
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}()

	writer := bufio.NewWriter(f)
	defer func() {
		err = writer.Flush()
		if err != nil {
			panic(err)
		}
	}()

	wg := &sync.WaitGroup{}
	resultsWG := &sync.WaitGroup{}
	jobs := make(chan ledger.Payload)
	resultsChan := make(chan pathDataPoint, 100)

	workerCount := runtime.NumCPU()

	results := make([]pathDataPoint, 0)

	resultsWG.Add(1)
	go func() {
		for point := range resultsChan {
			results = append(results, point)
		}
		resultsWG.Done()
	}()

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go r.pathReporterWorker(jobs, wg, resultsChan)
	}

	wg.Add(1)
	go func() {
		for _, p := range payload {
			jobs <- p
		}

		close(jobs)
		wg.Done()
	}()

	wg.Wait()

	//drain results chan
	close(resultsChan)
	resultsWG.Wait()

	tc, err := json.Marshal(struct {
		Data []pathDataPoint
	}{
		Data: results,
	})
	if err != nil {
		panic(err)
	}
	_, err = writer.Write(tc)
	if err != nil {
		panic(err)
	}

	return nil
}

func (r *PathReporter) pathReporterWorker(jobs chan ledger.Payload, wg *sync.WaitGroup, dataChan chan pathDataPoint) {
	for payload := range jobs {
		err := r.handlePayload(payload, dataChan)
		if err != nil {
			r.Log.Err(err).Msg("Error handling payload")
		}
	}

	wg.Done()
}

func (r *PathReporter) handlePayload(p ledger.Payload, dataChan chan pathDataPoint) error {
	id, err := keyToRegisterID(p.Key)
	if err != nil {
		return err
	}

	// Ignore known payload keys that are not Cadence values
	if state.IsFVMStateKey(id.Owner, id.Controller, id.Key) {
		return nil
	}

	value, version := interpreter.StripMagic(p.Value)

	err = storageMigrationV5DecMode.Valid(value)
	if err != nil {
		return nil
	}

	decodeFunction := interpreter.DecodeValue
	if version <= 4 {
		decodeFunction = interpreter.DecodeValueV4
	}

	// Decode the value
	owner := common.BytesToAddress([]byte(id.Owner))
	cPath := []string{id.Key}

	cValue, err := decodeFunction(value, &owner, cPath, version, nil)
	if err != nil {
		return fmt.Errorf(
			"failed to decode value: %w\n\nvalue:\n%s\n",
			err, hex.Dump(value),
		)
	}

	typename := ""

	if composite, isComposite := cValue.(*interpreter.CompositeValue); isComposite {
		typename = string(composite.TypeID())
	} else {
		typename = "non-composite"
	}

	dataChan <- pathDataPoint{
		Path:    id.Key,
		Address: owner.Hex(),
		Type:    typename,
		Size:    len(p.Value),
	}

	return nil
}
