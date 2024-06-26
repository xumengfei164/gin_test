package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("in my branch")
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.PUT("/tt", func(context *gin.Context) {
		context.JSON(209, gin.H{"msg": "success."})
	})
	server.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{
			"info": gin.H{
				"name": "徐梦飞",
				"sex":  true,
				"like": []string{"吃饭", "觅爱", "学习", "健身", "冥想"},
				"age":  26.3,
			},
		})
	})
	server.Run(":8082")
}
