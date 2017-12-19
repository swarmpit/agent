package main

func main() {
	args := getArgs()
	logPrintf("INFO: Swarmpit event collector starting...")
	logPrintf("INFO: EVENT_ENDPOINT: %s", args.EventEndpoint)
	logPrintf("INFO: Waiting for Swarmpit...")
	HealthCheck(args.HealthCheckEndpoint)
	HandleEvents(args.EventEndpoint)
}
