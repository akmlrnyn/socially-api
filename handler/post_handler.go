package handler

import (
	"fmt"
	"gin_social/dto"
	"gin_social/errorhandler"
	"gin_social/helpers"
	"gin_social/services"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct{
	services services.PostService
}

func NewPostHandler(service services.PostService) *postHandler{
	return &postHandler{
		services: service,
	}
}

func (h *postHandler) Create(c *gin.Context){
	var post *dto.PostRequest

	if err := c.ShouldBind(&post); err != nil{
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	//check if the user post image or not
	if post.Picture != nil{
		//make the directory for stored picture
		_, err := os.Stat("/public/picture")
		if err !=  nil && os.IsNotExist(err) {
			if err := os.MkdirAll("public/picture", 0755); err != nil{
			errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
			return
		}} 

		//rename the file thats stored
		ext := filepath.Ext(post.Picture.Filename)
		fileName := uuid.New().String() + ext

		//save image to directory
		dst := filepath.Join("public/picture", filepath.Base(fileName))
		c.SaveUploadedFile(post.Picture, dst)

		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", c.Request.Host, fileName)
	}

	userID, _ := c.Get("userID")
	post.UserId = userID.(int)

	if err := h.services.Create(post); err != nil{
		errorhandler.HandleError(c, err)
	}

	res := helpers.Response(&dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message: "Successfully upload your post",
	})

	c.JSON(http.StatusCreated, res)
}

func (h *postHandler) Get(c *gin.Context){
	filter := helpers.FilterParams(c)

	posts, paginate, err := h.services.FindAll(filter)

	if err != nil{
		errorhandler.HandleError(c, err)
	}

	res := &dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message: "Tweet Lists",
		Paginate: paginate,
		Data: posts,
	}

	c.JSON(http.StatusOK, res)
}