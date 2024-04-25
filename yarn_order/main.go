package main

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func newServer(ordersApi *API) *gin.Engine {
	swagger, err := GetSwagger()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.Use(middleware.OapiRequestValidator(swagger))
	RegisterHandlers(router, ordersApi)
	return router
}

func main() {
	Register()
	mq, err := NewRabbitQueue()
	if err != nil {
		panic(err)
	}
	mongo, err := NewDB()
	if err != nil {
		panic(err)
	}
	service := NewOrderService(mq, mongo)
	server := newServer(NewAPI(service))
	server.Run(":8080")
}
