package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/koirand/slack-balance/config"
	"github.com/koirand/slack-balance/slack"
	"github.com/koirand/slack-balance/ufj"
	"github.com/koirand/slack-balance/utils"
)

const jsonBody string = `{
	"text": ":moneybag: 三菱UFJ銀行の残高は %s です！",
}
`

func main() {
	utils.LoggingSetting()
	balance, err := ufj.GetBalance(config.Config.UfjId, config.Config.UfjPassword)
	if err != nil {
		log.Fatalf("Failed to get balance: %s\n", err.Error())
	}
	message := fmt.Sprintf(jsonBody, balance)

	if err := slack.SendMessage(config.Config.WebhookUrl, strings.NewReader(message)); err != nil {
		log.Fatalf("Failed to send to slack: %s\n", err.Error())
	}
}
