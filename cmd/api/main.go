package main

import (
	"log"
	"test-grpc/service"
)

func main() {
	serverPort := 50051
	gwPort := 50052
	go func() {
		if err := service.StartServer(serverPort); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()
	if err := service.StartGateway(serverPort, gwPort); err != nil {
		log.Fatalf("failed to start gateway: %v", err)
	}
	return
}
