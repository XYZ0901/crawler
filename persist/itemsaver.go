package persist

import (
	"context"
	"crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var (
	client *elastic.Client
	err    error
	// _index 后期可以写在config里面
	_index = "dating_profile"
)

func init() {
	client, err = elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
}

func ItemSaver() chan model.Item {
	out := make(chan model.Item)
	go func() {
		for {
			item := <-out
			err := save(item)
			if err != nil {
				log.Println("Item Saver: error "+"saving item ", item, ":", err)
			}
		}
	}()
	return out
}

func save(item model.Item) error {
	_, err := client.Index().
		Index(_index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item.Profile).
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
