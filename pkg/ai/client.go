package ai

import (
	"errors"

	openai "github.com/sashabaranov/go-openai"
	"github.com/slashbaseide/slashbase/internal/common/utils"
)

var (
	client                *openai.Client
	openAiModel           string = openai.GPT3Dot5Turbo
	supportedOpenAIModels        = []string{
		openai.GPT3Dot5TurboInstruct,
		openai.GPT3Dot5Turbo16K,
		openai.GPT3Dot5Turbo,
		openai.GPT3Dot5Turbo1106,
		openai.GPT4,
		openai.GPT432K,
		openai.GPT40613,
		openai.GPT432K0613,
	}
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

func ListSupportedOpenAiModels() []string {
	return supportedOpenAIModels
}

func SetOpenAiModel(model string) error {
	if model == "" {
		return errors.New("cannot be empty")
	}
	if utils.ContainsString(supportedOpenAIModels, model) {
		openAiModel = model
		return nil
	}
	return errors.New("invalid model")
}
