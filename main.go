package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/koirand/slack-balance/config"
	"github.com/koirand/slack-balance/moneyforward"
	"github.com/koirand/slack-balance/mufg"
	"github.com/koirand/slack-balance/slack"
	"github.com/koirand/slack-balance/utils"
)

const mufgJson string = `{
	"text": ":moneybag: Balance of MUFG: %s"
}
`

const mfJson string = `{
	"text": ":moneybag: Balance list of Money Forward:\n%s",
}
`

func main() {
	utils.LoggingSetting()

	// create chrome instance
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// get MUFG balance
	var mufgBalance string
	if err := chromedp.Run(ctx, mufg.GetBalance(config.Config.MUFGId, config.Config.MUFGPassword, &mufgBalance)); err != nil {
		log.Fatalf("Failed to get MUFG balance: %s\n", err.Error())
	}
	log.Printf("MUFG balance: %s", mufgBalance)

	// send MUFG balance to slack
	message := fmt.Sprintf(mufgJson, mufgBalance)
	if err := slack.SendMessage(config.Config.WebhookUrl, strings.NewReader(message)); err != nil {
		log.Fatalf("Failed to send to slack: %s\n", err.Error())
	}

	// get Money Forward balance
	var mfBalance string
	if err := chromedp.Run(ctx, moneyforward.GetBalance(config.Config.MFEmail, config.Config.MFPassword)); err != nil {
		log.Fatalf("Failed to get Money Forward balance: %s\n", err.Error())
	}
	for i := 0; i < len(moneyforward.Keys); i++ {
		mfBalance += moneyforward.Keys[i].NodeValue +
			": " + strings.Trim(moneyforward.Values[i].NodeValue, "\n") + "\n"
	}
	log.Printf("Money Forward balance list:\n%s\n", mfBalance)

	// send Money Forward balance to slack
	message = fmt.Sprintf(mfJson, mfBalance)
	if err := slack.SendMessage(config.Config.WebhookUrl, strings.NewReader(message)); err != nil {
		log.Fatalf("Failed to send to slack: %s\n", err.Error())
	}

}
