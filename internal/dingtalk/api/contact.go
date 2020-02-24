package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"peeka/pkg/common"
	"strconv"
	"strings"
)

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
	if userinfo.ErrCode != 0 {
		return nil, errors.New(userinfo.ErrMsg)
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

func (c *DingTalkClient) GetUIDbyUnionid(unionid string) (*UidByUnionid, error) {
	// user/getUseridByUnionid
	if unionid == "" {
		return nil, errors.New("unionid为空")
	}
	urlp := make(url.Values)
	urlp.Set("access_token", c.AccessToken)
	urlp.Set("unionid", unionid)
	d, e := c.Get("user/getUseridByUnionid", urlp)
	if e != nil {
		return nil, e
	}
	info := new(UidByUnionid)
	if e := json.Unmarshal(d, info); e != nil {
		return nil, e
	}
	if info.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("%s", info.ErrMsg))
	}
	return info, nil

}

func isFileChanged(filename string) bool {
	hash, err := common.ComputeFileSHA(filename)
	// 如果文件不存在，返回true让其继续访问api并创建缓存文件
	if err != nil {
		return true
	}
	data, err := ioutil.ReadFile(".cache")
	if err != nil {
		return false
	}
	// 文件是否改动
	if strings.Compare(hash, string(data)) == 0 {
		return false
	}
	return true
}
