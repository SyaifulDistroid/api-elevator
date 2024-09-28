package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"api-elevator/api/handler"
	"api-elevator/internal/repository"
	"api-elevator/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	rabbitMQRepo := repository.NewRabbitMQRepository()
	defer rabbitMQRepo.Close()

	elevatorService := service.NewElevatorService(rabbitMQRepo)
	go elevatorService.StartProcessing()

	router := gin.Default()
	elevatorHandler := handler.NewElevatorHandler(elevatorService)

	router.POST("/request-floor", elevatorHandler.RequestFloor)
	router.GET("/elevator-status", elevatorHandler.ElevatorStatus)

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %s", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down gracefully...")
	elevatorService.Stop()
	log.Println("Elevator service stopped.")
}
