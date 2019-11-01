package chatbot

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func Run(user, text string) {
	//取输入参数1和2
	fileName := time.Now().Format("20060102") + ".dingding.log"
	logFile, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Close()
	Log := log.New(logFile, "[Info]", log.Ldate|log.Ltime) // log.Ldate|log.Ltime|log.Lshortfile
	Log.Println("开始发送消息!")
	SendMsg(user, text, Log)
}

//发送消息到钉钉
func SendMsg(user, text string, Log *log.Logger) {
	jsonstring := `{
     "msgtype": "markdown",
     "markdown": {"title":"吾系炒鸡塞牙银!",
"text":"` + text + `"
     },
    "at": {
        "atMobiles": [
            "` + user + `"
        ], 
        "isAtAll": true
    }
 }`
	reader := bytes.NewReader([]byte(jsonstring))
	// cfg, err := goconfig.LoadConfigFile("dingding.conf")
	// if err != nil {
	// Log.SetPrefix("[Err]")
	// Log.Println("读取配置文件失败[config.ini]")
	// return
	// }
	// url, err := cfg.GetValue("setup", "url")
	// url := "https://oapi.dingtalk.com/robot/send?access_token=3c7b7c31b915460f621855e18a298aee782032bfdb27c56d300fa1f422766fbf"
	url := "https://oapi.dingtalk.com/robot/send?access_token=36e452cedd99e028151d3ce6f3b90b9a3994d9ab8a62811c49b4456e3924a92f"
	// if err != nil {
	// Log.SetPrefix("[Err]")
	// Log.Fatalf("无法获取键值（%s）：%s", "url", err)
	// return
	// }
	resp := Post(url, reader)
	Log.SetPrefix("[Info]")
	Log.Printf("消息发送完成,服务器返回内容：%s", string(resp))
}

func Post(url string, reader *bytes.Reader) []byte {
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {

	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	return respBytes
}
