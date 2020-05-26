package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 路由分组
	v1 := r.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			name := c.PostForm("name")
			c.JSON(200, gin.H{
				"name": name,
			})
		})
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			name := c.PostForm("name")
			c.JSON(200, gin.H{
				"name": name,
			})
		})
	}
	r.Run(":8080")
}
