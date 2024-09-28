package handler

import (
	"api-elevator/internal/domain"
	"api-elevator/internal/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ElevatorHandler struct {
	Service *service.ElevatorService
}

func NewElevatorHandler(s *service.ElevatorService) *ElevatorHandler {
	return &ElevatorHandler{Service: s}
}

func (h *ElevatorHandler) RequestFloor(c *gin.Context) {
	floor, _ := strconv.Atoi(c.PostForm("floor"))
	err := h.Service.RabbitMQRepo.SendRequest(domain.Request{Floor: floor})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send request"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "request added", "floor": floor})
}

func (h *ElevatorHandler) ElevatorStatus(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	for {
		status := h.Service.Elevator
		_, err := c.Writer.Write([]byte("data: " + fmt.Sprintf("CurrentFloor: %v, Direction: %v, Requests: %v", status.CurrentFloor, status.Direction, status.Requests) + "\n\n"))
		if err != nil {
			return
		}
		c.Writer.Flush()
		time.Sleep(1 * time.Second)
	}
}
