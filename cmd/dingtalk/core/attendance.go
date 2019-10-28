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

// ListRecord: 打卡记录
type ListRecord struct {
	ErrMsg       string `json:"errmsg"`
	ErrCode      int    `json:"errcode"`
	RecordResult []struct {
		IsLegal        string  `json:"isLegal"`
		BaseCheckTime  int64   `json:"baseCheckTime"`
		Id             int64   `json:"id"`
		UserAddress    string  `json:"userAddress"`
		UserId         string  `json:"userId"`
		CheckType      string  `json:"checkType"`
		TimeResult     string  `json:"timeResult"`
		DeviceId       string  `json:"deviceId"`
		CorpId         string  `json:"corpId"`
		SourceType     string  `json:"sourceType"`
		WorkDate       int64   `json:"workDate"`
		PlanCheckTime  int64   `json:"planCheckTime"`
		LocationMethod string  `json:"locationMethod"`
		LocationResult string  `json:"locationResult"`
		UserLongitude  float64 `json:"userLongitude"`
		PlanId         int64   `json:"planId"`
		GroupId        int64   `json:"groupId"`
		UserAccuracy   int     `json:"userAccuracy"`
		UserCheckTime  int64   `json:"userCheckTime"`
		UserLatitude   float64 `json:"userLatitude"`
		ProcInstId     string  `json:"procInstId"`
	}
}

func init() {
	if err := Client.UpdateAccessToken(); err != nil {
		log.Println(err)
		return
	}
}

// GetScheduleList: 返回size条结果, offset为偏移量, HasMore为false表示数据已完
// workDate: 只取年月日部分
// offset: 第一次为0, 之后传入offset+size
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

// GetListRecord: 获取打卡记录, 传入N个用户id查询打卡详情
func (l *ListRecord) GetListRecordDetails(uids []string, chkDateFrom, chkDateTo string) error {
	urlParam := make(url.Values)
	urlParam.Set("access_token", Client.AccessToken)
	params := make(misc.Data)
	params.Set("userIds", uids)
	params.Set("checkDateFrom", chkDateFrom)
	params.Set("checkDateTo", chkDateTo)
	data, err := Client.Post("attendance/listRecord", urlParam, params)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, l); err != nil {
		return err
	}
	return nil
}
