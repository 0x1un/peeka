package api

import (
	"boxes/internal/dingtalk/misc"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"
)

type UserInfoByCode struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	UserInfo `json:"user_info"`
}

type UserInfo struct {
	Nick    string `json:"nick"`
	Openid  string `json:"openid"`
	Unionid string `json:"unionid"`
}

func (c *DingTalkClient) GetUserInfoByCode(tmpAuthCode, appid, appsecret string) (*UserInfo, error) {
	urlParam := make(url.Values)
	params := make(misc.Data)
	tm := time.Now().UnixNano() / 1e6
	urlParam.Set("access_token", c.AccessToken)
	urlParam.Set("accessKey", appid)
	urlParam.Set("timestamp", fmt.Sprintf("%d", tm))
	urlParam.Set("signature", signatureByQR(tm, appsecret))
	params.Set("tmp_auth_code", tmpAuthCode)

	data, err := c.Post("sns/getuserinfo_bycode", urlParam, params)
	if err != nil {
		return nil, err
	}
	res := new(UserInfoByCode)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	if res.Errcode != 0 {
		return nil, errors.New(res.Errmsg)
	}
	return &res.UserInfo, nil
}

func signatureByQR(timestamp int64, appsecret string) string {
	// sha256 + current timestamp + appSecret => base64encode => urlEncoded

	h := hmac.New(sha256.New, []byte(appsecret))
	_, _ = h.Write([]byte(fmt.Sprintf("%d", timestamp)))
	enc := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return enc

}
