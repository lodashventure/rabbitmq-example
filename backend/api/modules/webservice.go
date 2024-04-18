package modules

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	h := handlers.NewHandler(customMiddleware.InitPublisher())

	api.Get("")
}
