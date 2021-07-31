package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context){
	name := c.DefaultQuery("name","jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func main() {
	//fmt.Println("hello")
	r := gin.Default()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20

	// 1. 上传单个文件
	/*r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})

		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})*/

	// 2. 上传多个文件
	/*r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})*/

	// 3. 路由分组
	// 路由组1 处理get请求
	v1 := r.Group("/v1")
	{
		v1.GET("/login",login)
		v1.GET("/submit",submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login",login)
		v2.POST("/submit",submit)
	}
	r.Run(":10000")
}
