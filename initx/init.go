package initx

import (
	"context"
	"fmt"
	"log"
	"os"
	"screenshot/action"
	"screenshot/loginzbx"
	"screenshot/man"
	"screenshot/savepic"
	"screenshot/util"
	"time"

	"github.com/chromedp/chromedp"
)

var (
	ctx     context.Context
	options []chromedp.ExecAllocatorOption
	version = "0.0.2"
)

var a util.Argv

func init() {
	util.SignalReading(func() {})
	a = util.ParamParser(version)
	man.HelpAsk()
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
		// chromedp.Flag("headless", false),        // 浏览器模式默认为headless, 有需求可将其改为true
		chromedp.Flag("hide-scrollbars", false), // 打开那个啥, 忘记中文名词叫啥来着了
		chromedp.Flag("mute-audio", false),      // 声音操作
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...) // 应用并覆盖选项
}

func Run() {
	// 存图片数据的buffer
	var buf []byte
	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc() // 资源释放
	ctx, cancel := context.WithTimeout(c, time.Second*20)
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel() // 资源释放
	util.SignalReading(cancel)

	if err := chromedp.Run(
		ctx,
		action.SigninAction(a.Host, a.Username, a.Password), // 这个账户密码我就先这么写吧, 一会儿再添加到配置文件
	); err != nil {
		// log.Println(err.Error())
		_ = err
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
