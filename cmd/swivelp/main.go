package main

import (
	"fmt"
	"os"

	"github.com/berlam/swivel-authenticator/pkg"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Call it with 'server_id', 'username' and 'provision_code'")
		os.Exit(1)
	}

	serverId := pkg.ServerId(args[0])
	username := pkg.Username(args[1])
	provisionCode := pkg.ProvisionCode(args[2])
	pkg.Provision(serverId, username, provisionCode)
	fmt.Println("OK")
}
