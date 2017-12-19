package main

func main() {
	args := getArgs()
	logPrintf("INFO: Event collector starting...")
	logPrintf("INFO: EVENT_ENDPOINT: %s", args.EventEndpoint)
	HandleEvents(args.EventEndpoint)
}
