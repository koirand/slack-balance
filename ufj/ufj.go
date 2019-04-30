package ufj

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

const url = "https://entry11.bk.mufg.jp/ibg/dfw/APLIN/loginib/login?_TRANID=AA000_001"

func GetBalance(id string, password string) (balance string, err error) {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// run task list
	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("#footer"),
		chromedp.SendKeys(`//input[@name='KEIYAKU_NO']`, id),
		chromedp.SendKeys(`//input[@name='PASSWORD']`, password),
		chromedp.Click(`//*[@id="login_frame"]/div/div/div[2]/a/img`),
		chromedp.WaitVisible("#footer"),
		chromedp.Text(`#setAmountDisplay`, &balance),
	)
	if err != nil {
		return "", err
	}

	return balance, nil
}
