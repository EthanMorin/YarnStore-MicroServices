package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type API struct{}

// PostUserLogin implements ServerInterface.
func (a *API) PostUserLogin(c *gin.Context) {
	var body PostUserLoginJSONRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := getUserByEmailAndPassword(body.Email, body.Password)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return
	}
	token, err := createToken(result.Email)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetUserCheck implements ServerInterface.
func (a *API) GetUserCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

// DeleteUserUserId implements ServerInterface.
func (a *API) DeleteUserUserId(c *gin.Context, userId string) {
	c.Header("Content-Type", "application/json")
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	tokenString = tokenString[len("Bearer "):]
	err := verifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = DeleteUser(objId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetUserUserId implements ServerInterface.
func (a *API) GetUserUserId(c *gin.Context, userId string) {
	objId, err := primitive.ObjectIDFromHex(userId)
	result, err := getUser(objId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// PostUser implements ServerInterface.
func (a *API) PostUser(c *gin.Context) {
	var body PostUserJSONRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := postUser(&body)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func NewAPI() *API {
	return &API{}
}
