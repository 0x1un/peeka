package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func ViewIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if fdata, err := ioutil.ReadFile("./resources/index.html"); err != nil {
			log.Fatal(err)
		} else {
			_, err = w.Write(fdata)
			CheckErr(err)
		}
	}
}
