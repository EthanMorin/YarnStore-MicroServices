package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderService struct {
	mq    *RabbitMQ
	mongo *mongo.Collection
}

func NewOrderService(mq *RabbitMQ, mongo *mongo.Collection) *OrderService {
	return &OrderService{
		mq:    mq,
		mongo: mongo,
	}
}

func NewDB() (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		return nil, err
	}
	return client.Database("yarn_store").Collection("order"), err
}

func (c *OrderService) postOrder(obj *Order) (*mongo.InsertOneResult, error) {
	result, err := c.mongo.InsertOne(context.Background(), &obj)
	if err != nil {
		return nil, err
	}
	return result, nil
}
