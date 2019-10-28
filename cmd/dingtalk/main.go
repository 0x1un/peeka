package main

import (
	"fmt"
	"log"
	dingtalk "peeka/cmd/dingtalk/core"
)

func main() {
	l := new(dingtalk.ListRecord)
	err := l.GetListRecordDetails([]string{"156339501623369564"}, "2019-10-28 08:00:00", "2019-10-29 00:00:00")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(l)
}
