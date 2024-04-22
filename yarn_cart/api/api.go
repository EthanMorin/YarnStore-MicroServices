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
	panic("unimplemented")
}

// DeleteCartCartIdProductId implements ServerInterface.
func (a *API) DeleteCartCartIdProductId(c *gin.Context, cartId string, productId string) {
	panic("unimplemented")
}

// GetCartCartId implements ServerInterface.
func (a *API) GetCartCartId(c *gin.Context, cartId string) {
	panic("unimplemented")
}

// PatchCartCartIdProductId implements ServerInterface.
func (a *API) PatchCartCartIdProductId(c *gin.Context, cartId string, productId string) {
	panic("unimplemented")
}

// PostCartCartId implements ServerInterface.
func (a *API) PostCartCartId(c *gin.Context, cartId string) {
	panic("unimplemented")
}

// PostCartNew implements ServerInterface.
func (a *API) PostCartNew(c *gin.Context) {
	var cart models.Cart
	cartId := uuid.New()
	var yarn models.Yarn
	err := c.Bind(&yarn) // Make sure to pass a pointer to yarn
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart.CartId = &cartId
	if cart.Items == nil {
		cart.Items = &[]models.CartItem{} // Initialize cart.Items with the correct type
	}
	// Create a new CartItem instance and append it to cart.Items
	quantity := 1
	item := models.CartItem{
		Yarn:     &yarn,
		Quantity: &quantity, // Set the quantity as needed
	}
	*cart.Items = append(*cart.Items, item)
	data.PostRedis(&cart)
}

// GetCheck implements ServerInterface.
func (a *API) GetCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func NewAPI() *API {
	return &API{}
}
