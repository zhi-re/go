package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// r := gin.New() // 无中间件启动
	r := gin.Default() // 默认启动方式，包含 Logger、Recovery 中间件

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// r.POST("/somePost")
	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	// http://localhost:8080/user/cqeeee
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		log.Println("get请求入参：" + name)
		c.JSON(200, gin.H{
			"name": name,
		})
	})

	// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.JSON(200, gin.H{
			"name":   name,
			"action": action,
		})
	})

	// 获取get请求入参
	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		//c.JSON(200, gin.H{
		//	"name":   name,
		//	"age": age,
		//})
		c.String(http.StatusOK, "Hello %s %s", name, age)
	})

	// 获取post请求表单提交
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// 给表单限制上传大小 (默认 32 MiB)
	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 上传文件到指定的路径
		result := c.SaveUploadedFile(file, "D:/go/"+file.Filename)
		log.Println(result)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	r.POST("/upload2", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["uploads"]
		for _, file := range files {
			log.Println(file.Filename)
			// 上传文件到指定的路径
			c.SaveUploadedFile(file, "D:/go/"+file.Filename)
		}

		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 创建记录日志的文件
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	//
	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//
	//	// 你的自定义格式
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	// }))
	// r.Use(gin.Recovery())

	r.Run(":8080")
}
