package tasks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/slashbaseide/slashbase/internal/config"
	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/models"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
	"github.com/slashbaseide/slashbase/pkg/sshtunnel"
)

func InitCron() {
	scheduler := gocron.NewScheduler(time.UTC)
	clearQueryEngineUnusedConnections(scheduler)
	clearOldLogs(scheduler)
	telemetryPings(scheduler)
	scheduler.StartAsync()
}

func clearQueryEngineUnusedConnections(s *gocron.Scheduler) {
	s.Every(5).Minutes().Do(func() {
		sshtunnel.RemoveUnusedTunnels()
		queryengines.RemoveUnusedConnections()
	})
}

func clearOldLogs(s *gocron.Scheduler) {
	s.Every(1).Day().Do(func() {
		setting, err := dao.Setting.GetSingleSetting(models.SETTING_NAME_LOGS_EXPIRE)
		if err != nil {
			return
		}
		days := setting.Int()
		err = dao.DBQueryLog.ClearOldLogs(days)
		if err != nil {
			return
		}
	})
}

func telemetryPings(s *gocron.Scheduler) {
	if !config.IsLive() {
		return
	}
	s.Every(1).Day().Do(func() {
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
		values := map[string]interface{}{
			"api_key": "phc_XSWvMvnTUEH9pLJDVmYfaKaKH8QZtK5fJO8NIiFoNwv",
			"event":   "Telemetry Ping",
			"properties": map[string]string{
				"distinct_id": setting.UUID().String(),
				"version":     config.GetConfig().Version,
			},
		}
		json_data, _ := json.Marshal(values)
		http.Post("https://app.posthog.com/capture/", "application/json",
			bytes.NewBuffer(json_data))
	})
}
