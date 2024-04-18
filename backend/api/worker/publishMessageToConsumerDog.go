package worker

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lodashventure/rabbitmq-example/models"
)

func PublishMessageToConsumerDog(body []byte) {
	var data models.MessageDog

	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("faile: ", err.Error())
		return
	}

	message := fmt.Sprintf("[DOG] breed: %s, message: %s", data.Breed, data.Message)
	log.Println(message)
}
