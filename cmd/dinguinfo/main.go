package main

import (
	"boxes/internal/logger"
	"log"
	"net/http"
)

type Data struct {
	TmpAuthCode string `json:"tmp_auth_code"`
}

func main() {
	logger.Info.Println("错误")
	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/cert_info", ViewProfile)

	http.HandleFunc("/", ViewIndex)
	log.Fatal(http.ListenAndServe(":8848", nil))

}
