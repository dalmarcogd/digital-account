package main

import (
	"github.com/dalmarcogd/digital-account/transactions-persist/brokers/events"
	"github.com/dalmarcogd/digital-account/transactions-persist/brokers/rabbit"
	"github.com/dalmarcogd/digital-account/transactions-persist/database"
	"github.com/dalmarcogd/digital-account/transactions-persist/utils"
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
		"transactions-persist", // name
		true,                   // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
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
			transactionEvent := events.NewTransactionCreateEvent("", "", 0, 0)
			failOnError(utils.NewJsonConverter().Decode(d.Body, transactionEvent), "Error when decode message")
			failOnError(database.CreateTransaction(transactionEvent.TransactionId, transactionEvent.AccountId, transactionEvent.OperationTypeId, transactionEvent.Amount), "Erro when save transaction")
			log.Printf("Transaction saved %s", transactionEvent.AccountId)

		}
	}()

	log.Printf("Consuming messages...")
	<-forever
}
