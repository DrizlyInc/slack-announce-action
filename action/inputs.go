package main

import (
	"os"

	"github.com/sethvargo/go-githubactions"
)

type ActionInputs struct {
	webhookUrl string
	channel string
	username string
}

func ParseInputs() *ActionInputs {
	webhookUrl := EnvOrFatal("INPUT_WEBHOOK_URL", "Input 'webhook_url' is required")
	channel := EnvOrDefault("INPUT_CHANNEL", "webhook-playground")
	username := EnvOrDefault("INPUT_USERNAME", "GitHub Actions")

	return &ActionInputs{
		webhookUrl: webhookUrl,
		channel: channel,
		username: username,
	}
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