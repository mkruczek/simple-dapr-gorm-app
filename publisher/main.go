package main

import (
	"context"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"log"
	"simple-gorm-app/common"
	"time"
)

const (
	pubSubName = "productpubsub"
	topicName  = "added-product"
)

func main() {

	ctx := context.Background()
	db := common.InitializeDatabase()

	client, err := dapr.NewClient()
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 10; i++ {
		product := common.CreateProduct(db, fmt.Sprintf("Product_%d", i), 100)

		if err := client.PublishEvent(ctx, pubSubName, topicName, []byte(product.Code)); err != nil {
			log.Fatal(err)
		}
		time.Sleep(5000)
	}
}
