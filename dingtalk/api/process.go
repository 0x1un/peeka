package api

import (
	network "github.com/0x1un/boxes/component/net"
	"bytes"
	"encoding/json"
	"os"
	"strconv"
)

// Methods

func (f *FormValues) add(key, value string) {
	*f = append(*f, FormComponentValuesVo{
		Name:  key,
		Value: value,
	})
}

// 城市，臺席號，域帳號，聯繫方式，故障類型，故障範圍，故障現象
// 這裏的工單爲現在自用的，具體實現靈活的用map或struct實現即可。
func FillForm(city, local, ad_user, contact, fault_type, fault_range, faults string) (formValues FormValues) {
	formValues.add("城市", city)
	formValues.add("台席号", local)
	formValues.add("域账号", ad_user)
	formValues.add("联系方式", contact)
	formValues.add("故障类型", fault_type)
	formValues.add("故障范围", fault_range)
	formValues.add("故障现象", faults)
	return formValues
}

func (self *DingTalkClient) SendProcessForTest(formComponent FormValues) (*CreateProcessInstanceResp, error) {

	var (
		processCreator CreateProcessInstanceReq
		formValues     FormValues
	)
	agentid, err := strconv.Atoi(os.Getenv("APP_AGENT_ID"))
	senderdep_id, err := strconv.Atoi(os.Getenv("DEPID"))
	if err != nil {
		return nil, err
	}
	processCreator.AgentId = int64(agentid)
	processCreator.ProcessCode = os.Getenv("APROVID")
	processCreator.OriginatorUserId = os.Getenv("USERID")
	processCreator.DeptId = int64(senderdep_id)
	formValues = formComponent

	processCreator.FormComponentValues = formValues

	data, err := json.Marshal(processCreator)
	if err != nil {
		return nil, err
	}
	d := bytes.NewReader(data)
	url := "https://oapi.dingtalk.com/topapi/processinstance/create?access_token=" + self.AccessToken
	result := network.Post(url, d)

	var processResp *CreateProcessInstanceResp
	if err := json.Unmarshal(result, &processResp); err != nil {
		return nil, err
	}
	return processResp, nil
}
