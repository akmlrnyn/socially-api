package routes

import (
	"gin_social/configs"
	"gin_social/handler"
	"gin_social/repository"
	"gin_social/services"

	"github.com/gin-gonic/gin"
)


func AuthRoutes(api *gin.RouterGroup) {
	AuthRepo := repository.NewAuthRepository(configs.DB)
	AuthService := services.NewAuthService(AuthRepo)
	AuthHandler := handler.NewAuthHandler(AuthService)


	api.POST("/register", AuthHandler.Register)
	api.POST("/login", AuthHandler.Login)
}