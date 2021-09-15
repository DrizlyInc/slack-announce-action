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
	titleSuffix := GetTitleSuffixForStatus(inputs.status)

	title := fmt.Sprintf("%s %s %s", env.githubRepositoryName, inputs.titleEntity, titleSuffix)

	webhookMsg := &slack.WebhookMessage{
		Username: inputs.username,
		Channel:  inputs.channel,
		Attachments: []slack.Attachment{{
			Fallback: title,
			Color:    GetColorForStatus(inputs.status),
			Blocks: slack.Blocks{
				BlockSet: GetMessageBlocks(*env, *inputs, titleSuffix),
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
func GetMessageBlocks(env Environment, inputs ActionInputs, titleSuffix string) []slack.Block {
	blocks := []slack.Block{
		NewTitleBlock(env, inputs.titleEntity, titleSuffix),
		NewActorContextBlock(env),
		NewRefContextBlock(env),
		NewCommitContextBlock(env),
	}
	if len(inputs.indicators) > 0 {
		blocks = append(blocks, slack.NewDividerBlock(), NewIndicatorsSectionBlock(inputs))
	}
	return blocks
}

// Returns a color representing the given status
func GetColorForStatus(status string) string {
	switch status {
	case Success:
		return "#4caf50"
	case Failure:
		return "#f44336"
	case Cancelled:
		return "#808080"
	case Skipped:
		return "#808080"
	default:
		githubactions.Fatalf("Provided status '%s' is invalid", status)
		return ""
	}
}

// Returns the end of a sentence announcing the given status
func GetTitleSuffixForStatus(status string) string {
	switch status {
	case Success:
		return "completed successfully!"
	case Failure:
		return "failed!"
	case Cancelled:
		return "was cancelled."
	case Skipped:
		return "was skipped."
	default:
		githubactions.Fatalf("Provided status '%s' is invalid", status)
		return ""
	}
}

// Returns an emoji which represents a given status
func GetEmojiForStatus(status string) string {
	switch status {
	case Success:
		return ":white_check_mark:"
	case Failure:
		return ":x:"
	case Cancelled:
		return ":grey_exclamation:"
	case Skipped:
		return ":heavy_minus_sign:"
	default:
		githubactions.Fatalf("Provided status '%s' is invalid", status)
		return ""
	}
}
