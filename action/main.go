package main

import (
	"encoding/json"
	"fmt"

	"github.com/sethvargo/go-githubactions"
	"github.com/slack-go/slack"
)

func main() {

	env := ParseEnv()
	inputs := ParseInputs()

	color, titleSuffix := GetColorAndTitleSuffix(inputs.status)

	title := fmt.Sprintf("%s build %s", env.githubRepositoryName, titleSuffix)

	webhookMsg := &slack.WebhookMessage{
		Username: inputs.username,
		Channel:  inputs.channel,
		Attachments: []slack.Attachment{{
			Fallback: title,
			Color:    color,
			Blocks: slack.Blocks{
				BlockSet: GetMessageBlocks(*env, *inputs),
			},
		}},
	}

	b, err := json.Marshal(webhookMsg)
	githubactions.Infof(fmt.Sprintf("%s\n", string(b)))

	err = slack.PostWebhook(inputs.webhookUrl, webhookMsg)
	if err != nil {
		githubactions.Fatalf("Error posting to slack webhook: %v", err.Error())
	}

}

// https://app.slack.com/block-kit-builder
func GetMessageBlocks(env Environment, inputs ActionInputs) []slack.Block {
	blocks := []slack.Block{
		NewTitleBlock(env),
		NewActorContextBlock(env),
		NewRefContextBlock(env),
		NewCommitContextBlock(env),
	}
	if len(inputs.steps) > 0 {
		blocks = append(blocks, slack.NewDividerBlock(), NewStepsSectionBlock(inputs))
	}
	return blocks
}

// Returns a color for the message attachment and an
// ending for the title based on the status being reported
func GetColorAndTitleSuffix(status string) (string, string) {
	switch status {
	case "success":
		return "#4caf50", "completed successfully!"
	case "failure":
		return "#f44336", "failed!"
	case "cancelled":
		return "#808080", "was cancelled."
	default:
		githubactions.Fatalf("Provided status %s is invalid", status)
		return "", ""
	}
}

// Returns an emoji which represents a given status
func GetStatusEmoji(status string) string {
	switch status {
	case "success":
		return ":white_check_mark:"
	case "failure":
		return ":x:"
	case "cancelled":
		return ":heavy_minus_sign:"
	default:
		githubactions.Fatalf("Provided status %s is invalid", status)
		return ""
	}
}
