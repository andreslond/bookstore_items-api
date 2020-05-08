package elasticsearch

import (
	"context"
	"fmt"
	"github.com/andrestor2/bookstore_items-api/logger"
	"github.com/olivere/elastic"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	SetClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetErrorLog(log.New()),
		//elastic.SetInfoLog(),
	)
	if err != nil {
		panic(err)
	}
	Client.SetClient(client)
}

func (c *esClient) SetClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger.Error(
			fmt.Sprintf("Error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil

}
