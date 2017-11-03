package uptimerobot

import (
	"github.com/WileESpaghetti/go-uptimemonitor-v2/v2"
)

type Config struct {
	ApiKey string
}

func (c *Config) Client() (uptimerobot.Client, error) {
	return (*uptimerobot.New(c.ApiKey)), nil
}
