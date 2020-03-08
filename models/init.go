/*
Author: Landers
Github: Landers1037
Date: 2020-02
Name: cloudp
*/
package models
//数据库的初始化
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"cloudp/models/sql"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	//CreatedOn int `json:"created_on"`
	//ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err error
	)

	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	//	return tablePrefix + defaultTableName
	//}
	//使用sqlite
	db,err = sql.Sqlite()
	if err != nil {
		log.Println(err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
