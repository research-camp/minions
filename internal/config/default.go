package config

import (
	"github.com/amirhnajafiz/xerox/internal/logger"
	"github.com/amirhnajafiz/xerox/internal/metric"
	"github.com/amirhnajafiz/xerox/proxy"
)

func Default() Config {
	return Config{
		Logger: logger.Config{
			Level: "debug",
		},
		Metric: metric.Config{
			Enable: false,
			Host:   "",
		},
		Proxy: proxy.Config{
			Address: ":8080",
			BaseURL: "",
		},
	}
}
