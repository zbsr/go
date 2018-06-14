package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"net/http"
)

var DB *xorm.Engine

func init() {
	godotenv.Load()
}

func main()  {
	fmt.Println("test...")

	// db
	var err error
	DB, err = xorm.NewEngine("mysql", fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Asia%%2FShanghai`,
		GetEnv("DB_USERNAME"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_HOST"),
		GetEnv("DB_PORT"),
		GetEnv("DB_DATABASE"),
	))
	if err != nil {
		fmt.Println("wrong")
		panic(err)
	}

	DB.ShowSQL()
	DB.ShowExecTime()
	DB.SetLogLevel(core.LOG_DEBUG)
	DB.IsTableExist("test")

	user := make([]UserModel, 0)
	err = DB.Find(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8000")

}