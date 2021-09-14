package main

import (
	"encoding/json"
	"os"

	"github.com/sethvargo/go-githubactions"
)

type ActionInputs struct {
	webhookUrl string
	channel string
	username string
	status string
	steps []Step
}

type Step struct {
	Title string `json:"title"`
	Status string `json:"status"`
}

func ParseInputs() *ActionInputs {
	webhookUrl := EnvOrFatal("INPUT_WEBHOOK_URL", "Input 'webhook_url' is required")
	status := EnvOrFatal("INPUT_STATUS", "Input 'status' is required")
	channel := EnvOrDefault("INPUT_CHANNEL", "webhook-playground")
	username := EnvOrDefault("INPUT_USERNAME", "GitHub Actions")
	steps := ParseStepsInput()

	return &ActionInputs{
		webhookUrl: webhookUrl,
		status: status,
		channel: channel,
		username: username,
		steps: steps,
	}
}

func ParseStepsInput() []Step {
	stepsJson := EnvOrDefault("INPUT_STEPS", "[]")
	var steps []Step
	err := json.Unmarshal([]byte(stepsJson), &steps)
	if err != nil {
		githubactions.Fatalf("Error parsing input 'steps': %v", err.Error())
	}
	for _, step := range steps {
		if step.Status == "" || step.Title == "" {
			githubactions.Fatalf("Missing property in provided step")
		}
	}
	return steps
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