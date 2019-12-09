package swarmpit

import (
	"log"
	"time"
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/Kenits/agent/setup"
)

var arg = setup.GetArgs()

type EventType string

const (
	EVENT EventType = "event"
	STATS EventType = "stats"
)

type Event struct {
	EventType EventType   `json:"type"`
	Message   interface{} `json:"message"`
}

func SendEvent(eventType EventType, message interface{}) {
	event := Event{EventType: eventType, Message: message}
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
