package main

import (
	"blogger/controller"
	"blogger/dao/db"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("blogger")

func main() {
	r := gin.Default()
	driverSql := "root:rootroot@tcp://(127.0.0.1:3306)/blogger?parseTime=true"
	db.InitDB(driverSql)

	//加载静态文件
	r.Static("/static/", "./static")
	//加载模板
	r.LoadHTMLGlob("views/*")

	r.GET("/", controller.IndexHandler)
	r.GET("/category/", controller.CategoryList)

	r.Run("0.0.0.0:8000")
}
