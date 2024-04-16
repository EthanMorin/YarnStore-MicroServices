package data

import (
	"context"
	"yarn_cart/models"
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

func yarnCollection() *mongo.Collection {
	return client.Database("yarn_store").Collection("yarn")
}

func PostYarn[T any](obj *T) (*mongo.InsertOneResult, error) {
	result, err := yarnCollection().InsertOne(context.Background(), &obj)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TODO: Convert to generics

func GetCatalog() (*[]models.Yarn, error) {
	var catalog []models.Yarn
	curser, err := yarnCollection().Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err := curser.All(context.Background(), &catalog); err != nil {
		return nil, err
	}
	return &catalog, nil
}

func GetYarn(objId primitive.ObjectID) (*models.Yarn, error) {
	if err := yarnCollection().FindOne(context.Background(), bson.M{"_id": objId}).Decode(&yarn); err != nil {
		return nil, err
	}
	return &yarn, nil
}

func PatchYarn[T](objId primitive.ObjectID, yarn models.PatchCartProductIdJSONBody) (*mongo.UpdateResult, error) {
	result, err := yarnCollection().UpdateOne(context.Background(), bson.M{"_id": objId}, bson.M{"$set": bson.M{"Available": &yarn.Amount}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteYarn(objId primitive.ObjectID) error {
	_, err := yarnCollection().DeleteOne(context.Background(), bson.M{"_id": objId})
	if err != nil {
		return err
	}
	return nil
}
