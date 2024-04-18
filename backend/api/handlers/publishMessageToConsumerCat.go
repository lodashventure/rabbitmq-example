package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/lodashventure/rabbitmq-example/models"
	"github.com/streadway/amqp"
)

func (h *Handler) PublishMessageToConsumerCat(c *fiber.Ctx) error {
	var requestBody models.MessageCat
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.Status(http.StatusOK).JSON(
			struct {
				Status  string
				Message string
			}{
				"faile",
				err.Error(),
			})
	}

	data, err := json.Marshal(requestBody)
	if err != nil {
		return c.Status(http.StatusOK).JSON(
			struct {
				Status  string
				Message string
			}{
				"faile",
				err.Error(),
			})
	}

	err = h.Rabbitmq.RabbitmqPublish(
		os.Getenv("EXCHANGE_NAME_PUBLISH_MESSAGE_TO_CONSUMER_CAT"), // exchange name
		"", // routing key
		amqp.Publishing{
			Timestamp:   time.Now(),
			ContentType: "application/json",
			Body:        data,
		})
	if err != nil {
		return c.Status(http.StatusOK).JSON(
			struct {
				Status  string
				Message string
			}{
				"faile",
				err.Error(),
			})
	}

	return c.Status(http.StatusOK).JSON(struct{ Status string }{"success"})
}
