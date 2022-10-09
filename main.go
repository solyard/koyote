package main

import (
	"github.com/koyote/pkg/api"
	"github.com/koyote/pkg/telegram"
)

func main() {
	go telegram.StartBot()
	api.StartPolling()
}
