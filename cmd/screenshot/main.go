package main

import (
	"bytes"
	"fmt"
	"peeka/internal/chatbot"
	"peeka/internal/screenshot"
	"time"
)

func main() {
	urls := screenshot.Run()
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
