package main

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"io"
)

func HandleEvents(eventEndpoint string) {
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
		}
	}
}
