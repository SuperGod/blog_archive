package model

import (
	"fmt"

	"github.com/GerardSoleCa/wordpress-hash-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
)

// Init init mysql connect
func Init(mysqlSource string) {
	var err error
	engine, err = xorm.NewEngine("mysql", mysqlSource)
	if err != nil {
		panic(err)
	}
}

// GetData return datas form db
func GetData(datas interface{}) (err error) {
	err = engine.Find(datas)
	return
}

func GetUserWithPwd(name, pwd string) (user User, err error) {
	user.UserLogin = name
	_, err = engine.Get(&user)
	if err != nil {
		return
	}
	if !wphash.CheckPassword(pwd, user.UserPass) {
		fmt.Println(user.UserPass, pwd)
		err = fmt.Errorf("password error")
		return
	}
	user.UserPass = ""
	return
}

func GetUserWithID(userID int64) (user User, err error) {
	_, err = engine.Id(userID).Get(&user)
	if err == nil {
		user.UserPass = ""
	}
	return
}

// GetUsers return all users
func GetUsers(u *User) (users []User, err error) {
	err = engine.Find(&users, u)
	return
}

// GetPosts find posts
func GetPosts(p *Post) (data []Post, err error) {
	err = engine.Find(&data, p)
	return
}

// GetPublishPosts get all publish posts
func GetPublishPosts() (data []Post, err error) {
	err = engine.Where("post_parent=? and post_type=? and post_status=?", 0, "post", "publish").Find(&data)
	return
}

// GetPublishPosts get all publish posts
func GetRecentPublishPosts(skip, limit int) (data []Post, err error) {
	err = engine.Where("post_parent=? and post_type=? and post_status=?", 0, "post", "publish").Desc("post_modified").Limit(limit, skip).Find(&data)
	return
}

// GetPublishPost get all publish posts
func GetPublishPost(name string) (data Post, err error) {
	_, err = engine.Where("post_parent=? and post_type=? and post_status=? and post_name=?", 0, "post", "publish", name).Get(&data)
	return
}
