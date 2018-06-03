package main

import (
	"log"
	"net/http"
	"github.com/docker/docker/client"
	"github.com/swarmpit/agent/setup"
	"github.com/swarmpit/agent/swarmpit"
	"github.com/swarmpit/agent/swarmpit/task"
)

var arg = setup.GetArgs()

func main() {
	router := NewRouter()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Printf("ERROR: Docker client initialization failed.")
		panic(err)
	}
	log.Printf("INFO: Waiting for Swarmpit...")
	swarmpit.HealthCheck()
	go task.HandleEvents(cli)
	log.Printf("INFO: Event collector started.")
	go task.HandleStats(cli)
	log.Printf("INFO: Stats collector started.")
	log.Printf("INFO: Swarmpit agent listening on port: %s", arg.AgentPort)
	log.Fatal(http.ListenAndServe(":"+arg.AgentPort, router))
}
