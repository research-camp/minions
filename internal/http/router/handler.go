package router

import "github.com/amirhnajafiz/minions/internal/config"

type Handler struct {
	Cfg     config.RouterConfig
	Metrics Metrics
}
