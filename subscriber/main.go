package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	daprCommon "github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"gorm.io/gorm"
	"log"
	"net/http"
	"simple-gorm-app/common"
)

const (
	pubSubName = "productpubsub"
	topicName  = "added-product"
)

var sub = &daprCommon.Subscription{
	PubsubName: pubSubName,
	Topic:      topicName,
	Route:      "/products",
}

func main() {

	db := common.InitializeDatabase()

	s := daprd.NewService(":6005")

	//Subscribe to a topic
	if err := s.AddTopicEventHandler(sub, eventHandler(db)); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
	if err := s.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("error listenning: %v", err)
	}
}

func eventHandler(db *gorm.DB) func(ctx context.Context, e *daprCommon.TopicEvent) (retry bool, err error) {
	return func(ctx context.Context, e *daprCommon.TopicEvent) (retry bool, err error) {

		productCode, ok := e.Data.(string)
		if !ok {
			return false, errors.New("invalid type")
		}

		product, err := common.GetProductByCode(db, productCode)
		if err != nil {
			return false, err
		}

		err = sendAuditNote(ctx, err, product)
		if err != nil {
			log.Printf("error sending audit note: %v", err)
			return false, err
		}

		log.Printf("Product: %s", product)

		return false, nil
	}
}

func sendAuditNote(ctx context.Context, err error, product *common.Product) error {

	const daprBindingService = "create-audit-product"

	data, err := json.Marshal(product)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:3500/v1.0/invoke/"+daprBindingService, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to invoke " + daprBindingService)
	}
	return nil
}
