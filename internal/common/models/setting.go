package models

import (
	"strconv"

	"github.com/google/uuid"
)

type Setting struct {
	Name  string `gorm:"primaryKey"`
	Value string `gorm:"not null"`
}

const (
	SETTING_NAME_APP_ID            = "APP_ID"
	SETTING_NAME_TELEMETRY_ENABLED = "TELEMETRY_ENABLED"
	SETTING_NAME_LOGS_EXPIRE       = "LOGS_EXPIRE"
	SETTING_NAME_OPENAI_KEY        = "OPENAI_KEY"
	SETTING_NAME_OPENAI_MODEL      = "OPENAI_MODEL"
)

func NewSetting(name string, value string) *Setting {
	return &Setting{
		Name:  name,
		Value: value,
	}
}

func (s *Setting) UUID() uuid.UUID {
	return uuid.MustParse(s.Value)
}

func (s *Setting) Int() int {
	value, _ := strconv.Atoi(s.Value)
	return value
}

func (s *Setting) Bool() bool {
	return s.Value == "true"
}
