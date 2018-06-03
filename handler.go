package main

import (
	"net/http"
	"encoding/json"
)

func Info(w http.ResponseWriter, r *http.Request) {
	args := getArgs()
	json.NewEncoder(w).Encode(args)
}
