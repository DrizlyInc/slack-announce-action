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
		Channel: inputs.channel,
		Attachments: []slack.Attachment{{
			Fallback: title,
			Color: color,
			Blocks: slack.Blocks{ // https://app.slack.com/block-kit-builder
				BlockSet: []slack.Block{
					NewTitleBlock(*env),
					NewActorContextBlock(*env),
					NewRefContextBlock(*env),
					NewCommitContextBlock(*env),
					slack.NewDividerBlock(),
					slack.NewSectionBlock(
						&slack.TextBlockObject{
							Type: slack.MarkdownType,
							Text: ":white_check_mark: `terraform apply` on dev-general",
						},
						nil,
						nil,
					),
				},
			},
		}},
	}

	b, err := json.Marshal(webhookMsg);
	githubactions.Infof(string(b))

	err = slack.PostWebhook(inputs.webhookUrl, webhookMsg)
	if err != nil {
		githubactions.Fatalf("Error posting to slack webhook: %v", err.Error())
	}

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