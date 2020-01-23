package api

type AttenGroup struct {
	Result struct {
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Wifis struct {
			String []string `json:"string"`
		} `json:"wifis"`
		AddressList struct {
			String []string `json:"string"`
		} `json:"address_list"`
		WorkDayList struct {
			Number []interface{} `json:"number"`
		} `json:"work_day_list"`
		MemberCount int    `json:"member_count"`
		Type        string `json:"type"`
		URL         string `json:"url"`
		ManagerList string `json:"manager_list"`
	} `json:"result"`
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// 使用gorm标签，如有特殊用途，自行修改tag
type Schedule struct {
	PlanID         int    `json:"plan_id" gorm:"column:planid"`
	CheckType      string `json:"check_type" gorm:"column:checktype"`
	ApproveID      int    `json:"approve_id" gorm:"column:approveid"`
	Userid         string `json:"userid" gorm:"column:userid"`
	ClassID        int    `json:"class_id" gorm:"column:classid"`
	ClassSettingID int    `json:"class_setting_id" gorm:"column:classsettingid"`
	PlanCheckTime  string `json:"plan_check_time" gorm:"column:planchecktime"`
	GroupID        int    `json:"group_id" gorm:"column:groupid"`
	CreatedAt      string `gorm:"column:createdat"`
	UserName       string `gorm:"-"`
}

type ListSchedule struct {
	ErrResponse
	Result struct {
		HasMore bool `json:"has_more"`
		// Schedules []misc.Data
		Schedules []Schedule `json:"schedules"`
	} `json:"result"`
}

// 获取考勤的班次摘要信息
type GroupMinimalismList struct {
	CreatedAt string `gorm:"column:createdat"`
	ErrResponse
	Result struct {
		HasMore bool `json:"has_more" gorm:"-"`
		Cursor  int  `json:"cursor" gorm:"-"`
		Result  []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"result"`
	} `json:"result"`
}

// ListRecord: 打卡记录
type ListRecord struct {
	ErrResponse
	RecordResult []struct {
		BaseCheckTime  int64   `json:"baseCheckTime"`
		Id             int64   `json:"id"`
		WorkDate       int64   `json:"workDate"`
		PlanCheckTime  int64   `json:"planCheckTime"`
		PlanId         int64   `json:"planId"`
		GroupId        int64   `json:"groupId"`
		UserCheckTime  int64   `json:"userCheckTime"`
		UserLongitude  float64 `json:"userLongitude"`
		UserAccuracy   float64 `json:"userAccuracy"`
		UserLatitude   float64 `json:"userLatitude"`
		IsLegal        string  `json:"isLegal"`
		UserAddress    string  `json:"userAddress"`
		UserId         string  `json:"userId"`
		CheckType      string  `json:"checkType"`
		TimeResult     string  `json:"timeResult"`
		DeviceId       string  `json:"deviceId"`
		CorpId         string  `json:"corpId"`
		SourceType     string  `json:"sourceType"`
		LocationMethod string  `json:"locationMethod"`
		LocationResult string  `json:"locationResult"`
		ProcInstId     string  `json:"procInstId"`
	}
}

// 考勤组摘要结构
type AttdGroup struct {
	ErrResponse
	Result []struct {
		Name string `json:"Name"`
		Id   int    `json:"id"`
	} `json:"result"`
}
