package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"peeka/cmd/dingtalk/misc"
	"strconv"
)

type UserId struct {
	ErrResponse
	Userid string `json:"userid"`
}

type UserList struct {
	Userid string `json:"userid"`
	Name   string `json:"name"`
}

type UsersOfDepartment struct {
	ErrResponse
	HasMore  bool       `json:"hasMore"`
	Userlist []UserList `json:"userlist"`
}

type UserInfoDetails struct {
	ErrResponse
	UserId
	Unionid         string `json:"unionid"`
	Remark          string `json:"remark"`
	IsLeaderInDepts string `json:"isLeaderInDepts"`
	IsBoss          bool   `json:"isBoss"`
	HiredDate       int64  `json:"hiredDate"`
	IsSenior        bool   `json:"isSenior"`
	Tel             string `json:"tel"`
	Department      []int  `json:"department"`
	WorkPlace       string `json:"workPlace"`
	Email           string `json:"email"`
	OrderInDepts    string `json:"orderInDepts"`
	Mobile          string `json:"mobile"`
	Active          bool   `json:"active"`
	Avatar          string `json:"avatar"`
	IsAdmin         bool   `json:"isAdmin"`
	IsHide          bool   `json:"isHide"`
	Jobnumber       string `json:"jobnumber"`
	Name            string `json:"name"`
	Extattr         struct {
	} `json:"extattr"`
	StateCode string `json:"stateCode"`
	Position  string `json:"position"`
	Roles     []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		GroupName string `json:"groupName"`
	} `json:"roles"`
}

// userid为用户的id, 企业内唯一， lang为通讯录语言(默认为zh_CN)
// 如果lang为空("")默认为zh_CN
func (c *DingTalkClient) GetUserInfoDetailsByUid(userid, lang string) (*UserInfoDetails, error) {
	if userid == "" {
		return nil, errors.New("请传入正确的userid")
	}
	urlParam := make(url.Values)
	urlParam.Set("access_token", c.AccessToken)
	urlParam.Set("userid", userid)
	if lang == "" {
		lang = "zh_CN"
	}
	urlParam.Set("lang", lang)
	data, err := c.Get("user/get", urlParam)
	if err != nil {
		return nil, err
	}
	userinfo := new(UserInfoDetails)
	if err := json.Unmarshal(data, userinfo); err != nil {
		return nil, err
	}
	return userinfo, nil
}

func (c *DingTalkClient) GetUserIdByMobile(mobile string) (*UserId, error) {
	if mobile == "" {
		return nil, errors.New("传入的手机号不规范")
	}
	urlParam := make(url.Values)
	urlParam.Set("access_token", c.AccessToken)
	urlParam.Set("mobile", mobile)
	data, err := c.Get("user/get_by_mobile", urlParam)
	if err != nil {
		return nil, err
	}
	userid := new(UserId)
	if err := json.Unmarshal(data, userid); err != nil {
		return nil, err
	}
	return userid, nil
}

// 获取一个部门下的所有摘要用户信息
// offset与size配合使用, size不能大于100
// order为排序规则, 默认不传为加入部门时间的升序
func (c *DingTalkClient) GetUsersOfDepartmentByDepId(depId, offset, size, order string) (*UsersOfDepartment, error) {
	size_, err := strconv.Atoi(size)
	if err != nil {
		return nil, err
	}
	if size_ > 100 {
		return nil, errors.New("size不能大于100!")
	}
	if order == "" {
		order = "entry_asc"
	}
	urlParam := make(url.Values)
	urlParam.Set("access_token", c.AccessToken)
	urlParam.Set("department_id", depId)
	urlParam.Set("offset", offset)
	urlParam.Set("size", size)
	urlParam.Set("order", order)
	data, err := c.Get("user/simplelist", urlParam)
	if err != nil {
		return nil, err
	}
	users := new(UsersOfDepartment)
	if err := json.Unmarshal(data, users); err != nil {
		return nil, err
	}
	return users, nil
}

// 将用户信息写入到文件方便以后快速读取
func (c *UsersOfDepartment) WriteToCacheFile() error {
	if c.ErrCode != 0 {
		return errors.New("获取部门用户失败!")
	}
	buffer := new(bytes.Buffer)
	for _, v := range c.Userlist {
		buffer.WriteString(v.Name + ":" + v.Userid + "\n")
	}
	err := ioutil.WriteFile(".users", buffer.Bytes(), 0666)
	if err != nil {
		return err
	}
	return nil
}

// 从cache文件中读取用户信息
func (c *UsersOfDepartment) ReadFromCacheFile() error {
	if !misc.IsExist(".users") {
		return errors.New(".users文件没有找到")
	}
	return nil
}
