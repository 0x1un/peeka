package main

import (
	"bytes"
	"fmt"
	"os"
<<<<<<< HEAD

=======
>>>>>>> 7d58b8f4feafc60ba892e4a2ecc19ab21ad74c27
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
<<<<<<< HEAD
	chatbot.Run(os.Getenv("ROBOT_TOKEN"), "ALL", buffer.String())
=======
	if os.Getenv("ENABLE_ROBOT") == "1" {
		chatbot.Run(os.Getenv("ROBOT_TOKEN"), "ALL", buffer.String())
	}
>>>>>>> 7d58b8f4feafc60ba892e4a2ecc19ab21ad74c27
}
