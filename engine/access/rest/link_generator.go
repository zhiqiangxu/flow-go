package rest

import (
	"github.com/gorilla/mux"
	"github.com/onflow/flow-go/model/flow"
)

type LinkGenerator interface {
	BlockLink(id flow.Identifier) (string, error)
	TransactionLink(id flow.Identifier) (string, error)
	TransactionResultLink(id flow.Identifier) (string, error)
}

type LinkGeneratorImpl struct {
	router *mux.Router
}

func NewLinkGeneratorImpl(router *mux.Router) *LinkGeneratorImpl {
	return &LinkGeneratorImpl{
		router: router,
	}
}

func (generator *LinkGeneratorImpl) BlockLink(id flow.Identifier) (string, error) {
	url, err := generator.router.Get(getBlocksByIDRoute).URLPath("id", id.String())
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func (generator *LinkGeneratorImpl) TransactionLink(id flow.Identifier) (string, error) {
	url, err := generator.router.Get(getTransactionByIDRoute).URLPath("id", id.String())
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func (generator *LinkGeneratorImpl) TransactionResultLink(id flow.Identifier) (string, error) {
	url, err := generator.router.Get(getTransactionResultRoute).URLPath("id", id.String())
	if err != nil {
		return "", err
	}
	return url.String(), nil
}
