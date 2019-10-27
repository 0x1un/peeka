package loginzbx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Param struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Data struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Param  `json:"params"`
	Id      int    `json:"id"`
}

type ResultData struct {
	JsonRpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	Id      int    `json:"id"`
}

func post(url string, reader *bytes.Reader) []byte {
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Println(err)
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	return respBytes
}

func ValidateAccount(url, username, password string) bool {
	d := Data{
		JsonRpc: "2.0",
		Method:  "user.login",
		Params: Param{
			User:     username,
			Password: password,
		},
		Id: 1,
	}
	data, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}
	read := bytes.NewReader(data)
	result := post("http://"+url+"/api_jsonrpc.php", read)
	ret := &ResultData{}
	if err := json.Unmarshal(result, ret); err != nil {
		log.Println(err)
		return false
	}
	if ret.Result == "" {
		return false
	}
	return true
}
