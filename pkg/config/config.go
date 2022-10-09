package config

import (
	"log"
	"os"

	"github.com/koyote/pkg/redis"
	"gopkg.in/yaml.v3"
)

type ApplicationConfig struct {
	Global struct {
		ListenPort string `yaml:"listen_port"`
	} `yaml:"global"`
	Events struct {
		Job          bool `yaml:"job"`
		MergeRequest bool `yaml:"merge_request"`
		Note         bool `yaml:"note"`
		Pipeline     bool `yaml:"pipeline"`
		Push         bool `yaml:"push"`
		TagPush      bool `yaml:"tag_push"`
	} `yaml:"events"`
	Redis struct {
		Enabled bool   `yaml:"enabled"`
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Auth    struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"auth"`
	} `yaml:"redis"`
}

var GlobalAppConfig ApplicationConfig

func LoadConfig() {
	configFile, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal("Error while reading application config file. Application crashed. Error: ", err)
		return
	}

	err = yaml.Unmarshal(configFile, &GlobalAppConfig)
	if err != nil {
		log.Fatal("Error while unmarshal application config to struct. Error: ", err)
		return
	}

	if GlobalAppConfig.Redis.Enabled {
		redis.ConnectToRedis()
	}
}
