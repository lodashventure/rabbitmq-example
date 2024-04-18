package worker

import (
	"os"

	"github.com/lodashventure/rabbitmq-example/helpers"
)

func CreateQueueAndExchange(conn *helpers.Connection) {
	/** =============================================
	 create PUBLISH_MESSAGE_TO_CONSUMER_DOG queues and exchange
	============================================= **/

	nameQueue := os.Getenv("QUEUE_NAME_PUBLISH_MESSAGE_TO_CONSUMER_DOG")       // queue name
	nameExchange := os.Getenv("EXCHANGE_NAME_PUBLISH_MESSAGE_TO_CONSUMER_DOG") // exchange name
	conn.RabbitmqAddExchangeDeclare(nameExchange, "fanout")                    // Declare new exchange with queue name and exchange name
	conn.RabbitmqAddQueueDeclare(nameQueue)                                    // Declare new queue
	conn.RabbitmqAddQueueBind(nameQueue, "", nameExchange)                     // Bind queue and exchange

	/** =============================================
	 create PUBLISH_MESSAGE_TO_CONSUMER_CAT queues and exchange
	============================================= **/

	nameQueue = os.Getenv("QUEUE_NAME_PUBLISH_MESSAGE_TO_CONSUMER_CAT")       // queue name
	nameExchange = os.Getenv("EXCHANGE_NAME_PUBLISH_MESSAGE_TO_CONSUMER_CAT") // exchange name
	conn.RabbitmqAddExchangeDeclare(nameExchange, "fanout")                   // Declare new exchange with queue name and exchange name
	conn.RabbitmqAddQueueDeclare(nameQueue)                                   // Declare new queue
	conn.RabbitmqAddQueueBind(nameQueue, "", nameExchange)                    // Bind queue and exchange

	/** ============================================= **/
}
