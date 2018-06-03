package setup

import (
	"os"
	"strconv"
)

type args struct {
	AgentPort           string `json:"agent_port"`
	StatsFrequency      int    `json:"stats_frequency"`
	EventEndpoint       string `json:"event_endpoint"`
	HealthCheckEndpoint string `json:"healthcheck_endpoint"`
}

func GetArgs() *args {
	return &args{
		AgentPort:           getStringValue("8080", "PORT"),
		StatsFrequency:      getIntValue(30, "STATS_FREQUENCY"),
		EventEndpoint:       getStringValue("http://app:8080/events", "EVENT_ENDPOINT"),
		HealthCheckEndpoint: getStringValue("http://app:8080/version", "HEALTH_CHECK_ENDPOINT"),
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

func getIntValue(defValue int, varName string) int {
	value := defValue
	env := os.Getenv(varName)
	if len(env) > 0 {
		i, err := strconv.Atoi(env)
		if err != nil {
			return value
		}
		value = i
	}
	return value
}
