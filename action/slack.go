package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

// Creates a section block containing a summary of the build
// status and a link to the repository
func NewTitleBlock(env Environment, titleEntity, titleSuffix string) *slack.SectionBlock {
	repositoryUrl := fmt.Sprintf("%s/%s", env.GithubServerUrl, env.GithubRepository)
	formattedRepoLink := FormatLink(repositoryUrl, env.GithubRepositoryName)

	return &slack.SectionBlock{
		Type: "section",
		Text: &slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: fmt.Sprintf("*%s* %s %s", formattedRepoLink, titleEntity, titleSuffix),
		},
		Accessory: NewViewBuildAccessory(&env),
	}
}

// Creates a slack accessory containing a button which
// links to the GitHub actions build
func NewViewBuildAccessory(env *Environment) *slack.Accessory {
	viewBuildButton := slack.NewButtonBlockElement(
		"view_build",
		"",
		slack.NewTextBlockObject("plain_text", "View Build", false, false),
	)
	viewBuildButton.URL = fmt.Sprintf("%s/%s/actions/runs/%d", env.GithubServerUrl, env.GithubRepository, env.GithubRunId)
	return slack.NewAccessory(viewBuildButton)
}

// Creates a context block containing the profile photo
// and name of the github actor (as a link)
func NewActorContextBlock(env Environment) *slack.ContextBlock {
	authorUrl := fmt.Sprintf("%s/%s", env.GithubServerUrl, env.GithubActor)

	return slack.NewContextBlock(
		"actor",
		NewContextTitle("Actor"),
		&slack.ImageBlockElement{
			Type:     "image",
			ImageURL: fmt.Sprintf("%s.png?size=32", authorUrl),
			AltText:  env.GithubActor,
		},
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: FormatLink(authorUrl, env.GithubActor),
		},
	)
}

// Creates a context block containing the ref which triggered
// this action (ex. refs/heads/main)
func NewRefContextBlock(env Environment) *slack.ContextBlock {
	return slack.NewContextBlock(
		"ref",
		NewContextTitle("Ref"),
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: fmt.Sprintf("`%s`", env.GithubRef),
		},
	)
}

// Creates a context block containing a link to the commit
// which triggered this action
func NewCommitContextBlock(env Environment) *slack.ContextBlock {
	commitUrl := fmt.Sprintf("%s/%s/commit/%s", env.GithubServerUrl, env.GithubRepository, env.GithubSha)

	return slack.NewContextBlock(
		"commit",
		NewContextTitle("Commit"),
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: FormatLink(commitUrl, env.GithubSha[0:6]),
		},
	)
}

// Creates a section block containing a list of indicator titles with
// emoji representation of their statuses
func NewIndicatorsSectionBlock(inputs ActionInputs) *slack.SectionBlock {

	text := ""
	for _, indicator := range inputs.indicators {
		text = text + fmt.Sprintf("%s %s\n", GetEmojiForStatus(indicator.Status), indicator.Name)
	}

	return slack.NewSectionBlock(
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: text,
		},
		nil,
		nil,
	)
}

// Creates a TextBlockObject with a bolded string to be used
// as a title in a ContextBlock
func NewContextTitle(title string) *slack.TextBlockObject {
	return &slack.TextBlockObject{
		Type: slack.MarkdownType,
		Text: fmt.Sprintf("*%s:*", title),
	}
}

// Creates a formatted link with display text
func FormatLink(url, text string) string {
	return fmt.Sprintf("<%s|%s>", url, text)
}
