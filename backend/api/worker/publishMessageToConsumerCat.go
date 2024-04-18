package worker

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lodashventure/rabbitmq-example/models"
)

func PublishMessageToConsumerCat(body []byte) {
	var data models.MessageCat

	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("faile: ", err.Error())
		return
	}

	message := fmt.Sprintf("[CAT] name: %s, age: %d, message: %s", data.Name, data.Age, data.Message)
	log.Println(message)
}
