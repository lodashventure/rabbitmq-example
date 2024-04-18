package modules

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func Consumer() {
	conn := helpers.NewRabbitmqConnection(os.Getenv("QUEUE_URL"))
	defer conn.RabbitmqClose()

	err := conn.RabbitmqConnect()
	if err != nil {
		Consumer()
	}

	nameQueue := "upload-queue"
	nameExchange := "uploader-exchange"

	conn.RabbitmqAddExchangeDeclare(nameExchange, "fanout")
	conn.RabbitmqAddQueueDeclare(nameQueue)
	conn.RabbitmqAddQueueBind(nameQueue, "", nameExchange)

	delivery, err := conn.RabbitmqConsume(nameQueue, "")
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go conn.RabbitmqHandleConsumedDeliveries(nameQueue, "", delivery, messageHandler)

	log.Println("Welcome to the server Consumer")

	<-forever
}

func messageHandler(c *helpers.Connection, deliveries <-chan amqp.Delivery) {
	config := helpers.PrepareConfiguration()

	for d := range deliveries {

		d.Ack(false)
	}
}
