package main

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func newServer(api *API) *gin.Engine {
	swagger, err := GetSwagger()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.Use(middleware.OapiRequestValidator(swagger))
	RegisterHandlers(router, api)
	return router
}

func main() {
	err := NewDB()
	if err != nil {
		panic(err)
	}
	Register()
	server := newServer(NewAPI())
	server.Run(":8080")
}