package action

import (
	"time"

	"github.com/chromedp/chromedp"
)

// NetworkTrafficAction: 进入指定线路，选取指定时间内的流量状况
func NetworkTrafficActionZBX(url string, timeRange string, sleepTime int) chromedp.Tasks {
	var tasks chromedp.Tasks
	tasks = append(tasks, chromedp.Navigate(url))
	switch timeRange {
	case "1h":
		oneHour := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(4) > a`
		tasks = append(tasks, chromedp.Click(oneHour, chromedp.NodeVisible))
	case "3h":
		threeHours := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(5) > a`

		tasks = append(tasks, chromedp.Click(threeHours, chromedp.NodeVisible))
	case "6h":
		sixHours := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(6) > a`
		tasks = append(tasks, chromedp.Click(sixHours, chromedp.NodeVisible))
	case "12h":
		twelveHours := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(7) > a`
		tasks = append(tasks, chromedp.Click(twelveHours, chromedp.NodeVisible))
	case "24h":
		twentyFourHours := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(8) > a`
		tasks = append(tasks, chromedp.Click(twentyFourHours, chromedp.NodeVisible))
	case "15m":
		m15 := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(2) > a`
		tasks = append(tasks, chromedp.Click(m15, chromedp.NodeVisible))
	case "30m":
		m30 := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(3) > a`
		tasks = append(tasks, chromedp.Click(m30, chromedp.NodeVisible))
	}
	tasks = append(tasks, chromedp.Sleep(time.Duration(sleepTime)*time.Millisecond))
	// tasks = append(tasks, chromedp.WaitVisible(`#graph_full`, chromedp.ByID))
	tasks = append(tasks, chromedp.WaitVisible(`#graph_full`, chromedp.ByID))
	return tasks
}

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
