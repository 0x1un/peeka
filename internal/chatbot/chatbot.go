package chatbot

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	baseURL = `https://oapi.dingtalk.com/robot/send?access_token=`
)

func Run(token, user, text string) {
	fileName := time.Now().Format("20060102") + ".dingding.log"
	logFile, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Close()
	Log := log.New(logFile, "[Info]", log.Ldate|log.Ltime) // log.Ldate|log.Ltime|log.Lshortfile
	Log.Println("开始发送消息!")
	SendMsg(token, user, text, Log)
}

//发送消息到钉钉
func SendMsg(token, user, text string, Log *log.Logger) {
	jsonstring := `{
     "msgtype": "markdown",
     "markdown": {"title":"皮卡丘!",
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
	resp := Post(baseURL+token, reader)
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
