package main

import (
	"fmt"
	"log"
	"peeka/cmd/dingtalk/api"
)

func main() {
	client := api.Client
	result, err := client.GetListRecord([]string{"156339501623369564"}, "2019-10-28 08:00:00", "2019-10-29 00:00:00", 0, 50)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)
}
