package models

import (
	"database/sql"
	//"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func init() {
	log.SetFlags(log.Lshortfile)

	var err error
	DB, err = sql.Open("mysql", "root:@(localhost:3306)/htz_search?parseTime=true")
	if nil != err {
		log.Fatal(err)
	}
}
