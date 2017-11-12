package main

import (
	"fmt"
	"os"

	"github.com/berlam/swivel-authenticator/domain"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Call it with 'server_id' and 'user_pin'")
		os.Exit(1)
	}

	serverId := domain.ServerId(args[0])
	pin := domain.Pin(args[1])
	fmt.Println(domain.Token(serverId, pin))
}
