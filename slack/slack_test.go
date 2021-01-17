package slack

import (
	"testing"
)

func TestSlackPost(t *testing.T) {
	c := Client{WebhookURL: GetWebhookURL()}
	c.SendText("Hello Slack from Golang")
}
