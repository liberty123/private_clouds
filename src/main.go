package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"fmt"
)

func main() {
	r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
	})
	r.POST("/form_post", func(c *gin.Context) {
		id := c.Query("id")
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonumous")
		c.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
			"id": id,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		c.JSON(200, gin.H{
			"status": "posted",
			"ids": ids,
			"names": names,
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		dst := "./new_file"
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
    r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}