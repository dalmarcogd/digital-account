package rabbit

import (
	"fmt"
	"github.com/dalmarcogd/digital-account/transactions-persist/environments"
	"github.com/streadway/amqp"
)

func GetConnection() (*amqp.Connection, error) {
	env := environments.GetEnvironment()
	return amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v/%v", env.RabbitUsername, env.RabbitPassword, env.RabbitURL, env.RabbitPort, env.RabbitVHost))
}
