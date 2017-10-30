package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
)

type Event struct {
	From    string
	Message interface{}
}

func SendEvent(endpoint string, message interface{}) {
	event := Event{From: "DOCKER", Message: message}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(event)

	res, err := http.Post(endpoint, "application/json; charset=utf-8", buffer)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Event sending failed: %s\n", err)
	} else {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Event response mallformed: %s\n", err)
		}
		fmt.Fprintf(os.Stdout, "Event response: %s\n", body)
	}
}
