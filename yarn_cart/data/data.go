package data

import (
	"context"
	"encoding/json"
	"log"
	"yarn_cart/models"

	"github.com/redis/go-redis/v9"
)

func redisConn() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "cache:6379",
		Password: "",
		DB: 0,
	})
	defer client.Close()

	return client
}

func PostRedis(cart *models.Cart) {
	client := redisConn()
	cartMarshaled, err := json.Marshal(cart.Items)
	if err != nil {
		log.Println(err)
		return
	}
	err = client.Set(context.Background(), cart.CartId.String(), cartMarshaled, 0)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetRedis(cartId string) *models.Cart {
	client := redisConn()
	var cart models.Cart
	val, err := client.Get(context.Background(), cartId).Result()
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(val, cart)
	if err != nil {
		log.Println(err)
	}
	return cart 
}