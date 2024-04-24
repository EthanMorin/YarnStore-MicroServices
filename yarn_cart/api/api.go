package api

import (
	// "log"
	"net/http"
	"yarn_cart/data"
	"yarn_cart/models"

	// "yarn_cart/data"
	// "yarn_cart/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "github.com/google/uuid"
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
	cart, err := data.GetRedis(cartId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
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
	var items models.CartItem
	quantity := 1
	cartId := uuid.New()
	items.Quantity = &quantity
	cart.CartId = &cartId
	err := c.Bind(&items.Yarn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.PostRedis(&cart)
	c.JSON(http.StatusCreated, cart)
}

// GetCheck implements ServerInterface.
func (a *API) GetCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

// DeleteCartCartId implements ServerInterface.
// func (a *API) DeleteCartCartId(c *gin.Context, cartId string) {
// 	err := data.RemoveCart(cartId)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// 	c.Status(http.StatusNoContent)
// }

// // DeleteCartCartIdProductId implements ServerInterface.
// func (a *API) DeleteCartCartIdProductId(c *gin.Context, cartId string, productId string) {
// 	err := data.RemoveItem(cartId, productId)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.Status(http.StatusNoContent)
// }

// // GetCartCartId implements ServerInterface.
// func (a *API) GetCartCartId(c *gin.Context, cartId string) {
// 	cart, err := data.GetRedis(cartId)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, cart)
// }

// // PatchCartCartIdProductId implements ServerInterface.
// func (a *API) PatchCartCartIdProductId(c *gin.Context, cartId string, productId string) {
// 	var quantity models.PatchCartCartIdProductIdJSONRequestBody
// 	err := c.Bind(&quantity)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	err = data.PatchQuantity(cartId, productId, &quantity)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.Status(http.StatusNoContent)
// }

// // PostCartCartId implements ServerInterface.
// func (a *API) PostCartCartId(c *gin.Context, cartId string) {
// 	var item models.PostCartCartIdJSONRequestBody
// 	err := c.Bind(&item)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	cart, err := data.GetRedis(cartId)
// 	if err != nil {
// 		log.Println(err)
// 		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error())}
// 	}
// *cart.Items = append(*cart.Items, item)
// 	data.PostRedis(cart)
// }

// // PostCartNew implements ServerInterface.
// func (a *API) PostCartNew(c *gin.Context) {
// 	var yarn models.Yarn
// 	err := c.Bind(&yarn)
// 	var cart models.Cart
// 	cartId := uuid.New()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	cart.CartId = &cartId
// 	quantity := 1
// 	item := models.CartItem{
// 		Yarn:     &yarn,
// 		Quantity: &quantity,
// 	}
// 	*cart.Items = append(*cart.Items, item)
// 	data.PostRedis(&cart)
// }

func NewAPI() *API {
	return &API{}
}
