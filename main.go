package main

import (
	"os"
	"fmt"
	"swarmpit-ec/event"
)

func main() {
	var dockerEndpoint, eventEndpoint string
	dockerEndpoint = os.Getenv("DOCKER_ENDPOINT")
	eventEndpoint = os.Getenv("EVENT_ENDPOINT")

	if dockerEndpoint == "" {
		dockerEndpoint = "unix:///var/run/docker.sock"
		fmt.Fprintf(os.Stdout, "DOCKER_ENDPOINT not defined. Used default: %s\n", dockerEndpoint)
	}

	if eventEndpoint == "" {
		fmt.Fprintf(os.Stdout, "EVENT_ENDPOINT not defined. Please specify correct url!!!\n")
		os.Exit(3)
	}

	event.HandleEvents(dockerEndpoint, eventEndpoint)
}
