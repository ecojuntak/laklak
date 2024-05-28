package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

type main struct {
	Server   server   `yaml:"app"`
	Database database `yaml:"database"`
}

type server struct {
	Host     string `yaml:"host"`
	GrpcPort string `yaml:"grpcPort"`
	HttpPort string `yaml:"httpPort"`
}

type database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var mainConfig *main

func Config() *main {
	return mainConfig
}

func DatabaseConfig() database {
	return mainConfig.Database
}

func ServerConfig() server {
	return mainConfig.Server
}

func LoadConfig(configFile string) (err error) {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		slog.Info("loading config from file", "filename", configFile)
	}

	if err := viper.Unmarshal(&mainConfig); err != nil {
		slog.Error("error unmarshilling config: %s\n", err.Error())
	}

	return
}
