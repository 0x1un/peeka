package main

import (
	"crypto/hmac"
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
				code string
			)
			for k, v := range r.URL.Query() {
				switch k {
				case "code":
					code = v[0]
				}
			}
			dc := api.NewClient("appkey", "appsecret")
			res, err := dc.GetUserInfoByCode(code, "appid", "secretkey")
			if err != nil {
				_, _ = w.Write([]byte(err.Error()))
			}
			// TODO: 获取用户信息后进行进一步的用户详细信息处理
			// get user detail by unionid

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

	h := hmac.New(sha256.New, []byte(appsecret))
	_, _ = h.Write([]byte(fmt.Sprintf("%d", timestamp)))
	enc := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return url.QueryEscape(enc)
}
