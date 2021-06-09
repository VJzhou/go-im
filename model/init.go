package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

func InitDatabase () *xorm.Engine{
	db, err := xorm.NewEngine("mysql", "")
	if err != nil {
		log.Fatal(err.Error())
	}
	db.ShowSQL(true)
	db.SetMaxOpenConns(1)
	err = db.Sync2(new(User))
	return db
}