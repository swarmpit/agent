package swarmpit

import (
	"log"
	"time"
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/swarmpit/agent/setup"
)

var arg = setup.GetArgs()

type EventType string

const (
	EVENT EventType = "event"
	STATS EventType = "stats"
	EMPTY           = ""
	TAB             = "\t"
)

type Event struct {
	EventType EventType   `json:"type"`
	Message   interface{} `json:"message"`
}

func SendEvent(eventType EventType, message interface{}) {
	event := Event{EventType: eventType, Message: message}
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(EMPTY, TAB)
	encoder.Encode(event)

	if eventType == STATS && arg.Debug.Stats == true {
		log.Printf("DEBUG: Host stats: %s", buffer)
	}

	if eventType == EVENT && arg.Debug.Event == true {
		log.Printf("DEBUG: Docker event: %s", buffer)
	}

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
