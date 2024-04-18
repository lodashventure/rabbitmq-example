package middlewares

import (
	"os"

	"github.com/lodashventure/rabbitmq-example/helpers"
)

func InitPublisher() *helpers.Connection {
	conn := helpers.NewRabbitmqConnection(os.Getenv("QUEUE_URL"))
	conn.RabbitmqConnect()
	return conn
}
