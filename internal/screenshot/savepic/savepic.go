package savepic

import (
	"context"
	"io/ioutil"
	"log"
	"path"
	"peeka/internal/screenshot/action"
	"peeka/internal/screenshot/util"
	"strings"

	"github.com/chromedp/chromedp"
	gim "github.com/ozankasikci/go-image-merge"
)

// SaveImg: 访问线路监控图并保存
func SaveImg(ctx context.Context, urls []map[string]string, dir, timeRange string, sleepTime int, quality int64, sltime, sptime int, buf []byte) ([]*gim.Grid, int, error) {
	util.CreateDirIfNotExist(dir)
	var file string
	var grids []*gim.Grid
	count := 0
	for _, x := range urls {
		for k, v := range x {
			if strings.Contains(k, "深信服") {
				if err := chromedp.Run(ctx, action.SangforLogin(v, "admin1", "goodluck@123", sltime, sptime), util.FullScreenshot(quality, &buf)); err != nil {
					return nil, 0, err
				}
			} else {
				if err := chromedp.Run(ctx, action.NetworkTrafficAction(v, timeRange, sleepTime), util.FullScreenshot(quality, &buf)); err != nil {
					return nil, 0, err
				}
			}
			count++
			file = path.Join(dir, "/", k+".png")
			if err := ioutil.WriteFile(file, buf, 0644); err != nil {
				return nil, 0, err
			}
			grids = append(grids, &gim.Grid{ImageFilePath: file})
			log.Printf("写入文件: %s\n", file)
		}
	}
	return grids, count, nil
}
