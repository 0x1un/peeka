package action

import (
	"time"

	"github.com/chromedp/chromedp"
)

// NetworkTrafficAction: 进入指定线路，选取指定时间内的流量状况
func NetworkTrafficAction(url string, timeRange, sleepTime int) chromedp.Tasks {
	var tasks chromedp.Tasks
	tasks = append(tasks, chromedp.Navigate(url))
	switch timeRange {
	case 1:
		oneHour := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(4) > a`
		tasks = append(tasks, chromedp.Click(oneHour, chromedp.NodeVisible))
	case 3:
		threeHours := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(5) > a`

		tasks = append(tasks, chromedp.Click(threeHours, chromedp.NodeVisible))
	case 6:
		sixHours := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(6) > a`
		tasks = append(tasks, chromedp.Click(sixHours, chromedp.NodeVisible))
	case 12:
		twelveHours := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(7) > a`
		tasks = append(tasks, chromedp.Click(twelveHours, chromedp.NodeVisible))
	case 24:
		twentyFourHours := `#tab_1 > div.time-quick-range > div:nth-child(4) > ul > li:nth-child(8) > a`
		tasks = append(tasks, chromedp.Click(twentyFourHours, chromedp.NodeVisible))
	}
	tasks = append(tasks, chromedp.Sleep(time.Duration(sleepTime)*time.Millisecond))
	// tasks = append(tasks, chromedp.WaitVisible(`#graph_full`, chromedp.ByID))
	tasks = append(tasks, chromedp.WaitVisible(`#graph_full`, chromedp.ByID))
	return tasks
}
