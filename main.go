package main

import (
	"github.com/koyote/pkg/api"
	"github.com/koyote/pkg/config"
	"github.com/koyote/pkg/telegram"
)

func main() {
	config.LoadConfig()
	go telegram.StartBot()
	api.StartPolling()
}
