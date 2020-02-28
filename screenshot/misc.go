package screenshot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fogleman/gg"
)

// screenshot config
type ScsConfig struct {
	EnableRobot      bool     `env:"ENABLE_ROBOT"`
	SangforPageTime  int      `env:"SANGFOR_PAGE_TIME"`
	Quality          int      `env:"QUALITY"`
	WaitTime         int      `env:"WAIT_TIME"`     // zabbix抓图间隔时间
	TotalTimeOut     int      `env:"TOTAL_TIMEOUT"` // 该程序超时时间，超时后直接退出
	SangforLoginTime int      `env:"SANGFOR_LOGIN_TIME"`
	ZBXUsername      string   `env:"ZABBIX_USERNAME"`
	ZBXPassword      string   `env:"ZABBIX_PASSWORD"`
	ZBXConfig        string   `env:"ZABBIX_CONFIG"`
	ZBXHost          string   `env:"ZABBIX_HOST"`
	ZBXTimeRange     string   `env:"ZABBIX_TIME_RANGE"`
	QiniuAK          string   `env:"QINIU_ACCESS_KEY"`
	QiniuSK          string   `env:"QINIU_SECRET_KEY"`
	QiniuBucket      string   `env:"QINIU_BUCKET"`
	QiniuBucketURL   string   `env:"QINIU_BUCKET_URL"`
	FontPath         string   `env:"FONT_PATH"`
	OutputPath       string   `env:"OUTPUT_PATH"` // 抓取下来的图片保存位置
	SangforUsername  string   `env:"SANGFOR_USERNAME"`
	SangforPassword  string   `env:"SANGFOR_PASSWORD"`
	RobotAtUsers     []string `env:"ROBOT_AT_USERS" envSeparator:";"`
	RobotTokens      []string `env:"ROBOT_TOKENS" envSeparator:";"`
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

// SignalReading: 捕获Ctrl-C信号并释放资源
func SignalReading(cancel func()) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
		print("\n已退出\n")
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
	if err := dc.LoadFontFace(CFG.FontPath, 20); err != nil {
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
