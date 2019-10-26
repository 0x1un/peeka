package action

import "github.com/chromedp/chromedp"

func SigninAction(loginUrl, username, password string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("http://" + loginUrl),
		chromedp.SendKeys(`#name`, username, chromedp.ByID),
		chromedp.SendKeys(`#password`, password, chromedp.ByID),
		chromedp.Click(`#enter`, chromedp.NodeVisible),
		chromedp.Click(`#sub_view > li:nth-child(1) > a`, chromedp.NodeVisible),
		// chromedp.Sleep(time.Duration(a.timeout) * time.Second),
	}
}
