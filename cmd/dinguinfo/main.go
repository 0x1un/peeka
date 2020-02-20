package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cert_info", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			code := r.URL.Query()["code"][0]
			state := r.URL.Query()["state"][0]
			w.Write([]byte(code + " " + state))
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if fdata, err := ioutil.ReadFile("./resources/index.html"); err != nil {
				log.Fatal(err)
			} else {
				w.Write(fdata)
				return
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8848", nil))

}
