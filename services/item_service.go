package services

import (
	"github.com/nishant01/mybookstore_items-api/domain/items"
	"github.com/nishant01/mybookstore_utils-go/rest_errors"
	"net/http"
)

var (
	ItemService itemServiceInterface = &itemsService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	// return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not_implemented", nil)
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not_implemented", nil)
}
