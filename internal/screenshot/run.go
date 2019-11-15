package screenshot

import (
	"context"
	"log"
	"net/url"
	"os"
	"path"
	"peeka/internal/screenshot/action"
	"peeka/internal/screenshot/checklogin"
	"time"

	"github.com/0x1un/env"
	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

var (
	options []chromedp.ExecAllocatorOption
	ctx     context.Context
	cancel  context.CancelFunc
	CFG     = ScsConfig{}
)

// init: 初始化一些必要配置
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	if err := env.Parse(&CFG); err != nil {
		log.Fatal("配置文件.env加载错误, 请检查!")
	}
	if !checklogin.ValidateAccountZBX(CFG.ZBXHost, CFG.ZBXUsername, CFG.ZBXPassword) {
		log.Println("帐号或密码验证错误, 请重新指定账户和密码!")
		os.Exit(0)
	}
	ctx = context.Background()
	options = []chromedp.ExecAllocatorOption{
		// chromedp.Flag("headless", false),
		// chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("mute-audio", false),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
		chromedp.WindowSize(1366, 768),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
}

// Run: 总调用接口
func Run() map[string]string {
	var buf []byte
	ctx, cancel = chromedp.NewExecAllocator(ctx, options...)
	ctx, cancel = context.WithTimeout(ctx, time.Second*time.Duration(CFG.TotalTimeOut))
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()
	SignalReading(cancel)

	log.Println("初始化中...")

	if err := chromedp.Run(
		ctx,
		action.SigninAction(CFG.ZBXHost, CFG.ZBXUsername, CFG.ZBXPassword),
	); err != nil {
		log.Println("未在$PATH中找到Chrome/Chromium")
		os.Exit(0)
	}

	CreateDirIfNotExist(CFG.OutputPath)
	remoteFiles := make(map[string]string)
	for name, v := range LoadJsonConfigToMap(CFG.ZBXConfig) {
		saveDir := CFG.OutputPath + "/" + name
		grids, num, err := SaveImg(ctx, v, saveDir, CFG.ZBXTimeRange, CFG.WaitTime, int64(CFG.Quality), CFG.SangforLoginTime, CFG.SangforPageTime, buf)
		if err != nil {
			log.Println(err)
		}
		fname, err := MergeImage(grids, 1, num, name, CFG.IsUpload)
		if err != nil {
			cancel()
			os.Exit(0)
		}
		u, err := url.Parse("http://" + CFG.QiniuBucketURL)
		if err != nil {
			log.Fatal(err)
		}
		u.Path = path.Join(u.Path, fname)
		remoteFiles[name] = u.String()
	}
	log.Println("抓取完成, 进入相关目录查看!")
	return remoteFiles
}
