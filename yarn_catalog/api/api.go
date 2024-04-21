package api

import (
	"net/http"
	"yarn_catalog/data"
	"yarn_catalog/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type API struct{}

// GetCheck implements ServerInterface.
func (a *API) GetCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// DeleteCatalogProductId implements ServerInterface.
func (a *API) DeleteCatalogProductId(c *gin.Context, productId string) {
	objId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		panic(err)
	}
	err = data.DeleteYarn(objId)
	if err != nil {
		panic(err)
	}
	c.Status(http.StatusNoContent)
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
	objId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		panic(err)
	}
	result, err := data.GetYarn(objId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, &result)
}

// PatchCatalogProductId implements ServerInterface.
func (a *API) PatchCatalogProductId(c *gin.Context, productId string) {
	objId, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		panic(err)
	}
	var yarn models.PatchCatalogProductIdJSONBody
	if err = c.ShouldBindJSON(&yarn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	result, err := data.PatchYarn(objId, &yarn)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, &result)
}

// PostCatalog implements ServerInterface.
func (a *API) PostCatalog(c *gin.Context) {
	var yarn models.PostCatalogJSONBody
	if err := c.ShouldBindJSON(&yarn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	result, err := data.PostYarn(&yarn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, &result)
}

func NewAPI() *API {
	return &API{}
}
