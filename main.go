package main

import (
	"github.com/koyote/pkg/api"
	"github.com/koyote/pkg/config"
	"github.com/koyote/pkg/redis"
	"github.com/koyote/pkg/telegram"
)

func main() {
	config.LoadConfig()
	if config.GlobalAppConfig.Redis.Enabled {
		go redis.ConnectToRedis()
	}
	go telegram.StartBot()
	api.StartPolling()
}
