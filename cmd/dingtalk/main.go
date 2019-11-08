package main

import (
	"fmt"
	"log"
	"peeka/internal/dingtalk/api"
)

// 2749481918775803

var (
	unionid = "YA2WZD0MiSokGNqsnbfoL1QiEiE"
)

func main() {
	client := api.Client
	// userids := []string{
	// 	"2749481918775803",
	// }

	// result, err := client.GetListRecord(userids, "2019-10-28 08:00:00", "2019-10-29 00:00:00", 0, 50)
	// result, err := client.GetUserInfoDetailsByUid(userids[0], "")
	// result, err := client.GetUserIdByMobile("17608035126")
	result, err := client.GetUsersOfDepartmentByDepId("105372678", "0", "100", "")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)
}
