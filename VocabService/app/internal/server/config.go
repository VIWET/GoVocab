package server

type Config struct {
	Addr     string `yaml:"addr"`
	LogLevel string `yaml:"logLevel"`
}

func NewConfig() *Config {
	return &Config{
		Addr:     ":8080",
		LogLevel: "debug",
	}
}
