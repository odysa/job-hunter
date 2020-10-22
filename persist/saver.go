package persist

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"job-hunter/engine"
	"log"
)

func Saver(collection string) (chan engine.Item, *mongo.Client, error) {
	in := make(chan engine.Item)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		itemCount := 1
		for {
			item := <-in
			log.Printf("Got Item #%d", itemCount)
			err := save(client, item, collection)
			if err != nil {
				log.Println(err)
			}
			itemCount++
		}
	}()
	return in, client, nil
}

func save(client *mongo.Client, item engine.Item, collection string) error {
	if client == nil {
		return fmt.Errorf("database client is empty")
	}
	col := client.Database("jobs").Collection(collection)
	_, err := col.InsertOne(context.TODO(), item)
	return err
}
