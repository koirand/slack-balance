package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/koirand/slack-balance/config"
	"github.com/koirand/slack-balance/slack"
	"github.com/koirand/slack-balance/ufj"
	"github.com/koirand/slack-balance/utils"
)

const jsonStr string = `{
	"text": ":moneybag: 三菱UFJ銀行の残高は %s です！",
}
`

func main() {
	utils.LoggingSetting()

	// create chrome instance
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var balance string
	err := chromedp.Run(ctx, ufj.GetBalance(config.Config.UfjId, config.Config.UfjPassword, &balance))
	if err != nil {
		log.Fatalf("Failed to get balance: %s\n", err.Error())
	}
	log.Printf("UFJ bank balance: %s", balance)

	// send to slack
	message := fmt.Sprintf(jsonStr, balance)
	if err := slack.SendMessage(config.Config.WebhookUrl, strings.NewReader(message)); err != nil {
		log.Fatalf("Failed to send to slack: %s\n", err.Error())
	}
}
