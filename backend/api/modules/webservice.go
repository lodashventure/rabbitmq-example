package modules

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/lodashventure/rabbitmq-example/handlers"
	"github.com/lodashventure/rabbitmq-example/middlewares"
)

func Webservice() {
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024, // this is the default limit of 1GB
	})

	app.Use(
		cors.New(),
		logger.New(),
	)

	apiName := "api"
	version := "v1"

	api := app.Group(fmt.Sprintf("/%s/%s", apiName, version)) // /api/v1

	h := handlers.NewHandler(middlewares.InitPublisher())

	api.Post("/publish_message_to_customer_cat", h.PublishMessageToCunsumerCat)
	api.Post("/publish_message_to_customer_dog", h.PublishMessageToCunsumerDog)

	log.Fatal(app.Listen(os.Getenv("SRVC_SERVER_PORT")))
}
