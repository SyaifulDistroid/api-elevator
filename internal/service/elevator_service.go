package service

import (
	"api-elevator/internal/domain"
	"api-elevator/internal/repository"
	"sync"
	"time"
)

type ElevatorService struct {
	Elevator     *domain.Elevator
	Requests     chan domain.Request
	wg           sync.WaitGroup
	RabbitMQRepo *repository.RabbitMQRepository
}

func NewElevatorService(rmqRepo *repository.RabbitMQRepository) *ElevatorService {
	return &ElevatorService{
		Elevator:     &domain.Elevator{CurrentFloor: 0, Direction: domain.Idle},
		Requests:     make(chan domain.Request, 100),
		RabbitMQRepo: rmqRepo,
	}
}

func (es *ElevatorService) StartProcessing() {
	es.wg.Add(1)
	go es.processRequests()
	go es.RabbitMQRepo.ConsumeRequests(es.Requests)
}

func (es *ElevatorService) processRequests() {
	defer es.wg.Done()
	for req := range es.Requests {
		es.Elevator.AddRequest(req.Floor)
		time.Sleep(1 * time.Second)
		es.Elevator.MoveElevator()
	}
}

func (es *ElevatorService) Stop() {
	close(es.Requests)
	es.wg.Wait()
}
