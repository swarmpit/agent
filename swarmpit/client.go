package swarmpit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"log"
	"../setup"
)

var arg = setup.GetArgs()

type Event struct {
	From    string      `json:"from"`
	Message interface{} `json:"message"`
}

func SendEvent(message interface{}) {
	event := Event{From: "DOCKER", Message: message}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(event)

	_, err := http.Post(arg.EventEndpoint, "application/json; charset=utf-8", buffer)
	if err != nil {
		log.Printf("ERROR: Event sending failed: %s", err)
	}
}

func HealthCheck() {
	for {
		<-time.After(5 * time.Second)
		_, err := http.Get(arg.HealthCheckEndpoint)

		if err == nil {
			log.Printf("INFO: Swarmpit OK")
			break;
		}
	}
}
