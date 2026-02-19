package main

import (
	"fmt"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connStr := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connStr)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer conn.Close()
	fmt.Println("Started Peril server")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	sig := <-signals
	fmt.Printf("Signal %v caught: shutting down...", sig)
}
