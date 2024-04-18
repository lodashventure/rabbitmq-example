package handlers

import "github.com/lodashventure/rabbitmq-example/helpers"

type Handler struct {
	Rabbitmq *helpers.Connection
}

func NewHandler(Rabbitmq *helpers.Connection) *Handler {
	return &Handler{
		Rabbitmq: Rabbitmq,
	}
}
