package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	WebhookUrl  string
	UfjId       string
	UfjPassword string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err.Error())
	}

	Config = ConfigList{
		WebhookUrl:  cfg.Section("slack").Key("webhook_url").String(),
		UfjId:       cfg.Section("ufj").Key("id").String(),
		UfjPassword: cfg.Section("ufj").Key("password").String(),
	}
}
