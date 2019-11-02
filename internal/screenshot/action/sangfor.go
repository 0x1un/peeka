package action

import (
	"github.com/chromedp/chromedp"
)

func SangforLogin(url, username, password string, sleepTime int) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.SendKeys(`#user`, username, chromedp.ByID),
		chromedp.SendKeys(`#password`, password, chromedp.ByID),
		chromedp.Click(`#button`, chromedp.ByID),
		// chromedp.Click(`#ext-gen167`, chromedp.NodeVisible),
	}
}

// https://10.6.2.5/index.php?name_jpgraph_antispam=747905443
//https://10.6.2.6/index.php?name_jpgraph_antispam=709812881
