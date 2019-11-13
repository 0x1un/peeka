package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"peeka/internal/dingtalk/misc"
	"time"
)

var (
	APPKEY    = os.Getenv("APPKEY")
	APPSECRET = os.Getenv("APPSECRET")
	Client    = NewClient(APPKEY, APPSECRET)
)

type ErrResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type Response struct {
	StatusCode int
	Text       []byte
	Url        string
}

type AccessTokenResponse struct {
	ErrResponse
	AccessToken string `json:"access_token"`
	Expires     int    `json:"expires_in"`
	Created     int64
}

// api主结构, 所有的api都围绕此结构体
type DingTalkClient struct {
	Client      *http.Client
	Parameters  url.Values
	Data        misc.Data
	APPKEY      string
	APPSECRET   string
	BaseURI     string
	AccessToken string
}

func (d *DingTalkClient) UpdateAccessToken() error {
	var rsp AccessTokenResponse
	if rp, err := ValidateToken(); err == nil {
		d.AccessToken = rp.AccessToken
		return nil
	}
	params := make(url.Values)
	params.Set("appkey", d.APPKEY)
	params.Set("appsecret", d.APPSECRET)
	text, err := d.Get("gettoken", params)
	if err != nil {
		return err
	}
	if rsp.ErrCode != 0 {
		return errors.New("Failed to get access_token")
	}
	rsp.Created = time.Now().Unix()
	if err := json.Unmarshal(text, &rsp); err != nil {
		return err
	}
	data, err := json.Marshal(rsp)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(".token.json", data, 0666); err != nil {
		return err
	}
	return nil
}

func NewClient(appkey, appsecret string) *DingTalkClient {
	dtc := new(DingTalkClient)
	dtc.Client = &http.Client{
		Timeout: 20 * time.Second,
	}
	dtc.BaseURI = "oapi.dingtalk.com"
	dtc.APPKEY = appkey
	dtc.APPSECRET = appsecret
	err := dtc.UpdateAccessToken()
	if err != nil {
		log.Fatalf("获取access_token失败: %s", err.Error())
	}
	return dtc
}

// check token from local
func ValidateToken() (*AccessTokenResponse, error) {
	jsonstr, err := ioutil.ReadFile(".token.json")
	if err != nil {
		return nil, err
	}
	var rsp AccessTokenResponse
	err = json.Unmarshal(jsonstr, &rsp)
	if err != nil {
		return nil, err
	}
	if (time.Now().Unix() - int64(rsp.Expires)) >= rsp.Created {
		return nil, errors.New("token已经过期了")
	}
	return &rsp, nil
}

func (d *DingTalkClient) Get(path string, params url.Values) ([]byte, error) {
	u := &url.URL{
		Scheme:   "https",
		Host:     d.BaseURI,
		Path:     path,
		RawQuery: params.Encode(),
	}
	_url := u.String()
	req, err := http.NewRequest("GET", _url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}
	text, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return text, nil
}

func (d *DingTalkClient) Post(path string, urlP url.Values, params misc.Data) ([]byte, error) {
	u := &url.URL{
		Scheme:   "https",
		Host:     d.BaseURI,
		Path:     path,
		RawQuery: urlP.Encode(),
		// RawQuery: d.Parameters.Encode(),
	}
	_url := u.String()
	paramsx, err := params.EncodeToJson()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", _url, bytes.NewReader(paramsx))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}
	resp, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}
	text, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return text, nil
}
