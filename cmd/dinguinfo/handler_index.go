package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func ViewIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("./resources/handler_index.html")
		if err != nil {
			panic(err)
		}
		params := struct{
			CallBackUrl string
		}{
			CallBackUrl: "http://localhost:8848/cert_info"
		}
		err = tpl.Execute(w, params)
		if err != nil {
			panic(err)
		}
	}
}
