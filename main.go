package main


import (
	"blogger/dao/mysql"
	"blogger/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	err := db.Init("root:youknow.@tcp(127.0.0.1:3306)/blogger?parseTime=true")
	if err != nil {
		panic(err)
		return
	}

	router := gin.Default()
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandler)
	router.GET("/category", controller.CategoryHandler)

	router.Run(":5566")
}
