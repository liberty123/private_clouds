package main

import (
	"github.com/gin-gonic/gin"
	"disposeImage"
	"net/http"
	"log"
	"fmt"
)

func main() {
	r := gin.Default()
	a,b := disposeImage.GetConfig()
	fmt.Println(a)
	fmt.Println(b)
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "OK",
        })
	})

	r.POST("/form_post", func(c *gin.Context) {
		id := c.Query("id")
		message := c.PostForm("message")
		//设定默认参数
		nick := c.DefaultPostForm("nick", "anonumous")
		c.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
			"id": id,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		// 获取数组类型的参数
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		c.JSON(200, gin.H{
			"status": "posted",
			"ids": ids,
			"names": names,
		})
	})

	//上传单个文件 
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		dst := "./new_file"
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	//上传文件列表
	r.POST("/upload_list", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		dst := "./"
		for _, file := range files {
			log.Println(file.Filename)
			c.SaveUploadedFile(file, dst+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files upload!", len(files)))
	})

    // r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}