package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andrestor2/bookstore_items-api/clients/elasticsearch"
	"github.com/andrestor2/bookstore_utils-go/rest_errors"
	"strings"
)

const (
	indexItems = "items"
	objectType = "_doc"
)

func (i *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, objectType, i)
	if err != nil {
		returnErr := rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
		return &returnErr
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, objectType, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no items found with id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}
	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	i.Id = itemId
	return nil
}
