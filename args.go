package main

import (
	"os"
)

type args struct {
	EventEndpoint string
}

func getArgs() *args {
	return &args{
		EventEndpoint: getStringValue("http://app:8080/events", "EVENT_ENDPOINT"),
	}
}

func getStringValue(defValue string, varName string) string {
	value := defValue
	env := os.Getenv(varName)
	if len(env) > 0 {
		value = env
	}
	return value
}
