package api

import (
	"encoding/json"
	"net/url"
	"peeka/cmd/dingtalk/misc"
)

type ListSchedule struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Result  struct {
		HasMore   bool `json:"has_more"`
		Schedules []misc.Data
	} `json:"result"`
}

// ListRecord: 打卡记录
type ListRecord struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	RecordResult []struct {
		BaseCheckTime  int64   `json:"baseCheckTime"`
		Id             int64   `json:"id"`
		WorkDate       int64   `json:"workDate"`
		PlanCheckTime  int64   `json:"planCheckTime"`
		PlanId         int64   `json:"planId"`
		GroupId        int64   `json:"groupId"`
		UserCheckTime  int64   `json:"userCheckTime"`
		UserLongitude  float64 `json:"userLongitude"`
		UserAccuracy   float64 `json:"userAccuracy"`
		UserLatitude   float64 `json:"userLatitude"`
		IsLegal        string  `json:"isLegal"`
		UserAddress    string  `json:"userAddress"`
		UserId         string  `json:"userId"`
		CheckType      string  `json:"checkType"`
		TimeResult     string  `json:"timeResult"`
		DeviceId       string  `json:"deviceId"`
		CorpId         string  `json:"corpId"`
		SourceType     string  `json:"sourceType"`
		LocationMethod string  `json:"locationMethod"`
		LocationResult string  `json:"locationResult"`
		ProcInstId     string  `json:"procInstId"`
	}
}

// GetScheduleList: 返回size条结果, offset为偏移量, HasMore为false表示数据已完
// workDate: 只取年月日部分
// offset: 第一次为0, 之后传入offset+size
func (c *DingTalkClient) GetScheduleList(workDate string, offset, size int) (*ListSchedule, error) {
	params := make(misc.Data)
	urlParma := make(url.Values)
	urlParma.Set("access_token", c.AccessToken)
	params.Set("workDate", workDate)
	params.Set("offset", offset)
	params.Set("size", size)
	data, err := Client.Post("topapi/attendance/listschedule", urlParma, params)
	if err != nil {
		return nil, err
	}
	res := new(ListSchedule)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetListRecordDetails: 获取打卡详情, 传入N个用户到数组
func (c *DingTalkClient) GetListRecordDetails(uids []string, chkDateFrom, chkDateTo string) (*ListRecord, error) {
	urlParam := make(url.Values)
	urlParam.Set("access_token", c.AccessToken)
	params := make(misc.Data)
	params.Set("userIds", uids)
	params.Set("checkDateFrom", chkDateFrom)
	params.Set("checkDateTo", chkDateTo)
	data, err := Client.Post("attendance/listRecord", urlParam, params)
	if err != nil {
		return nil, err
	}
	res := new(ListRecord)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetListRecord: 获取打卡结果, limit <= 50, offset初始为0, 后续offset=(offset+limit)
func (c *DingTalkClient) GetListRecord(uids []string, workDateFrom, workDateTo string, offset, limit int) (*ListRecord, error) {
	urlParam := make(url.Values)
	urlParam.Set("access_token", c.AccessToken)
	params := make(misc.Data)
	params.Set("userIdList", uids)
	params.Set("workDateFrom", workDateFrom)
	params.Set("workDateTo", workDateTo)
	params.Set("offset", offset)
	params.Set("limit", limit)
	data, err := Client.Post("attendance/list", urlParam, params)
	if err != nil {
		return nil, err
	}
	res := new(ListRecord)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *DingTalkClient) GetLeaveapproveDuration(uid string, fromDate, toDate string) {

}
