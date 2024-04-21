package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"yarn_catalog/api"
	"yarn_catalog/data"

	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	middleware "github.com/oapi-codegen/gin-middleware"
)

func register() {
	config := consulapi.DefaultConfig()
	config.Address = "consul:8500"
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}
	serviceId := "go-catalog-service"
	port, _ := strconv.Atoi(getPort()[1:len(getPort())])
	fmt.Printf("port:%v \n", port)
	address := getHostname()
	fmt.Printf("address:%v \n", address)

	registeration := &consulapi.AgentServiceRegistration{
		ID:      serviceId,
		Name:    "catalog",
		Port:    port,
		Address: address,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/check", address, port),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	regiErr := consul.Agent().ServiceRegister(registeration)
	if regiErr != nil {
		log.Panic(regiErr)
		log.Printf("Failed to register service: %s:%v ", address, port)
	} else {
		log.Printf("successfully register service: %s:%v", address, port)
	}
}

func getPort() (port string) {
	port = os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	port = ":" + port
	return
}

func getHostname() (hostname string) {
	hostname, _ = os.Hostname()
	return
}

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
	register()
	err := data.NewDB()
	if err != nil {
		panic(err)
	}
	server := newServer(api.NewAPI())
	server.Run(":8080")
}
