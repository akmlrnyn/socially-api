package helpers

import (
	"gin_social/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FilterParams(c *gin.Context) *dto.FilterParam {
page := c.DefaultQuery("page", "1")
limit := c.DefaultQuery("limit", "5")
search := c.Query("search")

pageNumber, _ := strconv.Atoi(page)
limitNumber, _ := strconv.Atoi(limit)
offset := (pageNumber - 1) * limitNumber

return &dto.FilterParam{
	Page: pageNumber,
	Limit: limitNumber,
	Offset: offset,
	Search: search,
} 

}
