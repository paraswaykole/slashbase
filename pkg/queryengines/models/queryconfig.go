package models

type QueryConfig struct {
	ReadOnly    bool
	CreateLogFn func(string)
}

func NewQueryConfig(readOnly bool, createLogFn func(string)) *QueryConfig {
	return &QueryConfig{
		ReadOnly:    readOnly,
		CreateLogFn: createLogFn,
	}
}
