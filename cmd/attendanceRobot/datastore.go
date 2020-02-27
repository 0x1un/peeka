package main

import (
	"boxes/internal/dingtalk/api"

	"github.com/jinzhu/gorm"
)

func InsertRecord(conn *gorm.DB, data interface{}) error {
	var err error
	handler := conn.Begin()
	if value, ok := data.(api.Schedule); ok {
		handler = handler.Table("atten_list")
		value.CreatedAt = GET_TIME.Format(DATE_FORMAT)
		err = handler.Create(value).Error
		if err != nil {
			handler.Rollback()
			return err
		}
	}
	if value, ok := data.(api.UserList); ok {
		handler = handler.Table("dep_users")
		value.CreatedAt = GET_TIME.Format(DATE_FORMAT)
		err = handler.Create(value).Error
		if err != nil {
			handler.Rollback()
			return err
		}
	}
	if value, ok := data.(api.GroupMinimalismList); ok {
		handler = handler.Table("atten_class")
		value.CreatedAt = GET_TIME.Format(DATE_FORMAT)
		err = handler.Create(value).Error
		if err != nil {
			handler.Rollback()
			return err
		}
	}
	handler.Commit()
	return nil
}

func UpdateRecord(conn *gorm.DB, data interface{}) error {
	var err error
	handler := conn.Begin()
	if value, ok := data.(api.UserList); ok {
		handler = handler.Table("dep_users")
		value.CreatedAt = GET_TIME.Format(DATE_FORMAT)
		err = handler.Save(value).Error
		if err != nil {
			handler.Rollback()
			return err
		}
	}
	handler.Commit()
	return nil
}

func QueryRecord(conn *gorm.DB, conditions []string, records *[]api.Schedule) error {
	err := conn.
		Table("atten_list").
		Where("checktype = ? AND createdat = ? AND userid in (?)",
			"OnDuty",
			GET_TIME.Format(DATE_FORMAT),
			conditions).
		Find(records).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDepUsers(conn *gorm.DB) *[]api.UserList {
	var err error
	users := new([]api.UserList)
	names := []string{
		"张军", "邹一",
		"唐顺", "唐建",
		"王彪", "李耀",
		"高远", "刘环",
		"陈浩", "尹升俊", "赵鹏辉",
	}
	err = conn.
		Table("dep_users").
		Select("name,userid").
		Where(
			"name in (?)",
			// GET_TIME.Format(DATE_FORMAT),
			names,
		).Scan(users).Error
	if err != nil {
		panic(err)
	}
	if len(*users) != len(names) {
		// 部门id, 调api获取吧
		err = GetAllUserInDepartment("105372678", "0", "100", "")
		if err != nil {
			panic(err)
		}
		return GetDepUsers(conn)
	}
	return users
}
