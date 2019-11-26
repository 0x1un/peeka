package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"peeka/cmd/attendanceRobot/db"
	"peeka/internal/chatbot"
	"peeka/internal/component/array"
	"peeka/internal/dingtalk/api"
	"peeka/internal/dingtalk/misc"
	"sort"
	"strconv"
	"time"

	"log"

	"github.com/jinzhu/gorm"
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
	result := Begin()
	fmt.Println(result)
	tokens = array.RemoveDuplicateElement(tokens)
	if *robot {
		chatbot.Send(tokens, nil, false, result)
	}
	defer conn.Close()
}

func Begin() string {
	collections := make(misc.Data)
	// 至于为什么这么写，我懒得写插入数据库了...简单粗暴即可
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
		"15:00:00-22:00:00") // V

	workUsers := Calling()
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

func Calling() []api.Schedule {
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

func InsertRecord(conn *gorm.DB, data interface{}) error {
	var err error
	handler := conn.Begin()
	if value, ok := data.(api.Schedule); ok {
		handler = handler.Table("atten_list")
		value.CreatedAt = GET_TIME.Format(DATE_FORMAT)
		err = handler.Create(value).Error
		if err != nil {
			handler.Rollback()
			return err
		}
	}
	if value, ok := data.(api.UserList); ok {
		handler = handler.Table("dep_users")
		value.CreatedAt = GET_TIME.Format(DATE_FORMAT)
		err = handler.Create(value).Error
		if err != nil {
			handler.Rollback()
			return err
		}
	}
	handler.Commit()
	return nil
}

func UpdateRecord(conn *gorm.DB, data interface{}) error {
	var err error
	handler := conn.Begin()
	if value, ok := data.(api.UserList); ok {
		handler = handler.Table("dep_users")
		value.CreatedAt = GET_TIME.Format(DATE_FORMAT)
		err = handler.Save(value).Error
		if err != nil {
			handler.Rollback()
			return err
		}
	}
	handler.Commit()
	return nil
}

func QueryRecord(conn *gorm.DB, conditions []string, records *[]api.Schedule) error {
	err := conn.
		Table("atten_list").
		Where("checktype = ? AND createdat = ? AND userid in (?)",
			"OnDuty",
			GET_TIME.Format(DATE_FORMAT),
			conditions).
		Find(records).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDepUsers(conn *gorm.DB) *[]api.UserList {
	var err error
	users := new([]api.UserList)
	names := []string{
		"张军", "邹一",
		"唐顺", "唐建",
		"王彪", "李耀",
		"高远", "刘环",
		"陈浩", "尹升俊", "赵鹏辉",
	}
	err = conn.
		Table("dep_users").
		Select("name,userid").
		Where(
			"name in (?)",
			// GET_TIME.Format(DATE_FORMAT),
			names,
		).Scan(users).Error
	if err != nil {
		panic(err)
	}
	if len(*users) != len(names) {
		// 部门id, 调api获取吧
		err = GetAllUserInDepartment("105372678", "0", "100", "")
		if err != nil {
			panic(err)
		}
		return GetDepUsers(conn)
	}
	return users
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
