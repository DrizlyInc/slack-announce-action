package main

import (
	"encoding/json"
	"fmt"

	"github.com/sethvargo/go-githubactions"
	"github.com/slack-go/slack"
)

// https://app.slack.com/block-kit-builder

func main() {

	env := ParseEnv()
	inputs := ParseInputs()

	title := fmt.Sprintf("%s build succeeded", env.githubRepositoryName)

	webhookMsg := &slack.WebhookMessage{
		Username: inputs.username,
		Channel: inputs.channel,
		Attachments: []slack.Attachment{{
			Fallback: title,
			Color: "#777777",
			Blocks: slack.Blocks{
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
	fmt.Println(string(b))

	err = slack.PostWebhook(inputs.webhookUrl, webhookMsg)

	if err != nil {
		githubactions.Fatalf("Error posting to slack webhook: %v", err.Error())
	}

}

