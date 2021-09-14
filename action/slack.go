package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

// Creates a section block containing a summary of the build
// status and a link to the repository
func NewTitleBlock(env Environment) *slack.SectionBlock {
	repositoryUrl := fmt.Sprintf("%s/%s", env.githubServerUrl, env.githubRepository)
	formattedRepoLink := FormatLink(repositoryUrl, env.githubRepositoryName)

	return &slack.SectionBlock{
		Type: "section",
		Text: &slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: fmt.Sprintf("*%s* build completed successfully!", formattedRepoLink),
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
	viewBuildButton.URL = fmt.Sprintf("%s/%s/actions/runs/%s", env.githubServerUrl, env.githubRepository, env.githubRunNumber)
	return slack.NewAccessory(viewBuildButton)
}

// Creates a context block containing the profile photo
// and name of the github actor (as a link)
func NewActorContextBlock(env Environment) *slack.ContextBlock {
	authorUrl := fmt.Sprintf("%s/%s", env.githubServerUrl, env.githubActor)

	return slack.NewContextBlock(
		"actor",
		NewContextTitle("Actor"),
		&slack.ImageBlockElement{
			Type:     "image",
			ImageURL: fmt.Sprintf("%s.png?size=32", authorUrl),
			AltText:  env.githubActor,
		},
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: FormatLink(authorUrl, env.githubActor),
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
			Text: fmt.Sprintf("`%s`", env.githubRef),
		},
	)
}

// Creates a context block containing a link to the commit
// which triggered this action
func NewCommitContextBlock(env Environment) *slack.ContextBlock {
	commitUrl := fmt.Sprintf("%s/%s/commit/%s", env.githubServerUrl, env.githubRepository, env.githubSha)

	return slack.NewContextBlock(
		"commit",
		NewContextTitle("Commit"),
		&slack.TextBlockObject{
			Type: slack.MarkdownType,
			Text: FormatLink(commitUrl, env.githubSha[0:6]),
		},
	)
}

// Creates a section block containing a list of step titles with
// emoji representation of their statuses
func NewStepsSectionBlock(inputs ActionInputs) *slack.SectionBlock {

	text := ""
	for _, step := range inputs.steps {
		text = text + fmt.Sprintf("%s %s\n", GetStatusEmoji(step.Status), step.Title)
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
