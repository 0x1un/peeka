package db

import (
	"time"

	"github.com/allegro/bigcache"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "accesstoken.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
}

func DBConnect() *gorm.DB {
	return db
}

func NewCache(min int) *bigcache.BigCache {
	if bc, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Duration(min) * time.Minute)); err != nil {
		return nil
	} else {
		return bc
	}
}
