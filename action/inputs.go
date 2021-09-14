package main

import (
	"encoding/json"
	"os"

	"github.com/sethvargo/go-githubactions"
)

type ActionInputs struct {
	webhookUrl string
	channel    string
	username   string
	status     string
	indicators []Indicator
}

type Indicator struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func ParseInputs() *ActionInputs {
	webhookUrl := EnvOrFatal("INPUT_WEBHOOK_URL", "Input 'webhook_url' is required")
	status := EnvOrFatal("INPUT_STATUS", "Input 'status' is required")
	channel := EnvOrFatal("INPUT_CHANNEL", "Input 'channel' is required")
	username := EnvOrDefault("INPUT_USERNAME", "GitHub Actions")
	indicators := ParseIndicatorsInput()

	return &ActionInputs{
		webhookUrl: webhookUrl,
		status:     status,
		channel:    channel,
		username:   username,
		indicators: indicators,
	}
}

func ParseIndicatorsInput() []Indicator {
	indicatorsJson := EnvOrDefault("INPUT_INDICATORS", "[]")
	var indicators []Indicator
	err := json.Unmarshal([]byte(indicatorsJson), &indicators)
	if err != nil {
		githubactions.Fatalf("Error parsing input 'indicators': %v", err.Error())
	}
	for _, indicator := range indicators {
		if indicator.Status == "" || indicator.Title == "" {
			githubactions.Fatalf("Missing property in provided indicator")
		}
	}
	return indicators
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
