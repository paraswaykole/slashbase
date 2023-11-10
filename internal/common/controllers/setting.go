package controllers

import (
	"errors"
	"strconv"

	"github.com/slashbaseide/slashbase/internal/common/analytics"
	"github.com/slashbaseide/slashbase/internal/common/dao"
	"github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/internal/common/utils"
	"github.com/slashbaseide/slashbase/pkg/ai"
)

type SettingController struct{}

func (SettingController) GetSingleSetting(name string) (interface{}, error) {
	setting, err := dao.Setting.GetSingleSetting(name)

	switch name {
	case models.SETTING_NAME_OPENAI_KEY:
		return setting.Value, nil
	case models.SETTING_NAME_OPENAI_MODEL:
		if setting.Value == "" {
			return ai.GetOpenAiModel(), nil
		}
		return setting.Value, nil
	}

	if err != nil {
		return "", errors.New("setting not found")
	}

	switch setting.Name {
	case models.SETTING_NAME_APP_ID:
		return setting.UUID().String(), nil
	case models.SETTING_NAME_TELEMETRY_ENABLED:
		return setting.Bool(), nil
	case models.SETTING_NAME_LOGS_EXPIRE:
		return setting.Int(), nil
	}
	return setting.Value, nil
}

func (SettingController) UpdateSingleSetting(name string, value string) error {
	switch name {
	case models.SETTING_NAME_APP_ID:
		return errors.New("cannot update the setting: " + name)
	case models.SETTING_NAME_TELEMETRY_ENABLED:
		if !utils.ContainsString([]string{"true", "false"}, value) {
			return errors.New("cannot update the setting: " + name)
		}
		analytics.SendUpdatedTelemetryEvent(value == "true")
	case models.SETTING_NAME_LOGS_EXPIRE:
		if _, err := strconv.Atoi(value); err != nil {
			return errors.New("cannot update the setting: " + name)
		}
	case models.SETTING_NAME_OPENAI_KEY:
		ai.InitClient(value)
	case models.SETTING_NAME_OPENAI_MODEL:
		err := ai.SetOpenAiModel(value)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid setting name: " + name)
	}
	err := dao.Setting.UpdateSingleSetting(name, value)
	if err != nil {
		return errors.New("there was some problem updating the setting")
	}
	return nil
}
