package main

import (
	"encoding/json"
	"testing"
)

func Test_NewViewBuildAccessory(t *testing.T) {
	env := Environment{
		githubServerUrl:  "https://github.com",
		githubRepository: "DrizlyInc/slack-announce-action",
		githubRunNumber:  "784",
	}

	accessory := NewViewBuildAccessory(&env)
	accessoryJson, _ := json.Marshal(accessory)

	expected := `{"type":"button","text":{"type":"plain_text","text":"View Build"},"action_id":"view_build","url":"https://github.com/DrizlyInc/slack-announce-action/actions/runs/784"}`

	assertEquals(t, expected, string(accessoryJson))
}

func Test_NewActorContextBlock(t *testing.T) {
	env := Environment{
		githubServerUrl: "https://github.com",
		githubActor:     "benreynolds-drizly",
	}

	block := NewActorContextBlock(env)
	blockJson, _ := json.Marshal(block)

	expected := `{"type":"context","block_id":"actor","elements":[{"type":"mrkdwn","text":"*Actor:*"},{"type":"image","image_url":"https://github.com/benreynolds-drizly.png?size=32","alt_text":"benreynolds-drizly"},{"type":"mrkdwn","text":"\u003chttps://github.com/benreynolds-drizly|benreynolds-drizly\u003e"}]}`

	assertEquals(t, expected, string(blockJson))
}

func Test_NewRefContextBlock(t *testing.T) {
	env := Environment{
		githubRef: "refs/heads/main",
	}

	block := NewRefContextBlock(env)
	blockJson, _ := json.Marshal(block)

	expected := "{\"type\":\"context\",\"block_id\":\"ref\",\"elements\":[{\"type\":\"mrkdwn\",\"text\":\"*Ref:*\"},{\"type\":\"mrkdwn\",\"text\":\"`refs/heads/main`\"}]}"

	assertEquals(t, expected, string(blockJson))
}

func Test_NewCommitContextBlock(t *testing.T) {
	env := Environment{
		githubServerUrl:  "https://github.com",
		githubRepository: "DrizlyInc/slack-announce-action",
		githubSha:        "ecfee3de6b694111add2576049ad73b18417b9ad",
	}

	block := NewCommitContextBlock(env)
	blockJson, _ := json.Marshal(block)

	expected := `{"type":"context","block_id":"commit","elements":[{"type":"mrkdwn","text":"*Commit:*"},{"type":"mrkdwn","text":"\u003chttps://github.com/DrizlyInc/slack-announce-action/commit/ecfee3de6b694111add2576049ad73b18417b9ad|ecfee3\u003e"}]}`

	assertEquals(t, expected, string(blockJson))
}

func Test_NewStepsSectionBlock(t *testing.T) {
	inputs := ActionInputs{
		steps: []Step{
			{Title: "Foo", Status: "success"},
			{Title: "Bar", Status: "failure"},
		},
	}

	block := NewStepsSectionBlock(inputs)
	blockJson, _ := json.Marshal(block)

	expected := `{"type":"section","text":{"type":"mrkdwn","text":":white_check_mark: Foo\n:x: Bar\n"}}`

	assertEquals(t, expected, string(blockJson))
}

func Test_NewContextTitle(t *testing.T) {
	title := NewContextTitle("Ref")
	titleJson, _ := json.Marshal(title)

	expected := `{"type":"mrkdwn","text":"*Ref:*"}`

	assertEquals(t, expected, string(titleJson))
}

func Test_FormatLink(t *testing.T) {
	url := "https://google.com"
	text := "Google"

	expected := "<https://google.com|Google>"

	assertEquals(t, expected, FormatLink(url, text))
}
