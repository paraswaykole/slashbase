package controllers

import (
	"errors"
	"strconv"

	"github.com/slashbaseide/slashbase/internal/analytics"
	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/models"
	"github.com/slashbaseide/slashbase/internal/utils"
)

type SettingController struct{}

func (SettingController) GetSingleSetting(name string) (interface{}, error) {
	setting, err := dao.Setting.GetSingleSetting(name)
	if err != nil {
		return "", errors.New("there was some problem")
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
	default:
		return errors.New("invalid setting name: " + name)
	}
	err := dao.Setting.UpdateSingleSetting(name, value)
	if err != nil {
		return errors.New("there was some problem updating the setting")
	}
	return nil
}
