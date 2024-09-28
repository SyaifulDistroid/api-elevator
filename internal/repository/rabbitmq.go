package repository

import (
	"api-elevator/internal/domain"
	"fmt"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

type RabbitMQRepository struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewRabbitMQRepository() *RabbitMQRepository {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	queue, err := channel.QueueDeclare(
		"elevator_requests", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}
	return &RabbitMQRepository{Conn: conn, Channel: channel, Queue: queue}
}

func (r *RabbitMQRepository) SendRequest(req domain.Request) error {
	body := fmt.Sprintf("%d", req.Floor)
	err := r.Channel.Publish("", r.Queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	return err
}

func (r *RabbitMQRepository) ConsumeRequests(requests chan domain.Request) {
	msgs, err := r.Channel.Consume(
		r.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	for msg := range msgs {
		floor, _ := strconv.Atoi(string(msg.Body))
		requests <- domain.Request{Floor: floor}
	}
}

func (r *RabbitMQRepository) Close() {
	r.Channel.Close()
	r.Conn.Close()
}
