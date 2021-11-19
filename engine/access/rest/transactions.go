package rest

import (
	"context"
	"fmt"
	"github.com/onflow/flow-go/engine/access/rest/middleware"
	"github.com/onflow/flow-go/model/flow"
	"net/http"

	"github.com/rs/zerolog"

	"github.com/onflow/flow-go/access"
	"github.com/onflow/flow-go/engine/access/rest/generated"
)

const ExpandableFieldProposalKey = "proposal_key"
const ExpandableFieldAuthorizers = "authorizers"
const ExpandableFieldPayloadSignatures = "payload_signatures"
const ExpandableFieldEnvelopeSignatures = "envelope_signatures"
const ExpandableFieldResult = "result"
const ExpandableFieldEvents = "events"

// getTransactionByID gets a transaction by requested ID.
func getTransactionByID(
	w http.ResponseWriter,
	r *http.Request,
	vars map[string]string,
	backend access.API,
	linkGenerator LinkGenerator,
	logger zerolog.Logger,
) (interface{}, StatusError) {

	id, err := toID(vars["id"])
	if err != nil {
		return nil, NewBadRequestError("invalid ID", err)
	}

	expandFields, _ := middleware.GetFieldsToExpand(r)
	selectFields, _ := middleware.GetFieldsToSelect(r)

	transactionFactory := newTransactionResponseFactory(expandFields, selectFields)

	transaction, err := transactionFactory.transactionResponse(r.Context(), id, backend, linkGenerator)
	if err != nil {
		return nil, NewBadRequestError("transaction fetching error", err)
	}

	return transaction, nil
}

type transactionResponseFactory struct {
	expandProposalKey        bool
	expandAuthorizers        bool
	expandPayloadSignatures  bool
	expandEnvelopeSignatures bool
	expandResult             bool
	selectFields             map[string]bool
}

func newTransactionResponseFactory(expandFields map[string]bool, selectFields map[string]bool) *transactionResponseFactory {
	txFactory := new(transactionResponseFactory)
	txFactory.expandProposalKey = expandFields[ExpandableFieldProposalKey]
	txFactory.expandAuthorizers = expandFields[ExpandableFieldAuthorizers]
	txFactory.expandPayloadSignatures = expandFields[ExpandableFieldPayloadSignatures]
	txFactory.expandEnvelopeSignatures = expandFields[ExpandableFieldEnvelopeSignatures]
	txFactory.expandResult = expandFields[ExpandableFieldResult]
	txFactory.selectFields = selectFields
	return txFactory
}

func (txRespFactory *transactionResponseFactory) transactionResponse(ctx context.Context, id flow.Identifier, backend access.API, linkGenerator LinkGenerator) (*generated.Transaction, StatusError) {
	var responseTransaction = new(generated.Transaction)

	tx, err := backend.GetTransaction(ctx, id)
	if err != nil {
		return nil, NewBadRequestError(err.Error(), err)
	}

	if txRespFactory.expandEnvelopeSignatures {
		responseTransaction.EnvelopeSignatures = transactionSignatureResponse(tx.EnvelopeSignatures)
	}
	if txRespFactory.expandAuthorizers {
		var auths []string
		for _, auth := range tx.Authorizers {
			auths = append(auths, auth.String())
		}

		responseTransaction.Authorizers = auths
	}
	if txRespFactory.expandProposalKey {
		responseTransaction.ProposalKey = proposalKeyResponse(&tx.ProposalKey)
	}
	if txRespFactory.expandPayloadSignatures {
		responseTransaction.PayloadSignatures = transactionSignatureResponse(tx.PayloadSignatures)
	}
	if txRespFactory.expandResult {
		txr, err := backend.GetTransactionResult(ctx, id)
		if err != nil {
			return nil, NewBadRequestError(err.Error(), err)
		}

		responseTransaction.Result = transactionResultResponse(txr)
	}

	transactionLink, err := linkGenerator.TransactionLink(id)
	if err != nil {
		msg := fmt.Sprintf("failed to generate respose for transaction ID %s", id.String())
		return nil, NewRestError(http.StatusInternalServerError, msg, err)
	}

	responseTransaction.Links = new(generated.Links)
	responseTransaction.Links.Self = transactionLink

	return responseTransaction, nil
}

type transactionResultResponseFactory struct {
	expandEvent  bool
	selectFields map[string]bool
}

func newTransactionResultResponseFactory(expandFields map[string]bool, selectFields map[string]bool) *transactionResponseFactory {
	txFactory := new(transactionResponseFactory)
	txFactory.expandProposalKey = expandFields[ExpandableFieldEvents]
	txFactory.selectFields = selectFields
	return txFactory
}

func (txResultRespFactory *transactionResultResponseFactory) transactionResultResponse(ctx context.Context, id flow.Identifier, backend access.API, linkGenerator LinkGenerator) (*generated.TransactionResult, StatusError) {
	var responseResultTransaction = new(generated.TransactionResult)

	txresult, err := backend.GetTransactionResult(ctx, id)
	if err != nil {
		return nil, NewBadRequestError(err.Error(), err)
	}

	responseResultTransaction = transactionResultResponse(txresult)

	if !txResultRespFactory.expandEvent {
		responseResultTransaction.Events = nil
	}

	transactionResultLink, err := linkGenerator.TransactionResultLink(id)
	if err != nil {
		msg := fmt.Sprintf("failed to generate respose for transaction ID %s", id.String())
		return nil, NewRestError(http.StatusInternalServerError, msg, err)
	}

	responseResultTransaction.Links = new(generated.Links)
	responseResultTransaction.Links.Self = transactionResultLink

	return responseResultTransaction, nil
}

func getTransactionResultByID(
	w http.ResponseWriter,
	r *http.Request,
	vars map[string]string,
	backend access.API,
	linkGenerator LinkGenerator,
	logger zerolog.Logger,
) (interface{}, StatusError) {
	id, err := toID(vars["id"])
	if err != nil {
		return nil, NewBadRequestError("invalid ID", err)
	}

	expandFields, _ := middleware.GetFieldsToExpand(r)
	selectFields, _ := middleware.GetFieldsToSelect(r)

	transactionFactory := newTransactionResponseFactory(expandFields, selectFields)

	transaction, err := transactionFactory.transactionResponse(r.Context(), id, backend, linkGenerator)
	if err != nil {
		return nil, NewBadRequestError("transaction fetching error", err)
	}

	return transaction, nil
}

// createTransaction creates a new transaction from provided payload.
func createTransaction(
	w http.ResponseWriter,
	r *http.Request,
	vars map[string]string,
	backend access.API,
	generator LinkGenerator,
	logger zerolog.Logger,
) (interface{}, StatusError) {

	var txBody generated.TransactionsBody
	err := jsonDecode(r.Body, &txBody)
	if err != nil {
		return nil, NewBadRequestError("invalid transaction request", err)
	}

	tx, err := toTransaction(&txBody)
	if err != nil {
		return nil, NewBadRequestError("invalid transaction request", err)
	}

	err = backend.SendTransaction(r.Context(), &tx)
	if err != nil {
		return nil, NewBadRequestError("failed to send transaction", err)
	}

	return transactionResponse(&tx), nil
}
