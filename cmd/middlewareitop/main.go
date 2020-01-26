package main

import (
	"fmt"
	"peeka/cmd/middlewareitop/db"
)

// 釘釘應用程序的agentid
const (
	ITOP_URL = `http://140.246.60.181:8096/itop/webservices/rest.php?version=1.3`
	// ITOP_LOCAL_URL = `http://localhost:8000/webservices/rest.php?version=1.3`
)

func main() {
	request_data, err := NewRestAPIAuthData("admin", "goodluck@123.")
	if err != nil {
		panic(err)
	}

	conn, err := db.NewDBConnect()
	if err != nil {
		panic(err)
	}

	// 从itop中获取所有状态为开启的工单
	resp := FetcheFromITOP(ITOP_URL, request_data)
	for _, v := range resp.Object {
		fmt.Println(v.Filed.Title)
		StoreTicketFromITOP(conn, v.Filed)
	}

	// client := api.NewClient(api.APPKEY, api.APPSECRET)
	// // 发送来自itop的工单
	// if err := SendToProv(client, resp); err != nil {
	// 	panic(err)
	// }
}
