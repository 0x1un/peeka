package api

import (
	"encoding/json"
	"errors"
	"net/url"
	"peeka/cmd/dingtalk/misc"
)

type ListSchedule struct {
	ErrResponse
	Result struct {
		HasMore   bool `json:"has_more"`
		Schedules []misc.Data
	} `json:"result"`
}

// ListRecord: 打卡记录
type ListRecord struct {
	ErrResponse
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

// 考勤组摘要结构
type AttdGroup struct {
	ErrResponse
	Result []struct {
		Name string `json:"Name"`
		Id   int    `json:"id"`
	} `json:"result"`
}

// GetScheduleList: 返回size条结果, offset为偏移量, HasMore为false表示数据已完
// workDate: 只取年月日部分
// offset: 第一次为0, 之后传入offset+size
func (c *DingTalkClient) GetScheduleList(workDate string, offset, size int) (*ListSchedule, error) {
	if size > 200 {
		return nil, errors.New("size不能大于200!")
	}
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
	if len(uids) == 0 {
		return nil, errors.New("请传入userid!")
	}
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
	if len(uids) == 0 {
		return nil, errors.New("请传入userid!")
	}
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

// 获取请假时长
func (c *DingTalkClient) GetLeaveapproveDuration(uid string, fromDate, toDate string) {

}

// 搜索考勤组摘要
// 按考勤组名称模糊搜索，获取考勤组的摘要信息
// opUid为操作者的userid
func (c *DingTalkClient) GetAttendanceGroup(opUid, groupName string) (*AttdGroup, error) {
	if opUid == "" || groupName == "" {
		return nil, errors.New("请传入有效的op_user_id或group name")
	}
	urlParam := make(url.Values)
	urlParam.Set("access_token", c.AccessToken)
	params := make(misc.Data)
	params.Set("op_user_id", opUid)
	params.Set("group_name", groupName)
	data, err := Client.Post("topapi/attendance/group/search", urlParam, params)
	if err != nil {
		return nil, err
	}
	res := new(AttdGroup)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}
