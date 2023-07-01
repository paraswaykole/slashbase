package ai

import (
	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

func InitClient(token string) {
	if token == "" {
		client = nil
		return
	}
	client = openai.NewClient(token)
}
