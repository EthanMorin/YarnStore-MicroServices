package main

import (
	"yarn_catalog/api"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func newServer(catalogApi *api.API) *gin.Engine {
	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(router, catalogApi)
	return router
}

func main() {
	server := newServer(api.NewAPI())
	server.Run(":8080")
}