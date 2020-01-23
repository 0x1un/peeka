package main

import (
	"encoding/json"
	"net/url"
	"strings"
)

// UserRequest structure
type Base struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Fileds struct {
	RequestType            string `json:"request_type"`
	ServiceSubcategoryName string `json:"servicesubcategory_name"`
	Urgency                string `json:"urgency"`
	Origin                 string `json:"origin"`
	CallerIdFriendlyName   string `json:"caller_id_friendlyname"`
	Impact                 string `json:"impact"`
	Title                  string `json:"title"`
	Description            string `json:"description"`
}

type ResponseContent struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Class   string `json:"class"`
	Key     string `json:"key"`
	Filed   Fileds `json:"fields"`
}

// UserRequest返回的响应内容
type UserReqResponse struct {
	Base
	Object map[string]ResponseContent `json:"objects"`
}

// Request api data struct
type RequestData struct {
	Operation    string `json:"operation"`
	Class        string `json:"class"`
	Key          string `json:"key"`
	OutPutFields string `json:"output_fields"`
}

func NewRestAPIAuthData(auth_user, auth_pwd string) (*strings.Reader, error) {
	req_data := make(url.Values)
	json_data := new(RequestData)
	json_data.Operation = "core/get"
	json_data.Class = "UserRequest"
	json_data.Key = "SELECT UserRequest WHERE operational_status = \"ongoing\""
	json_data.OutPutFields = "request_type,servicesubcategory_name,urgency,origin,caller_id_friendlyname,impact,title,description,contacts_list"
	data, err := json.Marshal(json_data)
	if err != nil {
		return nil, err
	}
	req_data.Add("auth_user", auth_user)
	req_data.Add("auth_pwd", auth_pwd)
	req_data.Add("json_data", string(data))
	return strings.NewReader(req_data.Encode()), nil
}
