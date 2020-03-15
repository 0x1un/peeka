package api

// Request Data

type FormValues []FormComponentValuesVo

type CreateProcessInstanceReq struct {
	AgentId          int64  `json:"agent_id"`
	ProcessCode      string `json:"process_code"`
	OriginatorUserId string `json:"originator_user_id"`
	DeptId           int64  `json:"dept_id"`
	// Approvers           string                      `json:"approvers"`
	// ApproversV2         []ProcessInstanceApproverVo `json:"approvers_v2"`
	// CcList              string                      `json:"cc_list"`
	// CcPosition          string                      `json:"cc_position"`
	FormComponentValues FormValues `json:"form_component_values"`
}

type ProcessInstanceApproverVo struct {
	UserIds        []string `json:"user_ids"`
	TaskActionType string   `json:"task_action_type"`
}

type FormComponentValuesVo struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	ExtValue string `json:"ext_value"`
}

type CreateProcessInstanceResp struct {
	BaseResp
	ProcessInstanceId string `json:"process_instance_id"`
}

type BaseResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
type OperationRecords []struct {
	Userid          string `json:"userid"`
	Date            string `json:"date"`
	OperationType   string `json:"operation_type"`
	OperationResult string `json:"operation_result"`
	Remark          string `json:"remark"`
}

type Tasks []struct {
	Userid     string `json:"userid"`
	TaskStatus string `json:"task_status"`
	TaskResult string `json:"task_result"`
	CreateTime string `json:"create_time"`
	FinishTime string `json:"finish_time"`
	Taskid     string `json:"taskid"`
}

type ProcessInstance struct {
	Title                      string `json:"title"`
	CreateTime                 string `json:"create_time"`
	FinishTime                 string `json:"finish_time"`
	OriginatorUserid           string `json:"originator_userid"`
	OriginatorDeptID           string `json:"originator_dept_id"`
	Status                     string `json:"status"`
	CcUserids                  string `json:"cc_userids"`
	Result                     string `json:"result"`
	BusinessID                 string `json:"business_id"`
	OriginatorDeptName         string `json:"originator_dept_name"`
	BizAction                  string `json:"biz_action"`
	FormComponentValuesVo      `json:"form_component_values"`
	OperationRecords           `json:"operation_records"`
	Tasks                      `json:"tasks"`
	AttachedProcessInstanceIds []interface{} `json:"attached_process_instance_ids"`
}

type ProcessInstanceDetail struct {
	BaseResp
	ProcessInstance `json:"process_instance"`
}
