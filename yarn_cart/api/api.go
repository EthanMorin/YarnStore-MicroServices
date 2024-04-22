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

// GetCheck implements ServerInterface.
func (a *API) GetCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// PatchCartCartIdProductId implements ServerInterface.
func (a *API) PatchCartCartIdProductId(c *gin.Context, cartId string, productId string) {
	panic("unimplemented")
}

// PostCartCartIdProductId implements ServerInterface.
func (a *API) PostCartCartIdProductId(c *gin.Context, cartId string, productId string) {
	panic("unimplemented")
}

// PostCartNew implements ServerInterface.
func (a *API) PostCartNew(c *gin.Context) {
	var cart models.Cart
	id := uuid.New()
	cart.CartId = &id
	
	data.PostRedis(cart)
	// panic("unimplemented")
}

func NewAPI() *API {
	return &API{}
}
