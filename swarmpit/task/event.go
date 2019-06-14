package task

import (
	"io"
	"log"
	"context"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/Kenits/agent/swarmpit"
)

func HandleEvents(cli *client.Client) {
	messages, errs := cli.Events(context.Background(), types.EventsOptions{})

loop:
	for {
		select {
		case err := <-errs:
			if err != nil && err != io.EOF {
				log.Printf("ERROR: Event channel error: %s", err)
			}
			break loop
		case msg, ok := <-messages:
			if !ok {
				log.Printf("ERROR: Event channel closed.")
				break loop
			}
			swarmpit.SendEvent(swarmpit.EVENT, msg)
		}
	}
	panic("Event collector is broken. Shutdown!!!")
}
