package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"peeka/internal/dingtalk/misc"
	"reflect"
	"time"
)

// 使用gorm标签，如有特殊用途，自行修改tag
type Schedule struct {
	PlanID         int    `json:"plan_id" gorm:"column:planid"`
	CheckType      string `json:"check_type" gorm:"column:checktype"`
	ApproveID      int    `json:"approve_id" gorm:"column:approveid"`
	Userid         string `json:"userid" gorm:"column:userid"`
	ClassID        int    `json:"class_id" gorm:"column:classid"`
	ClassSettingID int    `json:"class_setting_id" gorm:"column:classsettingid"`
	PlanCheckTime  string `json:"plan_check_time" gorm:"column:planchecktime"`
	GroupID        int    `json:"group_id" gorm:"column:groupid"`
	CreatedAt      string `gorm:"column:createdat"`
	UserName       string `gorm:"-"`
}

type ListSchedule struct {
	ErrResponse
	Result struct {
		HasMore bool `json:"has_more"`
		// Schedules []misc.Data
		Schedules []Schedule `json:"schedules"`
	} `json:"result"`
}

// 获取考勤的班次摘要信息
type ShiftList struct {
	ErrResponse
	Result struct {
		HasMore bool `json:"has_more"`
		Cursor  int  `json:"cursor"`
		Result  []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"result"`
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
func (c *DingTalkClient) GetScheduleList(workDate time.Time, offset, size int) (*ListSchedule, error) {
	if size > 200 {
		return nil, errors.New("size不能大于200!")
	}
	wrkDate := workDate.Format("2006-01-02")
	params := make(misc.Data)
	urlParma := make(url.Values)
	urlParma.Set("access_token", c.AccessToken)
	params.Set("workDate", wrkDate)
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

// 获取考勤班次摘要信息
func (c *DingTalkClient) GetShiftList(opuids string, cursor int) (*ShiftList, error) {
	if err := checkParameter(opuids, cursor); err != nil {
		return nil, err
	}
	return nil, nil
}

// 检查参数是否为空
func checkParameter(args ...interface{}) error {
	if len(args) == 0 {
		return errors.New("你必须传入一些参数")
	}
	for _, v := range args {
		if !reflect.ValueOf(v).IsValid() {
			return errors.New(fmt.Sprintf("参数* %s *为空!", reflect.TypeOf(v).String()))
		}
	}
	return nil
}
