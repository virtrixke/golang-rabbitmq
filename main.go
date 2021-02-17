package main

import (
	"fmt"
	"log"

	rh "github.com/michaelklishin/rabbit-hole/v2"
)

func main() {
	rmqc, err := rh.NewClient("http://rabbitmq:15672", "user", "password")

	if err != nil {
		log.Fatal("Cannot initialize: ", err)
	}

	resp, err := rmqc.Overview()

	if err != nil {
		log.Fatal("Cannot show overview: ", err)
	}

	fmt.Println(resp.RabbitMQVersion)

	xs, err := rmqc.ListNodes()

	if err != nil {
		log.Fatal("Failed to list nodes: ", err)
	}

	fmt.Println(xs)
}
