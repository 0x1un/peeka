package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/0x1un/boxes/cmd/attendanceRobot/db"
	"github.com/0x1un/boxes/chatbot"
	"github.com/0x1un/boxes/component/array"
	"github.com/0x1un/boxes/dingtalk/api"
	"log"

	"github.com/subosito/gotenv"
)

var (
	client      *api.DingTalkClient
	GET_TIME    = time.Now().AddDate(0, 0, 0)
	DATE_FORMAT = `2006-01-02`
	allRecord   = new([]api.Schedule)
	conn        = db.Conn
)

func init() {
	gotenv.Load()
	ak := os.Getenv("APPKEY")
	sk := os.Getenv("APPSECRET")
	if ak == "" || sk == "" {
		log.Fatal("请设置APPKEY,APPSECRET到环境变量或.env")
	}
	client = api.NewClient(ak, sk)
}

func main() {
	robot := flag.Bool("robot", false, "是否发送到钉钉机器人")
	whichDay := flag.Int("day", 0, "以当前时间为基准向前/向后推多少天")
	flag.Parse()
	if *whichDay != 0 {
		GET_TIME = time.Now().
			AddDate(0, 0, *whichDay)
	}
	tokens := []string{
		os.Getenv("ROBOT_TOKENS"),
	}
	result := FilterFormatter()
	fmt.Println(result)
	tokens = array.RemoveDuplicateElement(tokens)
	if *robot {
		chatbot.Send(tokens, nil, false, result)
	}
	results, err := client.GetShiftList("2749481918775803", 0)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(results)
	defer conn.Close()
}
