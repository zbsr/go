package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
)

var DB *xorm.Engine

func main()  {
	fmt.Println("test...")

	// db
	DB, err := xorm.NewEngine("mysql", fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Asia%%2FShanghai`,
		"test",
		"tzb984334645",
		"127.0.0.1",
		"3304",
		"test",
	))
	if err != nil {
		panic(err)
	}
	DB.SetLogLevel(core.LOG_DEBUG)

}