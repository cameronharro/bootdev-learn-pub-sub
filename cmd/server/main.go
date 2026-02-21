package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, _ := createConn("amqp://guest:guest@localhost:5672/")
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = pubsub.PublishJSON(ch, string(routing.ExchangePerilDirect), string(routing.PauseKey), routing.PlayingState{IsPaused: true})
	waitKill()
}

func createConn(connStr string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(connStr)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	fmt.Println("Started Peril server")
	return conn, nil
}

func waitKill() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	sig := <-signals
	fmt.Printf("Signal %v caught: shutting down...", sig)
}
