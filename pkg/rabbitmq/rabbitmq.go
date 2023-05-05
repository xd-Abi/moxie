package rabbitmq

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/xd-Abi/moxie/pkg/logging"
)

const (
	AuthExchangeKey    = "auth.events"
	ProfileQueueKey    = "profile"
	UserSignUpEventKey = "user.signup"
)

type RabbitMQConnection struct {
	Log                *logging.Log
	InternalConnection *amqp.Connection
	InternalChannel    *amqp.Channel
}

type Event struct {
	Key     string
	Payload interface{}
}

type UserSignUpEventPayload struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewConnection(url string, log *logging.Log) *RabbitMQConnection {
	connection, err := amqp.Dial(url)

	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ: %v", err)
		return nil
	}

	log.Info("Connected to RabbitMQ")
	channel, err := connection.Channel()

	if err != nil {
		log.Fatal("Failed to create RabbitMQ Channel: %v", err)
		return nil
	}

	return &RabbitMQConnection{
		Log:                log,
		InternalConnection: connection,
		InternalChannel:    channel,
	}
}

func (rmq *RabbitMQConnection) DeclareQueue(name string) amqp.Queue {
	q, err := rmq.InternalChannel.QueueDeclare(
		name,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		rmq.Log.Error("Could not create queue: %v", err)
	}

	return q
}

func (rmq *RabbitMQConnection) DeclareExchange(name string) {
	err := rmq.InternalChannel.ExchangeDeclare(
		name,
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		rmq.Log.Error("Failed to declare exchange: %v", err)
	}
}

func (rmq *RabbitMQConnection) Bind(name string, key string, exchange string) {
	err := rmq.InternalChannel.QueueBind(
		name,
		key,
		exchange,
		false,
		nil,
	)
	if err != nil {
		rmq.Log.Error("Failed to bind queue: %s", err)
	}
}

func (rmq *RabbitMQConnection) Publish(exchange string, event *Event) {
	payload, err := json.Marshal(event.Payload)
	if err != nil {
		rmq.Log.Error("Failed to marshal event payload: %s", err)
	}

	err = rmq.InternalChannel.Publish(
		exchange,
		event.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
			Headers:     amqp.Table{"x-event-type": event.Key},
		},
	)
	if err != nil {
		rmq.Log.Error("Failed to publish event: %v", err)
	} else {
		rmq.Log.Info("Published event successfully: %v", event)
	}
}

func (rmq *RabbitMQConnection) Consume(queueName string, handler func(*Event) error) {
	rmq.DeclareQueue(queueName)

	msgs, err := rmq.InternalChannel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		rmq.Log.Error("Failed to register consumer: %v", err)
	}

	go func() {
		for msg := range msgs {
			var payload map[string]interface{}
			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				rmq.Log.Error("Failed to unmarshal event payload: %v", err)
				continue
			}

			eventType, ok := msg.Headers["x-event-type"].(string)
			if !ok {
				rmq.Log.Error("Mssing or invalid x-event-type header")
			}

			event := &Event{
				Key:     eventType,
				Payload: payload,
			}

			if err := handler(event); err != nil {
				rmq.Log.Error("Error handling event: %v", err)
			}
		}
	}()
}

func NewSignUpEvent(payload UserSignUpEventPayload) *Event {
	return &Event{
		Key:     UserSignUpEventKey,
		Payload: payload,
	}
}
