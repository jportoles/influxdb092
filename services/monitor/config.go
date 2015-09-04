package monitor

import (
	"time"

	"github.com/jportoles/influxdb092/toml"
)

const (
	// DefaultStatisticsWriteInterval is the interval of time between internal stats are written
	DefaultStatisticsWriteInterval = 1 * time.Minute
)

// Config represents a configuration for the monitor.
type Config struct {
	Enabled       bool          `toml:"enabled"`
	WriteInterval toml.Duration `toml:"write-interval"`
}

func NewConfig() Config {
	return Config{
		Enabled:       false,
		WriteInterval: toml.Duration(DefaultStatisticsWriteInterval),
	}
}
