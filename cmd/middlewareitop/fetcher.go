package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func FetcheFromITOP(url string, data io.Reader) UserReqResponse {
	resp, err := request(http.MethodPost, url, data)
	if err != nil {
		panic(err)
	}
	t := new(UserReqResponse)
	if err := json.Unmarshal(resp, t); err != nil {
		panic(err)
	}
	return *t

}

func request(method, url string, data io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	switch data.(type) {
	case *strings.Reader:
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case *bytes.Reader:
		req.Header.Add("Content-Type", "application/json")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
