package controller

import (
	"strconv"

	"github.com/SuperGod/blog/model"
	"github.com/gin-gonic/gin"
)

type BlogController struct {
}

func (bc *BlogController) GetBlog(c *gin.Context) {
	blogName := c.Param("blogName")
	if blogName == "" {
		c.String(404, "no such blog")
		return
	}
	content, err := model.GetPublishPost(blogName)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(200, content)
	return
}

func (bc *BlogController) GetBlogList(c *gin.Context) {
	nStart := DefaultQueryInt(c, "start", 0)
	nLimit := DefaultQueryInt(c, "limit", 10)
	contentList, err := model.GetRecentPublishPosts(nStart, nLimit)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(200, contentList)
	return
}

// DefaultQueryInt query with default int value
func DefaultQueryInt(c *gin.Context, name string, nDefault int) int {
	value := c.DefaultQuery(name, "")
	if value == "" {
		return nDefault
	}
	nValue, err := strconv.Atoi(value)
	if err != nil {
		return nDefault
	}
	return nValue
}
