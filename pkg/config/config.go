package config

import (
	"time"

	"github.com/caarlos0/env/v6"
	log "github.com/gookit/slog"
	"github.com/sony/gobreaker"
)

var RedisCB = gobreaker.NewCircuitBreaker(
	gobreaker.Settings{
		Name:        "redis-circutit-breaker",
		MaxRequests: 10,
		Timeout:     10 * time.Second,
		Interval:    10 * time.Second,
	},
)

type ApplicationConfig struct {
	Global struct {
		ListenPort       string `env:"KOYOTE_API_PORT" envDefault:"8081"`
		TelegramBotToken string `env:"KOYOTE_TELEGRAM_BOT_TOKEN,required"`
	}
	Events struct {
		Job          bool `env:"KOYOTE_ENABLE_JOB_NOTIFICATION" envDefault:false`
		MergeRequest bool `env:"KOYOTE_ENABLE_MR_NOTIFICATION" envDefault:true`
		Note         bool `env:"KOYOTE_ENABLE_NOTE_NOTIFICATION envDefault:false`
		Pipeline     bool `env:"KOYOTE_ENABLE_PIPELINE_NOTIFICATION envDefault:true`
		Push         bool `env:"KOYOTE_ENABLE_PUSH_NOTIFICATION" envDefault:false`
		TagPush      bool `env:"KOYOTE_ENABLE_TAG_PUSH_NOTIFICATION" envDefault:false`
	}
	Redis struct {
		Enabled                     bool   `env:"KOYOTE_REDIS_ENABLED" envDefault:false`
		CheckUnsendedEventsInterval int    `env:"KOYOTE_REDIS_CHECK_UNSENDED_EVENTS_INTEVAL,unset"`
		UnsendendTaskTTL            int    `env:"KOYOTE_REDIS_UNSENDED_TASK_TTL,unset"`
		Host                        string `env:"KOYOTE_REDIS_INSTANCE_URI,unset"`
		Port                        string `env:"KOYOTE_REDIS_INSTANCE_PORT,unset"`
		Auth                        struct {
			Username string `env:"KOYOTE_REDIS_USERNAME,unset"`
			Password string `env:"KOYOTE_REDIS_PASSWORD,unset"`
		}
	}
}

var GlobalAppConfig ApplicationConfig

func LoadConfig() {
	if err := env.Parse(&GlobalAppConfig); err != nil {
		log.Fatal("Error while parse envs for config to struct. Error: ", err)
	}
}
