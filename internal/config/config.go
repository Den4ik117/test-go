package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"test-go/pkg/logging"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen struct {
		Type string `yaml:"type"`
		BindIp string `yaml:"bind_ip"`
		Port string `yaml:"port"`
	} `yaml:"listen"`
	MongoDB struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Database string `yaml:"database"`
		AuthDB string `yaml:"auth_db"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Collection string `yaml:"collection"`
	} `yaml:"mongodb"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
