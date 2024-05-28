package main

import (
	"fmt"
	"gin_social/configs"
	"gin_social/routes"

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

	routes.AuthRoutes(api)
	routes.PostRouter(api)
	
	r.Run(fmt.Sprintf(":%v", configs.ENV.PORT)) // listen and serve on 0.0.0.0:8080
}