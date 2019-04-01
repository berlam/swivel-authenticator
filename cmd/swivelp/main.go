package main

import (
	"fmt"
	"os"

	"crypto/tls"
	"log"
	"net/http"
	"swivel-authenticator/pkg"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Call it with 'server_id', 'username', 'provision_code' and optionally '--no-verify' to disable certificate validation")
		os.Exit(1)
	}

	if len(args) == 4 && args[3] == "--no-verify" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	serverId := pkg.ServerId(args[0])
	username := pkg.Username(args[1])
	provisionCode := pkg.ProvisionCode(args[2])
	err := pkg.Provision(serverId, username, provisionCode)
	if err != nil {
		log.Print(err)
		os.Exit(2)
	}

	fmt.Println("OK")
	os.Exit(0)
}
