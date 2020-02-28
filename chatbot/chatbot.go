package chatbot

import (
	network "github.com/0x1un/boxes/component/net"
	"bytes"
	"encoding/json"
	"log"
	"os"
)

const (
	baseURL  = `https://oapi.dingtalk.com/robot/send?access_token=`
	fileName = "chatbot.log"
)

type Message struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

func Send(tokens, atUsers []string, notifyAll bool, text string) {
	logFile, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Close()
	Log := log.New(logFile, "[Info]", log.Ldate|log.Ltime) // log.Ldate|log.Ltime|log.Lshortfile
	Log.Println("开始发送消息!")
	msg := new(Message)
	msg.MsgType = "markdown"
	msg.Markdown.Title = "[钉钉红包]恭喜发财 大吉大利!"
	msg.Markdown.Text = text
	msg.At.AtMobiles = atUsers
	msg.At.IsAtAll = notifyAll
	msgs, err := json.Marshal(msg)
	if err != nil {
		Log.Fatal(err)
	}
	for _, tk := range tokens {
		fillMsgAndSent(tk, msgs, Log)
	}
}

//发送消息到钉钉
func fillMsgAndSent(token string, msg []byte, Log *log.Logger) {
	reader := bytes.NewReader(msg)
	resp := network.Post(baseURL+token, reader)
	Log.SetPrefix("[Info]")
	Log.Printf("消息发送完成,服务器返回内容：%s", string(resp))
}
