package tasks

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/slashbaseide/slashbase/internal/common/analytics"
	"github.com/slashbaseide/slashbase/internal/common/config"
	commondao "github.com/slashbaseide/slashbase/internal/common/dao"
	"github.com/slashbaseide/slashbase/internal/common/models"
	"github.com/slashbaseide/slashbase/internal/server/dao"
	"github.com/slashbaseide/slashbase/pkg/queryengines"
	"github.com/slashbaseide/slashbase/pkg/sshtunnel"
)

func InitCron() {
	scheduler := gocron.NewScheduler(time.UTC)
	clearQueryEngineUnusedConnections(scheduler)
	clearOldLogs(scheduler)
	if !config.IsDesktop() {
		sendTelemetryEvents(scheduler)
	}
	scheduler.StartAsync()
}

func clearQueryEngineUnusedConnections(s *gocron.Scheduler) {
	s.Every(5).Minutes().Do(func() {
		sshtunnel.RemoveUnusedTunnels()
		queryengines.RemoveUnusedConnections()
	})
}

func clearOldLogs(s *gocron.Scheduler) {

	clear := commondao.DBQueryLog.ClearOldLogs
	if !config.IsDesktop() {
		clear = dao.DBQueryLog.ClearOldLogs
	}

	s.Every(1).Day().Do(func() {
		setting, err := commondao.Setting.GetSingleSetting(models.SETTING_NAME_LOGS_EXPIRE)
		if err != nil {
			return
		}
		days := setting.Int()
		err = clear(days)
		if err != nil {
			return
		}
	})
}

func sendTelemetryEvents(s *gocron.Scheduler) {
	s.Every(1).Day().Do(func() {
		analytics.SendTelemetryEvent()
	})
}
