package models

import (
	"blogapi/pkg/setting"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Model 数据模型基础字段
type Model struct {
	Id         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// init 初始化数据库连接
func init() {
	var (
		err         error
		dbType      string
		dbName      string
		user        string
		password    string
		host        string
		tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("加载区失败 err: %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true) // TODO 这是啥?
	db.LogMode(true)

	pool := db.DB()
	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(100)
}

// CloseDb 关闭数据库连接
func CloseDb() {
	defer db.Close()
}
