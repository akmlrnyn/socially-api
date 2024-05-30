package routes

import (
	"gin_social/configs"
	"gin_social/handler"
	"gin_social/middleware"
	"gin_social/repository"
	"gin_social/services"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	postRepo := repository.NewPostRepository(configs.DB)
	postService := services.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/tweets")
	r.Use(middleware.JWTMiddleware())

	r.GET("/", postHandler.Get)
	r.GET("/:id/detail", postHandler.Detail)
	r.POST("/create", postHandler.Create)
}
