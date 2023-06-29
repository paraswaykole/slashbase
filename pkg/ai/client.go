package ai

import (
	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

func InitClient(token string) {
	client = openai.NewClient(token)
}
