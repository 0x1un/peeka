package main

import (
	"bytes"
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
	// chatbot.Run("ALL", buffer.String())
}
