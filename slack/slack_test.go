package slack

import (
	"testing"
)

func TestSlackPost(t *testing.T) {
	c := Client{WebhookURL: "https://hooks.slack.com/services/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"}

	c.SendText("Hello Slack from Golang")
}
