package retention

import (
	"time"

	"github.com/jportoles/influxdb092/toml"
)

type Config struct {
	Enabled       bool          `toml:"enabled"`
	CheckInterval toml.Duration `toml:"check-interval"`
}

func NewConfig() Config {
	return Config{Enabled: true, CheckInterval: toml.Duration(10 * time.Minute)}
}
