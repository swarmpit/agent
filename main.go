package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	router := NewRouter()
	args := getArgs()
	logPrintf("INFO: Waiting for Swarmpit...")
	HealthCheck(args.HealthCheckEndpoint)
	logPrintf("INFO: Swarmpit event collector starting...")
	go HandleEvents(args.EventEndpoint)
	logPrintf("INFO: Swarmpit event collector running")
	logPrintf("INFO: Swarmpit agent listening on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
