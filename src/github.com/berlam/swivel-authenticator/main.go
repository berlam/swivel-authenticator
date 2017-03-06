package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Call it with 'otc' or 'provision'")
	}

	cmd := args[0]
	serverId := ServerId(args[1])
	switch cmd {
	case "otc":
		pin := Pin(args[2])
		fmt.Println(Token(serverId, pin))
		break
	case "provision":
		username := Username(args[2])
		provisionCode := ProvisionCode(args[3])
		Provision(serverId, username, provisionCode)
		fmt.Println("OK")
		break
	}
}
