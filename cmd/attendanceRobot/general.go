package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"peeka/internal/dingtalk/api"
	"peeka/internal/dingtalk/misc"
	"sort"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// 对获取到的数据进行清洗
func FilterFormatter() string {
	collections := make(misc.Data)
	// 先写死，考勤组api还没写
	collections.Set("227270072",
		"08:00:00-17:00:00") // A
	collections.Set("361950158",
		"07:00:00-17:00:00") // A1
	collections.Set("362555073",
		"07:00:00-16:00:00") // A2
	collections.Set("226810103",
		"10:00:00-19:00:00") // B1
	collections.Set("144280064",
		"15:00:00-00:00:00") // C1
	collections.Set("353705139",
		"17:00:00-02:00:00") // A1
	collections.Set("385620001",
		"20:00:00-08:00:00") // D1
	collections.Set("272220008",
		"09:00:00-18:00:00") // H
	collections.Set("447095022",
		"12:00:00-21:00:00") // V

	workUsers := getAllAttendanceResult()
	depUsers := GetDepUsers(conn)
	total := make(misc.TData)
	for _, v1 := range workUsers {
		for _, v2 := range *depUsers {
			if v1.Userid == v2.Userid {
				classid := strconv.Itoa(v1.ClassID)
				ret := collections.Get(classid).(string)
				total.Add(ret, v2.Name)
			}
		}
	}
	keys := make([]string, 0, len(total))
	for m := range total {
		keys = append(keys, m)
	}
	if len(keys) < 1 {
		os.Exit(1)
	}
	sort.Strings(keys)
	var buffer bytes.Buffer
	var content string
	title := fmt.Sprintf("# %s日IT到岗时间\n\n",
		GET_TIME.Format(DATE_FORMAT))
	buffer.WriteString(title)
	for _, date := range keys {
		content = fmt.Sprintf("> %s :%s\n\n", date, total[date])
		buffer.WriteString(content)
	}
	buffer.WriteString("**注: 合理联系该时间段在线的IT, 勿打扰休假人员!**")
	return buffer.String()
}

func getAllAttendanceResult() []api.Schedule {
	// TODO: 是否能从数据库中获取用户, 无法获取? 调用api获取并存入数据库
	// TODO: 是否能从数据库中获取当天上班用户, 无法获取? 调用api存入数据库并筛选
	var err error

	filterUsers := make(map[string]interface{})
	for _, user := range *GetDepUsers(conn) {
		filterUsers[user.Name] = user.Userid
	}
	result, err := FilterUsers(conn, filterUsers)
	if err != nil {
		panic(err)
	}
	return *result
}

func FilterUsers(conn *gorm.DB, users map[string]interface{}) (*[]api.Schedule, error) {
	var err error
	var conditions []string // 传入userid列表
	allRecord = new([]api.Schedule)
	count := 0 // 如果重试次数达到3次，直接返回错误!
	for _, uid := range users {
		conditions = append(conditions, uid.(string))
	}
	err = QueryRecord(conn, conditions, allRecord)
	if err != nil {
		return nil, err
	}
	// 每天上班人数至少4人，少于4人api或查询方式肯定有问题! 有待时间的考验就先写死了
	if len(*allRecord) < 4 {
		err = GetAllAttendanceResult(GET_TIME, 0, 1)
		count++
		if count == 3 {
			return nil, errors.New(
				fmt.Sprintf("查询数据库失败, 结果为空\n"),
			)
		}
		return FilterUsers(conn, users)
	}
	return allRecord, nil
}

func GetAllUserInDepartment(depid, offset, size, order string) error {
	users, err := client.
		GetUsersOfDepartmentByDepId(depid, offset, size, order)
	if err != nil {
		return err
	}
	if users.ErrCode != 0 {
		return errors.New(
			fmt.Sprintf(
				"部门用户获取失败: %d:%s",
				users.ErrCode,
				users.ErrMsg,
			))
	}
	for _, user := range users.Userlist {
		if err := UpdateRecord(conn, user); err != nil {
			return err
		}
	}
	return nil
}

// 获取所有的考勤信息存入数据库atten_list表中
func GetAllAttendanceResult(date time.Time, offset, size int) error {
	schedules, err := client.
		GetScheduleList(date, offset, size)
	if err != nil {
		return err
	}
	if schedules.ErrCode != 0 {
		return errors.New(
			fmt.Sprintf("获取考勤列表失败%d:%s",
				schedules.ErrCode,
				schedules.ErrMsg))
	}
	hasMore := schedules.Result.HasMore
	if hasMore {
		for _, atten := range schedules.Result.Schedules {
			err := InsertRecord(conn, atten)
			if err != nil {
				return err
			}
		}
		offset = offset + size
		return GetAllAttendanceResult(date, offset, size)
	}
	return nil
}
