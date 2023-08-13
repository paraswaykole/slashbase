package ai

import (
	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

var OpenAiApiKey string = ""

func InitClient(token string) {
	if token == "" {
		client = nil
		return
	}
	OpenAiApiKey = token
	client = openai.NewClient(token)
}

var OpenAiModel = "text-davinci-003"

func SetGptModel(token string) {
	if token == "" {
		return
	}
	OpenAiModel = token
}
