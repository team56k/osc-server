package main

import (
	"github.com/hypebeast/go-osc/osc"
	"fmt"
	"time"
)

var clients = []*osc.Client{}

func main() {
	addr := "127.0.0.1:8765"
	server := &osc.Server{Addr: addr}

	server.Handle("/register", register)
	go server.ListenAndServe()

	for _ = range time.Tick(time.Second) {
		for _, c := range clients {
			msg := osc.NewMessage("/make-cube")
			msg.Append(true)
			c.Send(msg)
		}
	}
}

func register(msg *osc.Message) {
	if len(msg.Arguments) != 2 {
		fmt.Println("Wrong num args")
		return
	}

	ip, ok := msg.Arguments[0].(string)
	if !ok {
		fmt.Println("IP not a string")
		return
	}

	port, ok := msg.Arguments[1].(int32)
	if !ok {
		fmt.Println("Port not an int")
		return
	}

	clients = append(clients, osc.NewClient(ip, int(port)))
	fmt.Println(clients)
}