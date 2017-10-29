package main

import (
	"os"
	"fmt"
)

func main() {
	var dockerSocket, eventEndpoint string
	dockerSocket = os.Getenv("DOCKER_SOCKET")
	eventEndpoint = os.Getenv("EVENT_ENDPOINT")

	if dockerSocket == "" {
		fmt.Fprintf(os.Stderr, "Please specify correct DOCKER_SOCKET property.\n")
		os.Exit(3)
	}

	if eventEndpoint == "" {
		fmt.Fprintf(os.Stderr, "Please specify correct EVENT_ENDPOINT property.\n")
		os.Exit(3)
	}

	HandleEvents(dockerSocket, eventEndpoint)
}
