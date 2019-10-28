package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"peeka/cmd/dingtalk/misc"
	"time"
)

type DingTalkClient struct {
	Client      *http.Client
	APPKEY      string
	APPSECRET   string
	BaseURI     string
	AccessToken string
	ATR         AccessTokenResponse // access_token所有返回值
}

type Response struct {
	StatusCode int
	Text       []byte
	Url        string
}

type AccessTokenResponse struct {
	ErrCode     int    `json:"errcode"`
	AccessToken string `json:"access_token"`
	ErrMsg      string `json:"errmsg"`
	ExpiresIn   int    `json:"expires_in"`
}

type Requests interface {
	Get()
	Post()
}

func NewClient(appkey, appsecret string) *DingTalkClient {
	dtc := new(DingTalkClient)
	dtc.Client = &http.Client{
		Timeout: 10 * time.Second,
	}
	dtc.BaseURI = "oapi.dingtalk.com"
	dtc.APPKEY = appkey
	dtc.APPSECRET = appsecret
	return dtc
}

func (d *DingTalkClient) UpdateAccessToken() error {
	params := make(url.Values)
	params.Set("appkey", d.APPKEY)
	params.Set("appsecret", d.APPSECRET)
	text, err := d.Get("gettoken", params)
	if err != nil {
		return err
	}
	var rsp AccessTokenResponse
	if err := json.Unmarshal(text, &rsp); err != nil {
		return err
	}
	if rsp.ErrCode != 0 {
		return errors.New("Failed to get access_token")
	}
	d.ATR = rsp
	d.AccessToken = rsp.AccessToken
	return nil
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

func (d *DingTalkClient) Post(path string, params misc.Params) ([]byte, error) {
	u := &url.URL{
		Scheme:   "https",
		Host:     d.BaseURI,
		Path:     path,
		RawQuery: params.Get("access_token").(string),
	}
	_url := u.String()
	paramsx, err := params.EncodeToJson()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", _url, bytes.NewReader(paramsx))

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
