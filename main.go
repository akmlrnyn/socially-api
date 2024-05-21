package main

import (
	"fmt"
	"gin_social/configs"

	"github.com/gin-gonic/gin"
)

func main() {

	configs.LoadConfig()
	configs.LoadDB()


	r := gin.Default()
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(":%v", configs.ENV.PORT)) // listen and serve on 0.0.0.0:8080
}