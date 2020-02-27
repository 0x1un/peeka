package main

import (
	"boxes/internal/dingtalk/api"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Printf("%s 来访\n", r.RemoteAddr)
		var (
			code string
		)
		for k, v := range r.URL.Query() {
			switch k {
			case "code":
				code = v[0]
			}
		}
		dc := api.NewClient("appkey", "appsecret")
		res, err := dc.GetUserInfoByCode(code, "appid", "appsecret")
		fmt.Println(code)
		if err != nil {
			fmt.Println("你不能通过链接直接访问该页面!")
			return
		}
		// TODO: 获取用户信息后进行进一步的用户详细信息处理
		// get user detail by unionid
		uid, err := dc.GetUIDbyUnionid(res.Unionid)
		if err != nil {
			panic(err)
		}
		// 在获得了uid后，使用它来获取用户的详细信息
		info, err := dc.GetUserInfoDetailsByUid(uid.Userid, "")
		if err != nil {
			panic(err)
		}
		i := struct {
			Name       string
			Department string
			Mobile     string
			Avatar     string
			HiredDate  string
		}{
			Name:       info.Name,
			Department: info.Roles[len(info.Roles)-1].Name,
			Mobile:     info.Mobile,
			Avatar:     info.Avatar,
			HiredDate:  time.Unix(0, info.HiredDate*int64(time.Millisecond)).Format("2006-01-02"),
		}
		tpl, err := template.ParseFiles("./resources/info.html")
		if err != nil {
			panic(err)
		}
		err = tpl.Execute(w, i)
		if err != nil {
			panic(err)
		}
	}
}
