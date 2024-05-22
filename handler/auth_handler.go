package handler

import (
	"gin_social/dto"
	"gin_social/errorhandler"
	"gin_social/helpers"
	"gin_social/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct{
	service services.AuthService
}

func NewAuthHandler(s services.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandleError(c, &errorhandler.ErrrorNotFound{
			Message: err.Error(),
		})
		return
	}

	if err := h.service.Register(&register); err != nil{
		errorhandler.HandleError(c, err)

		return
	}

	res := helpers.Response(&dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message: "Registered Successfully. Please Login",
	})

	c.JSON(http.StatusCreated, res)
}