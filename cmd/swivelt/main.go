package main

import (
	"fmt"
	"os"

	"github.com/berlam/swivel-authenticator/pkg"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Call it with 'server_id' and 'user_pin'")
		os.Exit(1)
	}

	serverId := pkg.ServerId(args[0])
	pin := pkg.Pin(args[1])
	fmt.Println(pkg.Token(serverId, pin))
}
