package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
func (a *API) GetOrderCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

// PostOrdersNewUserIdCartId implements ServerInterface.
func (a *API) PostOrderNewUserIdCartId(c *gin.Context, userId string, cartId string) {
	var order Order
	orderId := uuid.New()
	order.UserId = &userId
	order.CartId = &cartId
	order.OrderId = &orderId
	result, err := a.service.postOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	a.service.mq.CreateOrderNotification(userId)
	c.JSON(http.StatusCreated, result)
}
