package main

import (
	"github.com/0x1un/boxes/chatbot"
	"github.com/0x1un/boxes/screenshot"
	"bytes"
	"fmt"
	"time"
)

func main() {
	urls := screenshot.Run()
	if !screenshot.CFG.EnableRobot {
		return
	}
	var buffer bytes.Buffer
	buffer.WriteString("## 定时网络巡检播报!\n\n\n")
	for k, v := range urls {
		buffer.WriteString("**" + k + "**\n\n\n")
		buffer.WriteString("![" + k + "]" + "(" + v + ")\n")
	}
	buffer.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	if screenshot.CFG.EnableRobot {
		chatbot.Send(screenshot.CFG.RobotTokens, screenshot.CFG.RobotAtUsers, false, buffer.String())
	}
	fmt.Println(buffer.String())
}
