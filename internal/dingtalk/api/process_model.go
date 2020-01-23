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

// Response Data
