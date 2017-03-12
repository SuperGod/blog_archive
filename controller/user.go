package controller

import (
	"fmt"
	"strconv"

	"github.com/SuperGod/blog/model"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (uc *UserController) GetLogin(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("UserID")
	if v == nil {
		c.String(200, "")
		return
	}
	c.Redirect(301, fmt.Sprintf("/user/info/%d", v.(int64)))
}

func (uc *UserController) PostLogin(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("UserID")
	var user model.User
	var err error
	fmt.Println("userID:", v)
	if v != nil {
		user, err = model.GetUserWithID(v.(int64))
		if err != nil {
			c.String(500, err.Error())
			return
		}
		c.JSON(200, gin.H{"status": "ok", "user": user})
		return
	}
	var userInfo struct {
		Name string
		Pwd  string
	}
	err = c.Bind(&userInfo)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	// md5Buf := md5.Sum([]byte(userInfo.Pwd))
	// user, err := model.GetUserWithPwd(userInfo.Name, hex.EncodeToString(md5Buf[0:]))
	user, err = model.GetUserWithPwd(userInfo.Name, userInfo.Pwd)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	session.Set("UserID", user.ID)
	session.Save()
	c.JSON(200, gin.H{"status": "ok", "user": user})
}

func (uc *UserController) Info(c *gin.Context) {
	userID, ok := c.Params.Get("id")
	if !ok {
		c.Redirect(200, "/user/login")
		return
	}
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	user, err := model.GetUserWithID(id)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "ok", "user": user})
}
