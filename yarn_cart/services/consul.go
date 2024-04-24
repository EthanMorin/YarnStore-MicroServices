package services

import (
	"fmt"
	"log"
	"os"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

func Register() {
	config := consulapi.DefaultConfig()
	config.Address = "consul:8500"
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println(err)
	}
	serviceId := "go-cart-service:" + getHostname()
	port, _ := strconv.Atoi(getPort()[1:len(getPort())])
	fmt.Printf("port:%v \n", port)
	address := getHostname()
	fmt.Printf("address:%v \n", address)
	tags := []string{
		"traefik.enable=true",
		fmt.Sprintf("traefik.http.routers.%s.rule=PathPrefix(`/cart`)", getHostname()),
	}
	registeration := &consulapi.AgentServiceRegistration{
		Tags:    tags,
		ID:      serviceId,
		Name:    "cart",
		Port:    port,
		Address: address,
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
