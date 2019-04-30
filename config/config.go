package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	WebhookUrl   string
	MUFGId       string
	MUFGPassword string
	MFEmail      string
	MFPassword   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Failed to read file: %s", err.Error())
	}

	Config = ConfigList{
		WebhookUrl:   cfg.Section("slack").Key("webhook_url").String(),
		MUFGId:       cfg.Section("mufg").Key("id").String(),
		MUFGPassword: cfg.Section("mufg").Key("password").String(),
		MFEmail:      cfg.Section("moneyforward").Key("email").String(),
		MFPassword:   cfg.Section("moneyforward").Key("password").String(),
	}
}
