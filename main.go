package main

import (
	"os"
	"fmt"
)

func main() {
	var eventEndpoint string
	eventEndpoint = os.Getenv("EVENT_ENDPOINT")

	if eventEndpoint == "" {
		fmt.Fprintf(os.Stderr, "Please specify correct EVENT_ENDPOINT property.\n")
		os.Exit(3)
	}

	HandleEvents(eventEndpoint)
}
