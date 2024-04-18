package helpers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type queueDeclares struct {
	name string
}

type exchangeDeclares struct {
	name string
	kind string
}

type queueBinds struct {
	name     string
	key      string
	exchange string
}

// Connection is the connection created
type Connection struct {
	url                  string
	err                  chan error
	errReconnect         chan error
	timeSleepDuration    time.Duration
	conn                 *amqp.Connection
	channel              *amqp.Channel
	queueDeclares        []queueDeclares
	exchangeDeclares     []exchangeDeclares
	queueBinds           []queueBinds
	queueDeclaresConsume []queueDeclares
	serviceName          string
}

// NewConnection RabbitMQ returns the new connection object
func NewRabbitmqConnection(url string) *Connection {
	return &Connection{
		url:               url,
		err:               make(chan error),
		errReconnect:      make(chan error),
		timeSleepDuration: 2 * time.Second,
	}
}

func (c *Connection) RabbitmqConnect() error {
	var err error
	c.conn, err = amqp.Dial(c.url)
	if err != nil {
		time.Sleep(c.timeSleepDuration)
		log.Printf("could not connect to rabbitmq server: %s", err)
		return fmt.Errorf("error in creating rabbitmq connection with %s : %s", c.url, err.Error())
	}

	go func() {
		<-c.conn.NotifyClose(make(chan *amqp.Error)) //Listen to NotifyClose
		c.err <- errors.New("connection RabbitMQ service Closed")
	}()

	c.channel, err = c.conn.Channel()
	if err != nil {
		time.Sleep(c.timeSleepDuration)
		log.Printf("could not open channel to rabbitmq: %s", err)
		return fmt.Errorf("error in rabbitmq open channel: %s", err)
	}

	log.Println("Connected to RabbitMQ service")

	return nil
}

func (c *Connection) queueDeclare(name string) error {
	if _, err := c.channel.QueueDeclare(
		name,  //name
		true,  //durable
		false, //autoDelete
		false, //exclusive
		false, //noWait
		nil,   //arguments
	); err != nil {
		log.Printf("failed to declare a queue: %s", err)
		return fmt.Errorf("error in declaring the queue %s", err)
	}

	return nil
}

func (c *Connection) RabbitmqAddQueueDeclare(name string) error {
	if err := c.queueDeclare(name); err != nil {
		return err
	}

	c.queueDeclares = append(c.queueDeclares, queueDeclares{name: name})

	return nil
}

func (c *Connection) RabbitmqAddQueueDeclareConsume(name string) error {
	if err := c.queueDeclare(name); err != nil {
		return err
	}

	c.queueDeclaresConsume = append(c.queueDeclaresConsume, queueDeclares{name: name})

	return nil
}

func (c *Connection) exchangeDeclare(name, kind string) error {
	if err := c.channel.ExchangeDeclare(
		name,  // name
		kind,  // type
		true,  // durable
		true,  // auto-deleted
		false, // internal
		false, // noWait
		nil,   // arguments
	); err != nil {
		log.Printf("failed to declare an exchange: %s", err)
		return fmt.Errorf("error in declaring the exchange %s", err)
	}

	return nil
}

func (c *Connection) RabbitmqAddExchangeDeclare(name, kind string) error {
	if err := c.exchangeDeclare(name, kind); err != nil {
		return err
	}

	c.exchangeDeclares = append(c.exchangeDeclares, exchangeDeclares{name: name, kind: kind})

	return nil
}

func (c *Connection) queueBind(name, key, exchange string) error {
	if err := c.channel.QueueBind(
		name,     // queue name
		key,      // routing key
		exchange, // exchange
		false,
		nil,
	); err != nil {
		log.Printf("failed to queue bind: %s", err)
		return fmt.Errorf("error in queue bind %s", err)
	}

	return nil
}

func (c *Connection) RabbitmqAddQueueBind(name, key, exchange string) error {
	if err := c.queueBind(name, key, exchange); err != nil {
		return err
	}

	c.queueBinds = append(c.queueBinds, queueBinds{name: name, key: key, exchange: exchange})

	return nil
}

// Reconnect reconnects the connection
func (c *Connection) reconnect() error {
	log.Printf("could not connect to rabbitmq server")

	if err := c.RabbitmqConnect(); err != nil {
		return err
	}

	for _, q := range c.queueDeclares {
		if err := c.queueDeclare(q.name); err != nil {
			return err
		}
	}

	for _, q := range c.queueDeclaresConsume {
		if err := c.queueDeclare(q.name); err != nil {
			return err
		}
	}

	for _, ex := range c.exchangeDeclares {
		if err := c.exchangeDeclare(ex.name, ex.kind); err != nil {
			return err
		}
	}

	for _, qb := range c.queueBinds {
		if err := c.queueBind(qb.name, qb.key, qb.exchange); err != nil {
			return err
		}
	}

	return nil
}

// Publish publishes a request to the amqp queue
func (c *Connection) RabbitmqPublish(exchange, key string, message amqp.Publishing) error {
	select { //non blocking channel - if there is no error will go to default where we do nothing
	case err := <-c.err:
		if err != nil {
			go func() {
				for {
					err := c.reconnect()
					if err == nil {
						<-c.errReconnect
						break
					} else {
						c.errReconnect <- err
					}
				}
			}()
		}
	default:
	}

	select {
	case err := <-c.errReconnect:
		if err != nil {
			log.Printf("failed to publishing: %s", err)
			return fmt.Errorf("error in publishing: %s", err)
		}
	default:
	}

	var err error
	isFirst := true
	retry := 3
	for i := 0; i < retry; i++ {
		if isFirst || err != nil {
			err = c.channel.Publish(
				exchange, //exchange
				key,      //key
				false,    //mandatory
				false,    //immediate
				message,  //msg
			)
			isFirst = false
		} else {
			break
		}
	}

	if err != nil {
		defer c.reconnect()
		log.Printf("failed to publishing: %s", err)
		return fmt.Errorf("error in publishing: %s", err)
	}

	return nil
}

// Consume consumes the messages from the queues and passes it as map of chan of amqp.Delivery
func (c *Connection) RabbitmqConsume(queue, consumer string) (<-chan amqp.Delivery, error) {
	delivery, err := c.channel.Consume(
		queue,    //queue
		consumer, //consumer
		false,    //autoAck
		false,    //exclusive
		false,    //noLocal
		false,    //noWait
		nil,      //args
	)
	if err != nil {
		log.Printf("failed to consume: %s", err)
		return nil, err
	}

	return delivery, nil
}

// HandleConsumedDeliveries handles the consumed deliveries from the queues. Should be called only for a consumer connection
func (c *Connection) RabbitmqHandleConsumedDeliveries(queue, consumer string, delivery <-chan amqp.Delivery, fn func(*Connection, <-chan amqp.Delivery)) {
	for {
		go fn(c, delivery)

		if err := <-c.err; err != nil {
			for {
				err := c.reconnect()
				if err == nil {
					break
				}
			}

			delivery, err = c.RabbitmqConsume(queue, consumer)
			if err != nil {
				log.Panic("failed to consume")
			}
		}
	}
}

func (c *Connection) RabbitmqClose() {
	c.conn.Close()
}
