package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/hashicorp/consul/api"
)

func Register() {
	// Create a new Consul client
	config := api.DefaultConfig()
	config.Address = "consul:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	// Register the service with Consul
	port, err := strconv.Atoi(getPort()[1:len(getPort())])
	if err != nil {
		log.Fatalf("Failed to convert port to integer: %v", err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      "go-email-service:" + getHostname(),
		Name:    "email",
		Address: getHostname(),
		Port:    port,
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Failed to register service with Consul: %v", err)
	}

	// Deregister the service when the application shuts down
	defer func() {
		deregisterErr := client.Agent().ServiceDeregister("my-service-id")
		if deregisterErr != nil {
			log.Printf("Failed to deregister service with Consul: %v", deregisterErr)
		}
	}()

	// Keep the application running
	log.Println("Service registered with Consul. Press Ctrl+C to exit.")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
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
