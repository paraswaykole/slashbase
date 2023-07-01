package ai

import (
	"context"
	"errors"
	"fmt"
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
	prompt := fmt.Sprintf("### %s SQL tables, with their properties:\n#\n#%s\n#\n### A query to %s:\n\n", dbtype, dbDataModelDescription, text)

	req := openai.CompletionRequest{
		Model:            openai.GPT3TextDavinci003,
		Temperature:      0,
		MaxTokens:        150,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Stop:             []string{"#", ";"},
		Prompt:           prompt,
	}
	resp, err := client.CreateCompletion(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("completion error: %v", err)
	}

	return resp.Choices[0].Text, nil
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
		desc += fmt.Sprintf("# %s (%s) \n", dm.Name, strings.Join(fields, ", "))
	}
	return desc
}
