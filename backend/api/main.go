package main

import (
	"errors"
	"os"

	"github.com/lodashventure/rabbitmq-example/modules"
)

func main() {
	serviceName := os.Getenv("SERVICE_NAME") // from environment in docker-compose.yml

	switch serviceName {
	case "webservice":
		modules.Webservice()
	case "consumer":
		modules.Consumer()
	default:
		panic(errors.New(serviceName + "is not supported."))
	}
}
