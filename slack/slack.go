// Package slack ... This package provides slack integration to post message
// To use this
// 		- get person "webhook" URL from slack.com
//		- Create a client and send the message
//				c := Client{ WebhookURL: "https://hooks.slack.com/services/..."}
//				c.SendText("Hello Slack")
package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Message ... message to post on slack
type Message struct {
	Text      string `json:"text"`
	Channel   string `json:"channel:omitempty"`    // optional
	UserName  string `json:"username:omitempty"`   // optional
	IconURL   string `json:"icon_url:omitempty"`   // optional
	IconEmoji string `json:"icon_emoji:omitempty"` // optional
}

// Client ... used to send message to slack
type Client struct {
	WebhookURL string
	Channel    string
	UserName   string
	IconURL    string
	IconEmoji  string
}

// SendText ... sends text using client settings
func (c *Client) SendText(text string) error {
	var msg = Message{
		Text:      text,
		Channel:   c.Channel,
		UserName:  c.UserName,
		IconURL:   c.IconURL,
		IconEmoji: c.IconEmoji,
	}

	return c.SendMessage(&msg)
}

// SendMessage .. sends message with custom fields
func (c *Client) SendMessage(message *Message) error {
	data, _ := json.Marshal(message)
	fmt.Printf("Sending message: %s\n", string(data))
	resp, err := http.Post(c.WebhookURL, "application/json", bytes.NewBuffer(data))
	if resp != nil && resp.Body != nil {
		_ = resp.Body.Close()
	}

	return err
}
