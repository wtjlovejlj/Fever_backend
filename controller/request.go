package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func getCurrentUserName(c *gin.Context) (userName string, err error) {
	_userName, ok := c.Get(ContextUserNameKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userName, ok = _userName.(string)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func getPageInfo(c *gin.Context) (int, int) {
	//获取分页参数
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int
		size int
		err  error
	)

	page, err = strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}
	return page, size
}
