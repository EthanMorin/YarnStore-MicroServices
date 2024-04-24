package data

import (
	"context"
	"encoding/json"
	"log"
	"yarn_cart/models"

	"github.com/redis/go-redis/v9"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "cache:6379",
	Password: "",
	DB:       0,
})

func PostRedis(cart *models.Cart) error {
	// Marshal the cart into a JSON string
	val, err := json.Marshal(cart)
	if err != nil {
		log.Println(err)
		return err // Return an error if JSON marshaling fails
	}
	// Save the cart to Redis
	err = client.Set(context.Background(), cart.CartId.String(), val, 0).Err()
	if err != nil {
		log.Println(err)
		return err // Return an error if the Redis operation fails
	}
	return nil
}

func GetRedis(cartId string) (*models.Cart, error) {
	var cart models.Cart
	cartJson, err := client.Get(context.Background(), cartId).Result()
	if err != nil {
		log.Println(err)
		return nil, err // Return an error if the Redis operation fails
	}
	err = json.Unmarshal([]byte(cartJson), &cart.Items)
	if err != nil {
		log.Println(err)
		return nil, err // Return an error if JSON unmarshaling fails
	}
	return &cart, nil
}

func PatchQuantity(cartId string, productId string, newQuantity *models.PatchCartCartIdProductIdJSONRequestBody) error {
	// Retrieve the cart from Redis
	val, err := client.Get(context.Background(), cartId).Result()
	if err != nil {
		log.Println("Error retrieving cart:", err)
		return err
	}

	// Unmarshal the JSON string into a models.Cart struct
	var cart models.Cart
	err = json.Unmarshal([]byte(val), &cart)
	if err != nil {
		log.Println("Error unmarshaling cart:", err)
		return err
	}

	// Find the item in the cart
	found := false
	for i, item := range *cart.Items {
		if *item.Yarn.ProductId == productId {
			// Update the quantity
			*item.Quantity = *newQuantity.Quantity
			// Update the item in the cart
			(*cart.Items)[i] = item
			found = true
			break
		}
	}

	if !found {
		log.Println("Item not found in cart")
		return nil // Consider returning an error if the item is not found
	}

	// Save the updated cart back to Redis
	err = PostRedis(&cart)
	if err != nil {
		log.Println("Error saving updated cart:", err)
		return err
	}

	return nil
}

func RemoveCart(cartId string) error {
	err := client.Del(context.Background(), cartId).Err()
	if err != nil {
		log.Println("Error deleting cart:", err)
		return err
	}
	return nil
}

func RemoveItem(cartId string, productId string) error {
	// Retrieve the cart from Redis
	val, err := client.Get(context.Background(), cartId).Result()
	if err != nil {
		log.Println("Error retrieving cart:", err)
		return err
	}

	// Unmarshal the JSON string into a models.Cart struct
	var cart models.Cart
	err = json.Unmarshal([]byte(val), &cart)
	if err != nil {
		log.Println("Error unmarshaling cart:", err)
		return err
	}

	// Find the item in the cart and remove it
	items := *cart.Items
	for i, item := range items {
		if *item.Yarn.ProductId == productId {
			// Remove the item from the slice
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
	cart.Items = &items

	// Save the updated cart back to Redis
	err = PostRedis(&cart)
	if err != nil {
		log.Println("Error saving updated cart:", err)
		return err
	}

	return nil
}
