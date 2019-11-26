package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Conn *gorm.DB

func init() {
	var err error
	Conn, err = gorm.Open("sqlite3", "dingtalk.db")
	if err != nil {
		panic(err)
	}
}
