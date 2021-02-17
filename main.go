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

	// list users
	// xs, err := rmqc.ListUsers()

	// fmt.Println(xs)

	// declares a queue

	// qs := rh.QueueSettings{}
	// qs.Durable = true
	// qs.AutoDelete = false

	// resp, err := rmqc.DeclareQueue("/", "c.queue", qs)

	// fmt.Println(resp)

	// lq, _ := rmqc.ListQueues()

	// fmt.Println(lq)

	uq := rh.UserSettings{}
	uq.Name = "Vinny"
	uq.Password = "VeryDifficult"
	uq.Tags = "meetingslave"

	nu, err := rmqc.PutUser("Vinvli", uq)

	fmt.Println(nu)
	fmt.Println(rmqc.ListUsers())

}
