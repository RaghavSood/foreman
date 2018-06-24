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

func stopProcess(name string) error {
	result, err := client.StopProcess(name, true)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"result": result,
	}).Info("Stopped process")
	return nil
}

func startProcess(name string) error {
	result, err := client.StartProcess(name, true)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"result": result,
	}).Info("Stopped process")
	return nil
}

func restartProcess(name string) error {
	result, err := client.RestartProcess(name, true)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"result": result,
	}).Info("Stopped process")
	return nil
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

	id, err := client.ID()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error")
	}

	pid, err := client.PID()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error")
	}

	log.WithFields(log.Fields{
		"version": v,
		"api":     apiV,
		"id":      id,
		"pid":     pid,
	}).Info("Supervisor connected")
}
