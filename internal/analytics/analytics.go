package analytics

import (
	"github.com/posthog/posthog-go"
	"github.com/slashbaseide/slashbase/internal/config"
	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/models"
)

const (
	posthogApiKey          = "phc_XSWvMvnTUEH9pLJDVmYfaKaKH8QZtK5fJO8NIiFoNwv"
	posthogInstanceAddress = "https://app.posthog.com"
)

var client posthog.Client

func InitAnalytics() {
	if !config.IsLive() {
		return
	}
	client, _ = posthog.NewWithConfig(posthogApiKey, posthog.Config{
		Endpoint: posthogInstanceAddress,
	})
}

func sendEvent(eventName string, properties map[string]interface{}) {
	if !config.IsLive() {
		return
	}
	setting, err := dao.Setting.GetSingleSetting(models.SETTING_NAME_TELEMETRY_ENABLED)
	if err != nil {
		return
	}
	if !setting.Bool() {
		return
	}
	setting, err = dao.Setting.GetSingleSetting(models.SETTING_NAME_APP_ID)
	if err != nil {
		return
	}
	properties["version"] = config.GetConfig().Version
	client.Enqueue(posthog.Capture{
		DistinctId: setting.UUID().String(),
		Event:      eventName,
		Properties: properties,
	})
}

func SendTelemetryEvent() {
	projectCount, _ := dao.Project.GetAllProjectsCount()
	dbCount, _ := dao.DBConnection.GetAllDBConnectionsCount()
	sendEvent("Telemetry Ping", map[string]interface{}{
		"stats": map[string]interface{}{
			"projects":      projectCount,
			"dbconnections": dbCount,
		},
	})
}

func SendRunQueryEvent() {
	sendEvent("Run Query", map[string]interface{}{})
}

func SendSavedQueryEvent() {
	sendEvent("Saved Query", map[string]interface{}{})
}

func SendLowCodeDataViewEvent() {
	sendEvent("Low Code Data View", map[string]interface{}{})
}

func SendLowCodeModelViewEvent() {
	sendEvent("Low Code Model View", map[string]interface{}{})
}

func SendUpdatedTelemetryEvent(value bool) {
	sendEvent("Updated Telemetry Settings", map[string]interface{}{"value": value})
}
