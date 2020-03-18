package rabbit

import (
	"fmt"
	"github.com/dalmarcogd/digital-account/transactions/brokers/events"
	"github.com/dalmarcogd/digital-account/transactions/environments"
	"github.com/dalmarcogd/digital-account/transactions/errors"
	"github.com/dalmarcogd/digital-account/transactions/utils"
	"github.com/streadway/amqp"
)

func getConnection() (*amqp.Connection, error) {
	env := environments.GetEnvironment()
	return amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v/%v", env.RabbitUsername, env.RabbitPassword, env.RabbitURL, env.RabbitPort, env.RabbitVHost))
}

type rabbit struct{}

func NewRabbit() *rabbit {
	return &rabbit{}
}

func (r rabbit) Publish(event events.Event) error {
	connection, err := getConnection()
	if err != nil {
		return errors.NewRabbitConnectionError(err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	exchangeName := fmt.Sprintf("%v.master", event.GetName())

	err = channel.ExchangeDeclare(exchangeName, amqp.ExchangeFanout, true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, err := utils.NewJsonConverter().Encode(event)
	if err != nil {
		return err
	}

	return channel.Publish(exchangeName, event.GetChannel(), false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
}
