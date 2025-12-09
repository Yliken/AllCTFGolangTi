package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"flag": "http://flag{youseccess}.jpg/",
		})
	})
	engine.Run(":80")
}
