package dingtalk

import (
	"encoding/json"
	"log"
	"net/url"
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

type ListSchedule struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
	Result  struct {
		Schedules []misc.Data
		HasMore   bool `json:"has_more"`
	} `json:"result"`
}

func init() {
	if err := Client.UpdateAccessToken(); err != nil {
		log.Println(err)
		return
	}
}

// GetScheduleList: 返回size条结果, offset为偏移量, HasMore为false表示数据已完
func (l *ListSchedule) GetScheduleList(workDate string, offset, size int) error {
	params := make(misc.Data)
	urlParma := make(url.Values)
	urlParma.Set("access_token", Client.AccessToken)
	params.Set("workDate", workDate)
	params.Set("offset", offset)
	params.Set("size", size)
	data, err := Client.Post("topapi/attendance/listschedule", urlParma, params)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, l); err != nil {
		return err
	}
	return nil
}
