package main

import (
	"bytes"
	"fmt"
	"os"
	"peeka/internal/chatbot"
	"peeka/internal/screenshot/run"
)

func main() {
	urls := run.Run()
	var buffer bytes.Buffer
	buffer.WriteString("## 定时网络巡检播报!\n\n\n")
	for k, v := range urls {
		buffer.WriteString("**" + k + "**\n\n\n")
		buffer.WriteString("![" + k + "]" + "(" + v + ")\n")
	}
	fmt.Println(buffer.String())
	if os.Getenv("ENABLE_ROBOT") == "1" {
		chatbot.Run(os.Getenv("ROBOT_TOKEN"), "ALL", buffer.String())
	}
}
