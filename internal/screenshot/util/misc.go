package util

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"peeka/internal/screenshot/man"
	"runtime"
	"syscall"

	"github.com/fogleman/gg"
)

type LoginInfo struct {
	Url      string
	Username string
	Password string
}

// LoadJsonConfigToMap: 加载配置文件到map
func LoadJsonConfigToMap(filename string) map[string][]map[string]string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
	}
	var v = make(map[string][]map[string]string)
	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Println(err.Error())
	}

	return v
}

// CreateDirIfNotExist: 如果目录不存在就创建
func CreateDirIfNotExist(dir string) {
	log.Printf("创建目录: %s\n", dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// Argv: 参数解析结构体
type Argv struct {
	Timeout          int
	TotalTimeOut     int
	SangforLoginTime int
	SangforPageTime  int
	Quality          int64
	Help             bool
	Version          bool
	Username         string
	Password         string
	Config           string
	Host             string
	TimeRange        string
}

// ParamParser: 解析参数
func ParamParser(version string) Argv {
	var a Argv

	flag.BoolVar(&a.Help, "h", false, "帮助信息 `help`")
	flag.BoolVar(&a.Version, "v", false, "版本信息 `version`")
	flag.StringVar(&a.Username, "u", "Admin", "zabbix的用户名 `username`")
	flag.StringVar(&a.Password, "p", "zabbix", "zabbix的用户密码 `password`")
	flag.StringVar(&a.Config, "c", "./config.json", "配置文件地址 `config`")
	flag.StringVar(&a.Host, "s", "127.0.0.1", "zabbix服务器地址 `serverName`")
	flag.IntVar(&a.Timeout, "t", 1000, "单个页面抓取等待时间(ms) `waitTime`")
	flag.IntVar(&a.TotalTimeOut, "t-time", 120, "程序总超时时间, =waitTime*抓取数量(s) `TotalTimeOut`")
	flag.IntVar(&a.SangforLoginTime, "loginTime", 3, "深信服登录等待时间, unit(s) `sanfor-login-time`")
	flag.IntVar(&a.SangforPageTime, "pageTime", 20, "深信服进入页面等待加载时间, unit(s) `sanfor-page-time`")
	flag.Int64Var(&a.Quality, "q", 100, "抓取的图片质量 `Quality`")
	flag.StringVar(&a.TimeRange, "time-range", "24h", "设置抓取的时间范围(1h,3h,6h,12h,24h,15m,30m) `TimeRange`")
	flag.Usage = man.Usage
	flag.Parse()
	goos := runtime.GOOS
	arch := runtime.GOARCH

	if a.Help {
		flag.Usage()
		os.Exit(0)
	}
	if a.Version {
		fmt.Printf("MonitorCrawler version v%s %s/%s\n", version, goos, arch)
		os.Exit(0)
	}
	return a
}

// SignalReading: 捕获Ctrl-C信号并释放资源
func SignalReading(cancel func()) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nbye~~")
		cancel()
		os.Exit(0)
	}()
}

func SetTextImg(text, img string, width, height int) error {
	im, err := gg.LoadImage(img)
	if err != nil {
		return err
	}

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("./simkai.ttf", 96); err != nil {
		return err
	}
	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)

	// dc.DrawStringAnchored("阿里-电信深信服", WIDTH/3, HEIGHT/28, 0.5, 0.5)
	dc.DrawStringAnchored("Hello, world!", float64(width/2), float64(height/2), 0.5, 0.5)
	dc.Clip()
	dc.SavePNG("out.png")
	return nil
}
