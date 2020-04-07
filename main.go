package main

import (
	"./functions/sql"
	"./route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//echo框架初始化
	e := echo.New()
	//注册路由
	route.Route(e)
	//数据库连接池初始化
	sql.Init_Db()
	//应用停止时关闭数据库
	defer sql.DB_close()
	//允许跨域
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))
	//启动echo
	e.Logger.Fatal(e.Start(":2333"))

}
