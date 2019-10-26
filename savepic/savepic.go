package savepic

import (
	"context"
	"io/ioutil"
	"log"
	"path"
	"screenshot/action"
	"screenshot/util"

	"github.com/chromedp/chromedp"
	gim "github.com/ozankasikci/go-image-merge"
)

func SaveImg(ctx context.Context, urls []map[string]string, dir string, sleepTime int, buf []byte) ([]*gim.Grid, int, error) {
	util.CreateDirIfNotExist(dir)
	var file string
	grids := []*gim.Grid{}
	count := 0
	for _, x := range urls {
		for k, v := range x {
			if err := chromedp.Run(ctx, action.NetworkTrafficAction(v, 24, sleepTime), util.FullScreenshot(100, &buf)); err != nil {
				return nil, 0, err
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
