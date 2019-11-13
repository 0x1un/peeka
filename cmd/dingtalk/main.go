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
	userids := []string{
		"2749481918775803",
	}
	data, err := client.GetListRecord(userids, "2019-11-11 08:00:00", "2019-11-12 00:00:00", 0, 20)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}
