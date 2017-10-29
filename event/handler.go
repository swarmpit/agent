package event

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os"
)

func HandleEvents(dockerEndpoint string, eventEndpoint string) {
	client, err := docker.NewClient(dockerEndpoint)
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
		fmt.Fprintf(os.Stdout, "Failed to add event listener: %s", err)
	}

	for {
		msg, ok := <-listener
		if !ok {
			break
		}
		if msg != nil {
			SendEvent(eventEndpoint, *msg)
		}
	}
}
