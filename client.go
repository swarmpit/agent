package main

import (
	"github.com/fsouza/go-dockerclient"
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)

type Event struct {
	Type string
	Data docker.APIEvents
}

func SendEvent(endpoint string, eventData docker.APIEvents) {
	event := Event{Type: "DOCKER_EVENT", Data: eventData}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(event)

	res, err := http.Post(endpoint, "application/json; charset=utf-8", buffer)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Event sending failed: %s\n", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Fprintf(os.Stdout, "Event response: %s\n", body)
}
