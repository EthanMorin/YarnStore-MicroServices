package data

import (
	"context"
	"encoding/json"
	"yarn_cart/models"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "cache:6379",
	Password: "",
	DB:       0,
})

func PostCart(cart *models.Cart) error {
	jsonItems, err := json.Marshal(cart.Items)
	if err != nil {
		return err // Return an error if JSON marshaling fails
	}
	err = client.Set(context.Background(), cart.CartId.String(), jsonItems, 0).Err()
	if err != nil {
		return err // Return an error if the Redis operation fails
	}
	return nil
}

func GetCart(cartId string) (*models.Cart, error) {
	var cart models.Cart
	id, err := uuid.Parse(cartId)
	if err != nil {
		return nil, err
	}
	cart.CartId = &id
	val, err := client.Get(context.Background(), cartId).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), &cart.Items)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func RemoveCart(cartId string)  error {
	err := client.Del(context.Background(), cartId).Err()
	if err != nil {
		return err
	}
	return nil
}
