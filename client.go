package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Event struct {
	From    string
	Message interface{}
}

func SendEvent(endpoint string, message interface{}) {
	event := Event{From: "DOCKER", Message: message}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(event)

	_, err := http.Post(endpoint, "application/json; charset=utf-8", buffer)
	if err != nil {
		logPrintf("ERROR: Event sending failed: %s", err)
	}
}

func HealthCheck(healthcheck string) {
	for {
		<-time.After(5 * time.Second)
		_, err := http.Get(healthcheck)

		if err == nil {
			logPrintf("INFO: Swarmpit OK")
			break;
		}
	}
}
