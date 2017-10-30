package main

import (
	"fmt"
	"context"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"os"
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
				fmt.Fprintf(os.Stdout, "Event channel error: %s\n", err)
			}
			break loop
		case msg, ok := <-messages:
			if !ok {
				fmt.Fprintf(os.Stdout, "Event channel closed.\n")
				break loop
			}
			SendEvent(eventEndpoint, msg)
		}
	}
}
