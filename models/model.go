package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"PushServer/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID         int64 `gorm:"primary_key" json:"id"`
	CreatedOn  int64 `json:"created_on"`
	ModifiedOn int64 `json:"modified_on"`
}

func InitSetUp() {

	var (
		dbType      = setting.DbConfig.Type
		dbName      = setting.DbConfig.Name
		user        = setting.DbConfig.User
		password    = setting.DbConfig.PassWord
		host        = setting.DbConfig.Host
		tablePrefix = setting.DbConfig.TablePrefix
	)

	var err error
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println("db connect error :", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.LogMode(setting.DbConfig.Debug)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDb() {
	db.Close()
}
