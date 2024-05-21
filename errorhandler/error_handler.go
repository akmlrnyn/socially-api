package errorhandler

import (
	"gin_social/dto"
	"gin_social/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type){
	case *ErrrorNotFound:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *Unauthorized:
		statusCode = http.StatusUnauthorized
	}

	response := helpers.Response(&dto.ResponseParam{
		StatusCode: statusCode,
		Message: err.Error(),
	})

	c.JSON(statusCode, response)
}