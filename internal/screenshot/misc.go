package screenshot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fogleman/gg"
)

type LoginInfo struct {
	Url      string
	Username string
	Password string
}

type CommandArgv struct {
	WaitTime         int    `env:"WAIT_TIME"`     // zabbix抓图间隔时间
	TotalTimeOut     int    `env:"TOTAL_TIMEOUT"` // 该程序超时时间，超时后直接退出
	SangforLoginTime int    `env:"SANGFOR_LOGIN_TIME"`
	SangforPageTime  int    `env:"SANGFOR_PAGE_TIME"`
	Quality          int64  `env:"QUALITY"`
	IsUpload         bool   `env:"UPLOAD_QINIU"`
	ZBXUsername      string `env:"ZABBIX_USERNAME"`
	ZBXPassword      string `env:"ZABBIX_PASSWORD"`
	ZBXConfig        string `env:"ZABBIX_CONFIG"`
	ZBXHost          string `env:"ZABBIX_HOST"`
	ZBXTimeRange     string `env:"ZABBIX_TIME_RANGE"`
	EnableRobot      string `env:"ENABLE_ROBOT"`
	FONT_PATH        string `env:"FONT_PATH"`
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

// ParamParser: 解析参数
// func ParamParser(version string) CommandArgv {
// 	var a CommandArgv
//
// 	flag.BoolVar(&a.Help, "h", false, "帮助信息 `help`")
// 	flag.BoolVar(&a.Version, "v", false, "版本信息 `version`")
// 	flag.StringVar(&a.Username, "u", "Admin", "zabbix的用户名 `username`")
// 	flag.StringVar(&a.Password, "p", "zabbix", "zabbix的用户密码 `password`")
// 	flag.StringVar(&a.Config, "c", "./config.json", "配置文件地址 `config`")
// 	flag.StringVar(&a.Host, "s", "127.0.0.1", "zabbix服务器地址 `serverName`")
// 	flag.IntVar(&a.Timeout, "t", 1000, "单个页面抓取等待时间(ms) `waitTime`")
// 	flag.IntVar(&a.TotalTimeOut, "t-time", 120, "程序总超时时间, =waitTime*抓取数量(s) `TotalTimeOut`")
// 	flag.IntVar(&a.SangforLoginTime, "loginTime", 3, "深信服登录等待时间, unit(s) `sanfor-login-time`")
// 	flag.IntVar(&a.SangforPageTime, "pageTime", 20, "深信服进入页面等待加载时间, unit(s) `sanfor-page-time`")
// 	flag.Int64Var(&a.Quality, "q", 100, "抓取的图片质量 `Quality`")
// 	flag.StringVar(&a.TimeRange, "time-range", "24h", "设置抓取的时间范围(1h,3h,6h,12h,24h,15m,30m) `TimeRange`")
// 	flag.StringVar(&a.EnableRobot, "robot", "0", "发送到机器人, 默认0不发送, 1为发送 `enable-robot`")
// 	flag.StringVar(&a.IsUpload, "upload-qiniu", "0", "是否上传到七牛云 `upload-qiniu`")
// 	flag.Usage = man.Usage
// 	flag.Parse()
// 	goos := runtime.GOOS
// 	arch := runtime.GOARCH
//
// 	if a.Help {
// 		flag.Usage()
// 		os.Exit(0)
// 	}
// 	if a.Version {
// 		fmt.Printf("MonitorCrawler version v%s %s/%s\n", version, goos, arch)
// 		os.Exit(0)
// 	}
// 	return a
// }

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
	// 颜色设置成暗红
	dc.SetRGB(139, 0, 0)
	if err := dc.LoadFontFace("./simkai.ttf", 20); err != nil {
		return err
	}
	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)

	// 插入文本到合适的位置
	if strings.Contains(img, "深信") {
		dc.DrawStringAnchored(text, float64(width/3), float64(height/28), 0.5, 0.5)
	} else {
		dc.DrawStringAnchored(text, float64(width/8), float64(height/5), 0.5, 0.5)
	}
	dc.Clip()
	dc.SavePNG(img)
	return nil
}
