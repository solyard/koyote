package main

import (
	"github.com/koyote/pkg/api"
	"github.com/koyote/pkg/config"
	"github.com/koyote/pkg/redis"
	"github.com/koyote/pkg/telegram"
)

func main() {
	go telegram.StartBot()
	config.LoadConfig()
	api.StartPolling()
	if config.GlobalAppConfig.Redis.Enabled {
		redis.ConnectToRedis()
	}
}
