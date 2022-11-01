package tasks

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"slashbase.com/backend/internal/config"
)

func InitCron() {

	if !config.IsLive() {
		return
	}

	scheduler := gocron.NewScheduler(time.UTC)

	telemetryPings(scheduler)

	scheduler.StartAsync()
}

func telemetryPings(s *gocron.Scheduler) {
	s.Every(1).Day().Do(func() {
		values := map[string]interface{}{
			"api_key": "phc_XSWvMvnTUEH9pLJDVmYfaKaKH8QZtK5fJO8NIiFoNwv",
			"event":   "Telemetry Ping",
			"properties": map[string]string{
				"distinct_id": config.GetConfig().TelemetryID,
				"version":     config.VERSION,
			},
		}
		json_data, err := json.Marshal(values)
		if err != nil {
			log.Fatal(err)
		}

		_, err = http.Post("https://app.posthog.com/capture/", "application/json",
			bytes.NewBuffer(json_data))

		if err != nil {
			log.Fatal(err)
		}
	})
}
