package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nishant01/mybookstore_items-api/clients/elasticsearch"
	"github.com/nishant01/mybookstore_items-api/domain/queries"
	"github.com/nishant01/mybookstore_utils-go/rest_errors"
	"strings"
)

const (
	indexItem = "item"
	typeItem = "_doc"
)
func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItem, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save items", errors.New("database_error"))

	}
	i.Id =	result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItem, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database_error"))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying parse database response"), errors.New("database_error"))
	}

	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying parse database response"), errors.New("database_error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItem, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to search document"), errors.New("database_error"))
	}
	fmt.Println(result)

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to parse response"), errors.New("database_error"))
		}
		item.Id = hit.Id
		items[index] = item
	}
	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("No item find")
	}
	return items, nil
}

