package main

import (
	"strings"

	"github.com/sethvargo/go-githubactions"
)

type Environment struct {
	githubRunNumber string
	githubRepository string
	githubRepositoryName string
	githubRepositoryOwner string
	githubServerUrl string
	githubRef string
	githubEventName string
	githubActor string
	githubSha string
}

func ParseEnv() *Environment {

	githubRunNumber := EnvOrFatal("GITHUB_RUN_NUMBER", "Failed to read GITHUB_RUN_NUMBER from environment")
	githubServerUrl := EnvOrFatal("GITHUB_SERVER_URL", "Failed to read GITHUB_SERVER_URL from environment")
	githubRef := EnvOrFatal("GITHUB_REF", "Failed to read GITHUB_REF from environment")
	githubEventName := EnvOrFatal("GITHUB_EVENT_NAME", "Failed to read GITHUB_EVENT_NAME from environment")
	githubActor := EnvOrFatal("GITHUB_ACTOR", "Failed to read GITHUB_ACTOR from environment")
	githubSha := EnvOrFatal("GITHUB_SHA", "Failed to read GITHUB_SHA from environment")

	githubRepository := EnvOrFatal("GITHUB_REPOSITORY", "Failed to read GITHUB_REPOSITORY from environment")
	githubRepositorySplit := strings.Split(githubRepository, "/")
	if len(githubRepositorySplit) != 2 {
		githubactions.Fatalf("GITHUB_REPOSITORY env var was not formatted as <owner>/<name>")
	}

	return &Environment{
		githubRunNumber: githubRunNumber,
		githubRepository: githubRepository,
		githubRepositoryName: githubRepositorySplit[1],
		githubRepositoryOwner: githubRepositorySplit[0],
		githubServerUrl: githubServerUrl,
		githubRef: githubRef,
		githubEventName: githubEventName,
		githubActor: githubActor,
		githubSha: githubSha,
	}
}