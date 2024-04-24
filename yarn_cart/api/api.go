package api

import (
	"net/http"
	"yarn_cart/data"
	"yarn_cart/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type API struct{}

// DeleteCartCartId implements ServerInterface.
func (a *API) DeleteCartCartId(c *gin.Context, cartId string) {
	if err := data.RemoveCart(cartId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// DeleteCartCartIdProductId implements ServerInterface.
func (a *API) DeleteCartCartIdProductId(c *gin.Context, cartId string, productId string) {
	cart, err := data.GetCart(cartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, item := range *cart.Items {
		if *item.Yarn.ProductId == productId {
			*cart.Items = append((*cart.Items)[:i], (*cart.Items)[i+1:]...)
			break
		}
	}
	if err := data.PostCart(cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetCartCartId implements ServerInterface.
func (a *API) GetCartCartId(c *gin.Context, cartId string) {
	cart, err := data.GetCart(cartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

// PostCartCartId implements ServerInterface.
func (a *API) PostCartCartId(c *gin.Context, cartId string) {
	var items models.CartItems
	if err := c.BindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart, err := data.GetCart(cartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, newItem := range items {
		found := false
		for _, existingItem := range *cart.Items {
			if *existingItem.Yarn.ProductId == *newItem.Yarn.ProductId {
				*existingItem.Quantity += *newItem.Quantity
				found = true
				break
			}
		}
		if !found {
			*cart.Items = append(*cart.Items, newItem)
		}
	}
	if err := data.PostCart(cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

// PostCartNew implements ServerInterface.
func (a *API) PostCartNew(c *gin.Context) {
	var cart models.Cart
	var items models.PostCartNewJSONRequestBody
	if err := c.BindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := uuid.New()
	cart.CartId = &id
	cart.Items = &items
	if err := data.PostCart(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.PostCart(&cart)
	c.JSON(http.StatusOK, cart)
}

// GetCheck implements ServerInterface.
func (a *API) GetCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func NewAPI() *API {
	return &API{}
}
