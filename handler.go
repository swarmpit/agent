package main

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"io"
	"net/http"
	"encoding/json"
	"time"
)

func HandleEvents(eventEndpoint string) {
	fallbackAttempts := 5
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	messages, errs := cli.Events(context.Background(), types.EventsOptions{})

loop:
	for {
		select {
		case err := <-errs:
			if err != nil && err != io.EOF {
				logPrintf("ERROR: Event channel error: %s", err)
			}
			break loop
		case msg, ok := <-messages:
			if !ok {
				logPrintf("ERROR: Event channel closed.")
				break loop
			}
			SendEvent(eventEndpoint, msg)
			fallbackAttempts = 5
		}
	}

	if fallbackAttempts > 0 {
		fallbackAttempts -= 1
		time.Sleep(2 * time.Second)
		goto loop
	} else {
		panic("Event collector is broken. Shutdown!!!")
	}
}

func Info(w http.ResponseWriter, r *http.Request) {
	args := getArgs()
	json.NewEncoder(w).Encode(args)
}
