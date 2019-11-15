// merge images and screenshot
package screenshot

import (
	"context"
	"errors"
	"image/png"
	"math"
	"os"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	gim "github.com/ozankasikci/go-image-merge"
)

// MergeImage: 合并图片, 合并规则由(x,y)决定
// 返回参数需要添加一个string, 作用为返回图床链接
func MergeImage(grids []*gim.Grid, x, y int, filename, upload string) (string, error) {
	if len(grids) == 0 {
		return "", errors.New("No pictures..")
	}

	rgba, err := gim.New(grids, x, y).Merge()
	if err != nil {
		return "", err
	}
	// save the output to jpg or png
	file, err := os.Create(filename + ".png")
	if err != nil {
		return "", err
	}
	if err := png.Encode(file, rgba); err != nil {
		return "", err
	}
	// TODO: 这里需要将图片上传到图床
	fname, err := PostFileToStorage(file.Name(), upload)
	if fname == "" && err == nil {
		return "", nil
	}
	return fname, nil
}

// FullScreenshot: 截取完整图片
func FullScreenshot(quality int64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
			if err != nil {
				return err
			}
			// fmt.Println(contentSize.X, contentSize.Y, contentSize.Width, contentSize.Height)

			width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

			// force viewport emulation
			err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
				WithScreenOrientation(&emulation.ScreenOrientation{
					Type:  emulation.OrientationTypePortraitPrimary,
					Angle: 0,
				}).
				Do(ctx)
			if err != nil {
				return err
			}

			*res, err = page.CaptureScreenshot().
				WithQuality(quality).
				WithClip(&page.Viewport{
					X:      contentSize.X,
					Y:      contentSize.Y,
					Width:  contentSize.Width,
					Height: contentSize.Height,
					Scale:  1,
				}).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}
}

// 对指定元素进行截图, 小面积截图
func ElementScreenshot(sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
	}
}
