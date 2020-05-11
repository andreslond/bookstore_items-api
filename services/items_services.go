package services

import (
	"github.com/andrestor2/bookstore_items-api/domain/items"
	"github.com/andrestor2/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(itemId string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: itemId}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}
