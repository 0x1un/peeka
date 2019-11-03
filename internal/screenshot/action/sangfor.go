package action

import (
	"github.com/chromedp/chromedp"

	"time"
)

func SangforLogin(url, username, password string, logintime, pagetime int) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.SendKeys(`#user`, username, chromedp.ByID),
		chromedp.SendKeys(`#password`, password, chromedp.ByID),
		chromedp.Sleep(time.Duration(logintime) * time.Second),
		chromedp.Click(`#button`, chromedp.ByID),
		chromedp.Sleep(time.Duration(pagetime) * time.Second),
		chromedp.WaitVisible(`#ext-gen148`, chromedp.ByID),
	}
}
