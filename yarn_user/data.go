package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func NewDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		return err
	}
	return nil
}

func userCollection() *mongo.Collection {
	return client.Database("yarn_store").Collection("user")
}

func postUser(obj *PostUserJSONRequestBody) (*mongo.InsertOneResult, error) {
	result, err := userCollection().InsertOne(context.Background(), &obj)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getUser(objId primitive.ObjectID) (*User, error) {
	var user User
	if err := userCollection().FindOne(context.Background(), bson.M{"_id": objId}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(objId primitive.ObjectID) error {
	_, err := userCollection().DeleteOne(context.Background(), bson.M{"_id": objId})
	if err != nil {
		return err
	}
	return nil
}