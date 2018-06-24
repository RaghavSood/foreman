package main

import (
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var clientMode bool
var serverMode bool

func init() {
	log.SetFormatter(&log.TextFormatter{})
	flag.BoolVar(&clientMode, "client", false, "starts foreman in client mode")
	flag.BoolVar(&serverMode, "server", false, "starts foreman in server mode")

	flag.Parse()

	if clientMode == serverMode {
		log.Fatal("You must specify only one of client and server")
	}
}

func main() {
	if clientMode {
		runClient()
	}

	// if serverMode {
	// server()
	// }
}
