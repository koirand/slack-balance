package moneyforward

import (
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var Keys, Values []*cdp.Node

const url = "https://moneyforward.com"

func GetBalance(email string, password string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url + "/users/sign_in"),
		chromedp.WaitVisible("#new_sign_in_session_service", chromedp.ByID),
		chromedp.SendKeys(`//*[@id="sign_in_session_service_email"]`, email),
		chromedp.SendKeys(`//*[@id="sign_in_session_service_password"]`, password),
		chromedp.Click(`//*[@id="login-btn-sumit"]`),
		chromedp.Sleep(1 * time.Second),
		chromedp.Navigate(url + "/accounts"),
		chromedp.WaitVisible("#footer", chromedp.ByID),
		chromedp.Nodes(`//*[@id="account-table"]//tr//td[1]//a[1]/text()`, &Keys),
		chromedp.Nodes(`//*[@id="account-table"]//tr//td[2]/text()`, &Values),
	}
}
