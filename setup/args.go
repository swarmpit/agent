package setup

import (
	"os"
	"strconv"
)

type debug struct {
	Event bool `json:"event"`
	Stats bool `json:"stats"`
}

type args struct {
	StatsFrequency      int    `json:"stats_frequency"`
	EventEndpoint       string `json:"event_endpoint"`
	HealthCheckEndpoint string `json:"healthcheck_endpoint"`
	Debug               debug  `json:"debug"`
}

func GetArgs() *args {
	return &args{
		StatsFrequency:      getIntValue(30, "STATS_FREQUENCY"),
		EventEndpoint:       getStringValue("http://app:8080/events", "EVENT_ENDPOINT"),
		HealthCheckEndpoint: getStringValue("http://app:8080/version", "HEALTH_CHECK_ENDPOINT"),
		Debug: debug{
			Event: getBooleanValue(false, "DEBUG_EVENT"),
			Stats: getBooleanValue(false, "DEBUG_STATS")},
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

func getBooleanValue(defValue bool, varName string) bool {
	value := defValue
	env := os.Getenv(varName)
	if len(env) > 0 {
		i, err := strconv.ParseBool(env)
		if err != nil {
			return false
		}
		value = i
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
