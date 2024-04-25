package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	service *OrderService
}

func NewAPI(service *OrderService) *API {
	return &API{
		service: service,
	}
}

// GetOrdersCheck implements ServerInterface.
func (a *API) GetOrdersCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

// PostOrdersNewUserIdCartId implements ServerInterface.
func (a *API) PostOrdersNewUserIdCartId(c *gin.Context, userId string, cartId string) {
	var order Order
	order.UserId = &userId
	order.CartId = &cartId
	result, err := a.service.postOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	a.service.mq.CreateOrderNotification(userId)
	c.JSON(http.StatusCreated, result)
}
