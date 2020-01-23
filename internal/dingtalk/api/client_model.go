package api

import (
	"net/http"
	"net/url"
	"peeka/internal/dingtalk/misc"
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
