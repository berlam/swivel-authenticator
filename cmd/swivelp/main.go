package main

import (
	"fmt"
	"os"

	"github.com/berlam/swivel-authenticator/domain"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Call it with 'server_id', 'username' and 'provision_code'")
		os.Exit(1)
	}

	serverId := domain.ServerId(args[0])
	username := domain.Username(args[1])
	provisionCode := domain.ProvisionCode(args[2])
	domain.Provision(serverId, username, provisionCode)
	fmt.Println("OK")
}
