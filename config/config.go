package config

import (
	"log"

	"github.com/go-yaml/yaml"
	"github.com/spf13/viper"

	"github.com/puripat-hugeman/go-clean-todo/todo/repository/postgres"
)

type Config struct {
	Port     string          `mapstructure:"listen_address" yaml:"port"`
	Postgres postgres.Config `mapstructure:"postgres" yaml:"postgres"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetDefault("listen_address", "127.0.0.1:8000")
	viper.SetDefault("postgres.host", "host.docker.internal")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.password", "mypostgres")
	viper.SetDefault("postgres.name", "mypostgres")
	viper.SetDefault("postgres.user", "postgres")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	conf, _ := yaml.Marshal(config)
	log.Printf("Configuration:\n%s\n", conf)
	return config, nil
}
