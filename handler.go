package main

import (
	"net/http"
	"encoding/json"
	"github.com/swarmpit/agent/setup"
)

func Info(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(setup.GetArgs())
}
