package config

import (
	"os"

	"github.com/mattn/go/src/log"
	"gopkg.in/ini.v1"
)

type ConfList struct {
	LogFile string
	Address string
	Port    int
}

var Config ConfList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfList{
		LogFile: cfg.Section("gotello").Key("log_file").String(),
		Address: cfg.Section("web").Key("address").String(),
		Port:    cfg.Section("web").Key("port").MustInt(),
	}
}
