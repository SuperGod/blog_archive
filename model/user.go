package model

import "time"

// User table: wp_users
type User struct {
	ID                int64     `xorm:"ID pk"`
	UserLogin         string    `xorm:"user_login"`
	UserPass          string    `xorm:"user_pass"`
	UserNiceName      string    `xorm:"user_nicename"`
	UserEmail         string    `xorm:"user_email"`
	UserURL           string    `xorm:"user_url"`
	UserRegistered    time.Time `xorm:"created"`
	UserActivationKey string    `xorm:"user_activation_key"`
	UserStatus        int       `xorm:"user_status"`
	DisplayName       string    `xorm:"display_name"`
	Spam              int
	Deleted           int
}

// TableName return the tablename
func (u *User) TableName() string {
	return "wp_users"
}

type UserMeta struct {
	ID        int64  `xorm:"umeta_id pk"`
	UserID    int64  `xorm:"user_id"`
	MetaKey   string `xorm:"meta_key"`
	MetaValue string `xorm:"meta_value"`
}

// TableName return the tablename
func (u *UserMeta) TableName() string {
	return "wp_usermeta"
}
