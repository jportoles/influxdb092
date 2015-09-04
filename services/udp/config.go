package udp

import "github.com/jportoles/influxdb092/toml"

type Config struct {
	Enabled     bool   `toml:"enabled"`
	BindAddress string `toml:"bind-address"`

	Database     string        `toml:"database"`
	BatchSize    int           `toml:"batch-size"`
	BatchTimeout toml.Duration `toml:"batch-timeout"`
}
