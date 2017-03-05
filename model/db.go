package model

import (
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
