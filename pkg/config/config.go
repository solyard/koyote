package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/gookit/slog"
)

type ApplicationConfig struct {
	Global struct {
		ListenPort       string `env:"KOYOTE_API_PORT" envDefault:"8081"`
		TelegramBotToken string `env:"KOYOTE_TELEGRAM_BOT_TOKEN,required"`
	}
	Events struct {
		Job          bool `env:"KOYOTE_ENABLE_JOB_NOTIFICATION"`
		MergeRequest bool `env:"KOYOTE_ENABLE_MR_NOTIFICATION" envDefault:true`
		Note         bool `env:"KOYOTE_ENABLE_NOTE_NOTIFICATION"`
		Pipeline     bool `env:"KOYOTE_ENABLE_PIPELINE_NOTIFICATION" envDefault:true`
		Push         bool `env:"KOYOTE_ENABLE_PUSH_NOTIFICATION"`
		TagPush      bool `env:"KOYOTE_ENABLE_TAG_PUSH_NOTIFICATION"`
	}
}

var GlobalAppConfig ApplicationConfig

func LoadConfig() {
	if err := env.Parse(&GlobalAppConfig); err != nil {
		log.Fatal("Error while parse envs for config to struct. Error: ", err)
	}
}
