package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	// 使用LoadHTMLGlob() 或者 LoadHTMLFiles()
	// router.LoadHTMLGlob("src/main/resources/templates/index.tmpl")
	router.LoadHTMLGlob("src/main/resources/templates/**/**")
	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "嗯~ o(*￣▽￣*)o",
		})
	})
	// 在不同目录中使用具有相同名称的模板
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	//发布HTTP重定向很容易，支持内部和外部链接
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://zhi.re/")
	})

	// Gin路由重定向，使用如下的HandleContext
	router.GET("/test1", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})

	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
	router.Run(":8080")

}
