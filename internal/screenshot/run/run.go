package run

import (
	"context"
	"fmt"
	"log"
	"os"
	"peeka/internal/screenshot/action"
	"peeka/internal/screenshot/loginzbx"
	"peeka/internal/screenshot/savepic"
	"peeka/internal/screenshot/util"
	"time"

	"github.com/chromedp/chromedp"
)

var (
	version = "0.0.3"
	options []chromedp.ExecAllocatorOption
	ctx     context.Context
	cancel  context.CancelFunc
)

var a util.Argv

// init: 初始化一些必要配置
func init() {
	a = util.ParamParser(version)
	if len(os.Args) < 2 {
		fmt.Println("或许你需要指定些什么参数?")
		os.Exit(0)
	}

	if !loginzbx.ValidateAccount(a.Host, a.Username, a.Password) {
		log.Println("帐号或密码验证错误, 请重新指定账户和密码!")
		os.Exit(0)
	}
	ctx = context.Background()
	options = []chromedp.ExecAllocatorOption{
		// chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
}

// Run: 总调用接口
func Run() {
	var buf []byte
	ctx, cancel = chromedp.NewExecAllocator(ctx, options...)
	ctx, cancel = context.WithTimeout(ctx, time.Second*time.Duration(a.TotalTimeOut))
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()
	util.SignalReading(cancel)

	log.Println("初始化中...")

	if err := chromedp.Run(
		ctx,
		action.SigninAction(a.Host, a.Username, a.Password),
	); err != nil {
		log.Println("未在$PATH中找到Chrome/Chromium")
		os.Exit(0)
	}

	for k, v := range util.LoadJsonConfigToMap(a.Config) {
		grids, num, err := savepic.SaveImg(ctx, v, k, a.Timeout, buf)
		if err != nil {
			log.Println(err)
		}
		if err := util.MergeImage(grids, 1, num, k); err != nil {
			cancel()
			os.Exit(0)
		}
	}
	log.Println("抓取完成, 进入相关目录查看!")
}
