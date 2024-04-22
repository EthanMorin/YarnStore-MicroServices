package main

import (
	"yarn_cart/api"
	"yarn_cart/services"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func newServer(cartApi *api.API) *gin.Engine {
	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(router, cartApi)
	return router
}

func main() {
	services.Register()
	server := newServer(api.NewAPI())
	server.Run(":8080")
}
