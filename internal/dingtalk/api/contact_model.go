package api

type UserId struct {
	ErrResponse
	Userid string `json:"userid"`
}
type UidByUnionid struct {
	ErrResponse
	ContactType int    `json:"contactType"`
	Userid      string `json:"userid"`
}

type UserList struct {
	Userid    string `json:"userid" gorm:"column:userid"`
	Name      string `json:"name" gorm:"column:name"`
	CreatedAt string `gorm:"column:createdat"`
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
