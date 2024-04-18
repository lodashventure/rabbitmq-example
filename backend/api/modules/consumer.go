package modules

import (
	"log"
	"os"

	"github.com/lodashventure/rabbitmq-example/helpers"
	"github.com/lodashventure/rabbitmq-example/worker"
	"github.com/streadway/amqp"
)

func Consumer(serviceName string) {
	conn := helpers.NewRabbitmqConnection(os.Getenv("QUEUE_URL"))
	defer conn.RabbitmqClose()

	err := conn.RabbitmqConnect()
	if err != nil {
		Consumer(serviceName)
	}

	worker.CreateQueueAndExchange(conn)

	delivery, err := conn.RabbitmqConsume(serviceName, "")
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go conn.RabbitmqHandleConsumedDeliveries(serviceName, "", delivery, messageHandler)

	log.Println("Welcome to the server Consumer")

	<-forever
}

func messageHandler(c *helpers.Connection, deliveries <-chan amqp.Delivery) {

	for d := range deliveries {
		if os.Getenv("SERVICE_NAME") == os.Getenv("QUEUE_NAME_PUBLISH_MESSAGE_TO_CONSUMER_CAT") {
			worker.PublishMessageToConsumerCat(d.Body)
		} else if os.Getenv("SERVICE_NAME") == os.Getenv("QUEUE_NAME_PUBLISH_MESSAGE_TO_CONSUMER_DOG") {
			worker.PublishMessageToConsumerDog(d.Body)
		}

		d.Ack(false)
	}
}
