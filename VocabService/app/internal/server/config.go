package server

import "github.com/VIWET/GoVocab/app/internal/repository/sqlstore"

type Config struct {
	Addr     string           `yaml:"addr"`
	LogLevel string           `yaml:"logLevel"`
	DBConfig *sqlstore.Config `yaml:"db"`
}

func NewConfig() *Config {
	return &Config{
		Addr:     ":8080",
		LogLevel: "debug",
		DBConfig: sqlstore.NewConfig(),
	}
}
