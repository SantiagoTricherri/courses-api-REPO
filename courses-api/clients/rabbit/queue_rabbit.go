package rabbit

import (
	"courses-api/domain/courses"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitConfig struct {
	URI       string
	QueueName string
}

type Rabbit struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewRabbit(config RabbitConfig) Rabbit {
	connection, err := amqp.Dial(config.URI)
	if err != nil {
		log.Fatalf("error getting Rabbit connection: %v", err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("error creating Rabbit channel: %v", err)
	}
	queue, err := channel.QueueDeclare(config.QueueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("error declaring Rabbit queue: %v", err)
	}
	return Rabbit{
		connection: connection,
		channel:    channel,
		queue:      queue,
	}
}

func (r Rabbit) Publish(cursoNew courses.CursosNew) error {
	bytes, err := json.Marshal(cursoNew)
	if err != nil {
		return fmt.Errorf("error al serializar CursosNew: %w", err)
	}
	if err := r.channel.Publish(
		"",
		r.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		}); err != nil {
		return fmt.Errorf("error al publicar en Rabbit: %w", err)
	}
	return nil
}
