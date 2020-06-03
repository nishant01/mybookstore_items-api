package services

import (
	"github.com/nishant01/mybookstore_items-api/domain/items"
	"github.com/nishant01/mybookstore_items-api/domain/queries"
	"github.com/nishant01/mybookstore_utils-go/rest_errors"
)

var (
	ItemService itemServiceInterface = &itemsService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	// return nil, rest_errors.NewRestError("implement me", http.StatusNotImplemented, "not_implemented", nil)
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}
