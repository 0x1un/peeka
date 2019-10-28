package dingtalk

import (
	"fmt"
	"log"
	"os"
	"peeka/cmd/dingtalk/client"
	"peeka/cmd/dingtalk/misc"
)

var (
	APPKEY    = os.Getenv("APPKEY")
	APPSECRET = os.Getenv("APPSECRET")
)

var (
	Client = client.NewClient(APPKEY, APPSECRET)
)

/**
{
    "result":{
        "schedules":[
            {
                    "plan_id":1,
                    "check_type":"OnDuty",
                    "approve_id":1,
                    "userid":"0001",
                    "class_id":1,
                    "class_setting_id":1,
                    "plan_check_time":"2017-04-11 11:11:11",
                    "group_id":1
            }
        ],
        "has_more":false
    },
    "errmsg":"ok",
    "errcode":0
}
**/

type ListSchedule struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
	Result  struct {
		Schedules []misc.Params
		HasMore   bool `json:"has_more"`
	} `json:"result"`
}

func init() {
	if err := Client.UpdateAccessToken(); err != nil {
		log.Println(err)
		return
	}
}

func (l *ListSchedule) GetScheduleList(workDate string, offset, size int) error {
	params := make(misc.Params)
	params.Set("workDate", workDate)
	params.Set("offset", offset)
	params.Set("size", size)
	result, err := Client.Post("topapi/attendance/listschedule", params)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
