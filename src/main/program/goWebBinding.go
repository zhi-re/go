package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Login struct {
	UserName string `form:"userName" json:"userName" xml:"userName"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"-"`
}

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func main() {

	r := gin.Default()
	r.POST("/login", func(c *gin.Context) {
		var json Login
		err := c.ShouldBindJSON(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(json)
		if json.UserName != "cq" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "账号或密码错误!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "登录成功！"})
	})
	// 返回json数据
	r.POST("/getJSON", func(c *gin.Context) {
		// 为啥必须大写才可以？？？
		var message struct {
			Name string ` json:"userName"`
			Pass string ` json:"password"`
			Age  int
		}
		message.Name = "cq"
		message.Pass = "123"
		log.Println(message)
		c.JSON(http.StatusOK, message)
	})

	// JSONP
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// 访问 http://localhost:8080/JSONP?callback=call
		// 将会输出:   call({foo:"bar"})
		c.JSONP(http.StatusOK, data)
	})

	// AsciiJSON
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO",
			"tag":  "<br>",
		}

		// 将输出: {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// JSON会将特殊的HTML字符替换为对应的unicode字符
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// PureJSON 原样输出html
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	// 绑定Uri  localhost:8088/cq/987fbc97-4bed-5078-9f07-9141ba07c9f3
	//r.GET("/:username/:password", func(c *gin.Context) {
	//	var person Person
	//	if err := c.ShouldBindUri(&person); err != nil {
	//		c.JSON(400, gin.H{"msg": err})
	//		return
	//	}
	//	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	//})
	// 绑定Header
	r.GET("/", func(c *gin.Context) {
		h := testHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(200, err)
		}

		fmt.Printf("%#v\n", h)
		c.JSON(200, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	})
	// 访问静态资源
	// http://localhost:8080/resources/index.html
	// r.Static("/resources", "./resources")
	r.StaticFS("/resources", http.Dir("D:\\gowork\\gogogo\\src\\main\\resources"))
	// r.StaticFS("/more_static", http.Dir("my_file_system"))
	// r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 返回第三方获取的数据
	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://xile.xyz/logo.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	r.Run(":8080")

}
