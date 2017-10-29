package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os"
)

func HandleEvents(dockerSocket string, eventEndpoint string) {
	client, err := docker.NewClient(dockerSocket)
	if err != nil {
		panic(err)
	}

	listener := make(chan *docker.APIEvents)
	defer func() {
		if err = client.RemoveEventListener(listener); err != nil {
			panic(err)
		}
	}()

	err = client.AddEventListener(listener)
	if err != nil {
		panic(err)
		fmt.Fprintf(os.Stdout, "Event listener registration failed: %s\n", err)
	}

	for {
		msg, ok := <-listener
		if !ok {
			fmt.Fprintf(os.Stdout, "Event channel closed.\n")
			break
		}
		if msg != nil {
			SendEvent(eventEndpoint, *msg)
		}
	}
}
