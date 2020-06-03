package items

import (
	"errors"
	"github.com/nishant01/mybookstore_items-api/clients/elasticsearch"
	"github.com/nishant01/mybookstore_utils-go/rest_errors"
)

const (
	indexItem = "item"
)
func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save items", errors.New("database_error"))

	}
	i.Id =	result.Id
	return nil
}
