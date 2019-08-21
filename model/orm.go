package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zhengjianwen/hr_service/conf"
	"github.com/zhengjianwen/utils/log"
)

var Orm *xorm.Engine

func InitSQL() (err error) {
	switch conf.DbEngine {
	case "sqlite3":
		Orm, err = ConnSqlite()
		if err != nil || Orm == nil {
			log.Fatalf("[sqlite3] %v",err)
			return err
		}

	}

	initTables()
	return nil
}

func ConnSqlite() (*xorm.Engine,error) {
	db, err := xorm.NewEngine("sqlite3", "./db/hr_jbtm.db")
	if err != nil {
		log.Fatalf("[ConnSqlite] connect sqlite3 => %v\n", err)
		return nil,fmt.Errorf("连接失败")
	}
	if err = db.Ping(); err != nil {
		return nil,fmt.Errorf("链接数据库失败")
	}
	return db,err
}


func initTables()  {
	Orm.Sync(Service{})
	err := Orm.Sync(ServiceGroup{})
	if err != nil{
		log.Fatalf("[initTables] ")
	}
}
