package ai

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	qemodels "github.com/slashbaseide/slashbase/pkg/queryengines/models"
)

func GenerateSQL(dbtype, text string, datamodels []*qemodels.DBDataModel) (string, error) {

	if client == nil {
		return "", errors.New("update openai key in advanced settings")
	}

	if !(dbtype == qemodels.DBTYPE_POSTGRES || dbtype == qemodels.DBTYPE_MYSQL) {
		return "", errors.New("unsupported database type")
	}

	dbDataModelDescription := generateDBDataModelsDescription(datamodels)
	systemMessage := "No text, just write SQL query with ```sql."
	prompt := fmt.Sprintf("%s SQL tables, with their properties:\n%s\n\nWrite a query to %s", dbtype, dbDataModelDescription, text)

	req := openai.ChatCompletionRequest{
		Model: openAiModel,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemMessage,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("completion error: %v", err)
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("empty response")
	}

	messageContent := resp.Choices[0].Message.Content

	if strings.Contains(messageContent, "```sql") {
		re, _ := regexp.Compile("```sql[\\s\\S]([\\s\\S]*)[\\s\\S]```")
		submatches := re.FindStringSubmatch(messageContent)
		if len(submatches) == 2 {
			messageContent = submatches[1]
		}
	} else if strings.Contains(messageContent, "```") {
		re, _ := regexp.Compile("```[\\s\\S]([\\s\\S]*)[\\s\\S]```")
		submatches := re.FindStringSubmatch(messageContent)
		if len(submatches) == 2 {
			messageContent = submatches[1]
		}
	}

	return messageContent, nil
}

func generateDBDataModelsDescription(datamodels []*qemodels.DBDataModel) string {
	desc := ""
	for _, dm := range datamodels {
		fields := []string{}
		for _, field := range dm.Fields {
			fname := field.Name
			if field.IsPrimary {
				fname += " pk"
			}
			fields = append(fields, fname)
		}
		desc += fmt.Sprintf("- %s (%s) \n", dm.Name, strings.Join(fields, ", "))
	}
	return desc
}
