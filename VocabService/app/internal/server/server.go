package server

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type server struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *server {
	return &server{
		config: config,
		logger: logrus.New(),
	}
}

func (s *server) Run() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info(fmt.Sprintf("logger configured on level: %s", s.config.LogLevel))

	return nil
}

func (s *server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}
