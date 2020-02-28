package api

import (
	"github.com/0x1un/boxes/dingtalk/misc"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"time"
)

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

// 根据用户id获取考勤班次摘要信息
func (c *DingTalkClient) GetShiftList(opuids string, cursor int) (*GroupMinimalismList, error) {
	if err := checkParameter(opuids, cursor); err != nil {
		return nil, err
	}
	urlParam := make(url.Values)
	urlParam.Set("access_token", c.AccessToken)
	params := make(misc.Data)
	params.Set("op_user_id", opuids)
	params.Set("cursor", cursor)
	data, err := Client.Post("topapi/attendance/group/minimalism/list", urlParam, params)
	if err != nil {
		return nil, err
	}
	res := new(GroupMinimalismList)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

// 获取指定考勤组的详细信息
func (c *DingTalkClient) GetSpecShiftDetail(opuid, grpName string) (*AttenGroup, error) {
	if err := checkParameter(opuid, grpName); err != nil {
		return nil, err
	}
	urlParam := make(url.Values)
	urlParam.Set("access_token", c.AccessToken)
	params := make(misc.Data)
	params.Set("op_user_id", opuid)
	params.Set("group_id", grpName)
	data, err := Client.Post("topapi/attendance/group/query", urlParam, params)
	if err != nil {
		return nil, err
	}
	res := new(AttenGroup)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
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
