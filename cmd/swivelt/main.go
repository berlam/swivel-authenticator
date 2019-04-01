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

	if len(args) < 2 {
		fmt.Println("Call it with 'server_id', 'user_pin' and optionally '--no-verify' to disable certificate validation")
		os.Exit(1)
	}

	if len(args) == 3 && args[2] == "--no-verify" {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	serverId := pkg.ServerId(args[0])
	pin := pkg.Pin(args[1])
	token, err := pkg.Token(serverId, pin)
	if err != nil {
		log.Print(err)
		os.Exit(2)
	}

	fmt.Println(token)
	os.Exit(0)
}
