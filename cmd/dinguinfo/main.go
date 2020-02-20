package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"peeka/internal/dingtalk/api"
)

type Data struct {
	TmpAuthCode string `json:"tmp_auth_code"`
}

func main() {
	http.HandleFunc("/cert_info", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var (
				code  string
				state string
			)
			for k, v := range r.URL.Query() {
				switch k {
				case "code":
					code = v[0]
				case "state":
					state = v[0]
				}
			}
			dc := api.NewClient("dingkimljpj2mycouknv", "31eyI8v57QWhd5ub4L2-xw3m4z4SpI5OWyN0KhGdQqYpCqyw_DUeR2gP2FEQ66fP")

		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if fdata, err := ioutil.ReadFile("./resources/index.html"); err != nil {
				log.Fatal(err)
			} else {
				_, err = w.Write(fdata)
				checkErr(err)
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8848", nil))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func signatureByQR(timestamp int64, appsecret string) string {
	// sha256 + current timestamp + appSecret => base64encode => urlEncoded
	sum := sha256.Sum256([]byte(fmt.Sprintf("%d%s", timestamp, appsecret)))
	enc := base64.StdEncoding.EncodeToString(sum[:])
	return url.QueryEscape(enc)
}
