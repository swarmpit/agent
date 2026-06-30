package swarmpit

import (
	"io"
	"log"
	"time"
	"bytes"
	"io/ioutil"
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

	resp, err := http.Post(arg.EventEndpoint, "application/json; charset=utf-8", buffer)
	if err != nil {
		log.Printf("ERROR: Event sending failed: %s", err)
		return
	}
	// Drain and close the body so the underlying keep-alive connection is
	// reused/released. Without this every POST leaks a connection + file
	// descriptor (SendEvent runs on every stats tick), eventually exhausting
	// the agent's FD limit ("too many open files") and growing memory until OOM.
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
}

func HealthCheck() {
	for {
		<-time.After(5 * time.Second)
		resp, err := http.Get(arg.HealthCheckEndpoint)

		if err == nil {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
			log.Printf("INFO: Swarmpit OK")
			break;
		}
	}
}
