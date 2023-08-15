package ai

import (
	openai "github.com/sashabaranov/go-openai"
)

var (
	client      *openai.Client
	openAiModel string = "text-davinci-003"
)

func InitClient(token string) {
	if token == "" {
		client = nil
		return
	}
	client = openai.NewClient(token)
}

func GetOpenAiModel() string {
	return openAiModel
}

func SetOpenAiModel(model string) {
	if model == "" {
		return
	}
	openAiModel = model
}
