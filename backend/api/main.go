package main

import (
	"errors"
	"os"

	"github.com/lodashventure/rabbitmq-example/modules"
)

func main() {
	serviceName := os.Getenv("SERVICE_NAME") // from environment in docker-compose.yml

	switch serviceName {
	case "webservice-api":
		modules.Webservice()
	case "publish-message-to-consumer-dog", "publish-message-to-consumer-cat":
		modules.Consumer(serviceName)
	default:
		panic(errors.New(serviceName + " is not supported."))
	}
}
