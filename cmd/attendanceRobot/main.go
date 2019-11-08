package main

import (
	"fmt"
	"os"
	"peeka/internal/dingtalk/api"
	"peeka/internal/dingtalk/misc"
	"time"
)

var (
	client    = api.NewClient(os.Getenv("APPKEY"), os.Getenv("APPSECRET"))
	allRecord = new([]misc.Data)
)

func main() {
	// users, err := client.GetUsersOfDepartmentByDepId("105372678", "0", "100", "")
	GetAllAttendanceResult(time.Now(), 0, 1)
	fmt.Println(len(*allRecord))
}

// 获取所有的考勤信息放到一个[]misc.Data列表当中
func GetAllAttendanceResult(date time.Time, offset, size int) *[]misc.Data {
	schedules, err := client.GetScheduleList(date, offset, size)
	if err != nil {
		return nil
	}
	if schedules.ErrCode != 0 {
		return nil
	}
	// *allRecord = append(*allRecord, schedules.Result.Schedules[0])
	hasMore := schedules.Result.HasMore
	fmt.Println(offset, hasMore)
	if hasMore {
		for _, atten := range schedules.Result.Schedules {
			*allRecord = append(*allRecord, atten)
		}
		offset = offset + size
		return GetAllAttendanceResult(date, offset, size)
	}
	return nil
}
