package api

import (
	"net/http"
	"yarn_catalog/models"

	"github.com/gin-gonic/gin"
)

type API struct{}

// DeleteCatalogProductId implements ServerInterface.
func (a *API) DeleteCatalogProductId(c *gin.Context, productId string) {
	panic("unimplemented")
}

// GetCatalog implements ServerInterface.
func (a *API) GetCatalog(c *gin.Context) {
	panic("unimplemented")
}

// GetCatalogProductId implements ServerInterface.
func (a *API) GetCatalogProductId(c *gin.Context, productId string) {
	panic("unimplemented")
}

// PatchCatalogProductId implements ServerInterface.
func (a *API) PatchCatalogProductId(c *gin.Context, productId string) {
	panic("unimplemented")
}

// PostCatalog implements ServerInterface.
func (a *API) PostCatalog(c *gin.Context) {
	var yarn models.PostCatalogJSONRequestBody
	if err := c.ShouldBindJSON(&yarn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return 
	}
	// TODO add mongo post
	c.JSON(http.StatusCreated, &yarn)
}

func NewAPI() *API {
	return &API{}
}
