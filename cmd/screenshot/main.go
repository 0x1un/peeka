package main

import (
	"bytes"
	"os"
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
	tokens := []string{
		os.Getenv("ROBOT_TOKEN"),
	}
	if os.Getenv("ENABLE_ROBOT") == "1" {
		chatbot.Send(tokens, []string{}, false, buffer.String())
	}
}
