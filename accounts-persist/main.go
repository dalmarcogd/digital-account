package main

import (
	"github.com/dalmarcogd/digital-account/accounts-persist/brokers/events"
	"github.com/dalmarcogd/digital-account/accounts-persist/brokers/rabbit"
	"github.com/dalmarcogd/digital-account/accounts-persist/database"
	"github.com/dalmarcogd/digital-account/accounts-persist/utils"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	database.Migrate()
	defer database.CloseConnection()

	connection, err := rabbit.GetConnection()
	failOnError(err, "Error when get connection")
	defer connection.Close()

	channel, err := connection.Channel()
	failOnError(err, "Error when get channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"accounts-persist", // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Error when declare a queue")

	msgs, err := channel.Consume(queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	failOnError(err, "Error when create consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			accountEvent := events.NewAccountCreateEvent("", "")
			failOnError(utils.NewJsonConverter().Decode(d.Body, accountEvent), "Error when decode message")
			failOnError(database.CreateAccount(accountEvent.AccountId, accountEvent.DocumentNumber), "Erro when save account")
			log.Printf("Account saved %s", accountEvent.AccountId)

		}
	}()

	log.Printf("Consuming messages...")
	<-forever
}
