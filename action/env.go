package main

import (
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/sethvargo/go-githubactions"
)

type Environment struct {
	GithubRunNumber       int    `env:"GITHUB_RUN_NUMBER,notEmpty"`
	GithubRepository      string `env:"GITHUB_REPOSITORY,notEmpty"`
	GithubRepositoryName  string
	GithubRepositoryOwner string
	GithubServerUrl       string `env:"GITHUB_SERVER_URL,notEmpty"`
	GithubRef             string `env:"GITHUB_REF,notEmpty"`
	GithubEventName       string `env:"GITHUB_EVENT_NAME,notEmpty"`
	GithubActor           string `env:"GITHUB_ACTOR,notEmpty"`
	GithubSha             string `env:"GITHUB_SHA,notEmpty"`
}

func ParseEnv() *Environment {

	cfg := Environment{}
	if err := env.Parse(&cfg); err != nil {
		githubactions.Fatalf(err.Error())
	}

	githubRepositorySplit := strings.Split(cfg.GithubRepository, "/")
	if len(githubRepositorySplit) != 2 {
		githubactions.Fatalf("GITHUB_REPOSITORY env var was not formatted as <owner>/<name>")
	}
	cfg.GithubRepositoryOwner = githubRepositorySplit[0]
	cfg.GithubRepositoryName = githubRepositorySplit[1]

	return &cfg
}
