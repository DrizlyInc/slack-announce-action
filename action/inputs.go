package main

import (
	"encoding/json"
	"os"

	"github.com/sethvargo/go-githubactions"
)

type ActionInputs struct {
	webhookUrl string
	channel    string
	indicators []Indicator
	// Optional
	statusOverride string
	username       string
	// Derived
	status string
}

type Indicator struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

const (
	Success = "success"
	Failure = "failure"
	Skipped = "skipped"
	Cancelled = "cancelled"
)

func ParseInputs() *ActionInputs {
	webhookUrl := EnvOrFatal("INPUT_WEBHOOK_URL", "Input 'webhook_url' is required")
	channel := EnvOrFatal("INPUT_CHANNEL", "Input 'channel' is required")
	indicators := ParseIndicatorsInput()
	statusOverride := EnvOrDefault("INPUT_STATUS_OVERRIDE", "")
	username := EnvOrDefault("INPUT_USERNAME", "GitHub Actions")

	if len(indicators) == 0 && statusOverride == "" {
		githubactions.Fatalf("Input 'status_override' is required if no 'indicators' are provided")
	}

	return &ActionInputs{
		webhookUrl:     webhookUrl,
		channel:        channel,
		indicators:     indicators,
		statusOverride: statusOverride,
		username:       username,
		status: GetCummulativeStatus(statusOverride, indicators),
	}
}

func ParseIndicatorsInput() []Indicator {
	indicatorsJson := EnvOrDefault("INPUT_INDICATORS", "[]")

	var indicators []Indicator
	err := json.Unmarshal([]byte(indicatorsJson), &indicators)
	if err != nil {
		githubactions.Fatalf("Error parsing input 'indicators': %v", err.Error())
	}

	for idx, indicator := range indicators {
		if indicator.Status == "" {
			githubactions.Fatalf("Missing 'status' in indicator at position [%d]", idx)
		} else if !IsValidStatus(indicator.Status) {
			githubactions.Fatalf("Provided status '%s' is invalid in indicator at position [%d]", indicator.Status, idx)
		} else if indicator.Name == "" {
			githubactions.Fatalf("Missing 'name' in indicator at position [%d]", idx)
		}
	}

	return indicators
}

func IsValidStatus(status string) bool {
	return status == Success || status == Failure || status == Skipped || status == Cancelled
}

func EnvOrDefault(name, def string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}
	return def
}

func EnvOrFatal(name, msg string) string {
	val, ok := os.LookupEnv(name)
	if !ok {
		githubactions.Fatalf(msg)
	}
	return val
}

func GetCummulativeStatus(statusOverride string, indicators []Indicator) string {
	if statusOverride != "" {
		return statusOverride
	}

	for _, indicator := range indicators {
		if indicator.Status == Failure {
			return Failure
		} else if indicator.Status == Cancelled {
			return Cancelled
		}
	}

	return Success
}