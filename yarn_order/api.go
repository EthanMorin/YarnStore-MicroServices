package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct{}

// GetOrdersCheck implements ServerInterface.
func (a *API) GetOrdersCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// PostOrdersNewUserIdCartId implements ServerInterface.
func (a *API) PostOrdersNewUserIdCartId(c *gin.Context, userId string, cartId string) {
	// performs an http request to the cart service to get the cart
	// performs an http request to the user service to get the user
	// creates a new order
	// sends a message to rabbitmq to notify email service to send email to user
	// returns the orderId
	panic("not implemented")
}

func NewAPI() *API {
	return &API{}
}
