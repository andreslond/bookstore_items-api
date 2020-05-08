package items

import (
	"errors"
	"github.com/andrestor2/bookstore_items-api/clients/elasticsearch"
	"github.com/andrestor2/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		returnErr := rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
		return &returnErr
	}
	i.Id = result.Id
	return nil
}
