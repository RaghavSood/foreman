package main

import (
	"github.com/RaghavSood/foreman/supervisor"
	log "github.com/sirupsen/logrus"
)

var client supervisor.Client

func runClient() {
	log.Info("foreman")
	log.Info("client mode")

	var err error
	client, err = supervisor.NewClient("http://127.0.0.1:9001/RPC2")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error")
	}
	defer client.Close()

	logVersion()
}

func logVersion() {
	v, err := client.Version()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error")
	}

	apiV, err := client.APIVersion()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error")
	}

	log.WithFields(log.Fields{
		"version": v,
		"api":     apiV,
	}).Info("Supervisor connected")
}
