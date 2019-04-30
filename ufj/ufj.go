package ufj

import (
	"github.com/chromedp/chromedp"
)

const url = "https://entry11.bk.mufg.jp/ibg/dfw/APLIN/loginib/login?_TRANID=AA000_001"

func GetBalance(id string, password string, balance *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("#footer", chromedp.ByID),
		chromedp.SendKeys(`//input[@name='KEIYAKU_NO']`, id),
		chromedp.SendKeys(`//input[@name='PASSWORD']`, password),
		chromedp.Click(`//*[@id="login_frame"]/div/div/div[2]/a/img`),
		chromedp.WaitVisible("#footer", chromedp.ByID),
		chromedp.Text(`#setAmountDisplay`, balance, chromedp.ByID),
	}
}
