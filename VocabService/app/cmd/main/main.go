package main

import (
	"flag"

	"github.com/VIWET/GoVocab/app/internal/app"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.yaml", "path to server config file")
}

func main() {
	flag.Parse()

	app.Run(configPath)
}
