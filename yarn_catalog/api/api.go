package api

import (
	"net/http"
	"yarn_catalog/data"

	"github.com/gin-gonic/gin"
)

type API struct{}

// DeleteCatalogProductId implements ServerInterface.
func (a *API) DeleteCatalogProductId(c *gin.Context, productId string) {
	panic("unimplemented")
}

// GetCatalog implements ServerInterface.
func (a *API) GetCatalog(c *gin.Context) {
	results, err := data.GetCatalog()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, &results)
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
	// var yarn models.PostCatalogJSONRequestBody
	// if err := c.ShouldBindJSON(&yarn); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	// 	return 
	// }
	// result, err := data.PostYarn(&yarn)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusCreated, &result)
}

func NewAPI() *API {
	return &API{}
}
