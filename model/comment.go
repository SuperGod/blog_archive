package model

import "time"

type Comment struct {
	ID       int64     `xorm:"comment_ID pk"`
	PostID   int64     `xorm:"comment_post_ID"`
	Author   string    `xorm:"comment_author"`
	Email    string    `xorm:"comment_author_email"`
	URL      string    `xorm:"comment_author_url"`
	IP       string    `xorm:"comment_author_IP"`
	Date     time.Time `xorm:"comment_date"`
	DateGmt  time.Time `xorm:"comment_date_gmt"`
	Content  string    `xorm:"comment_content"`
	Karma    int       `xorm:"comment_karma"`
	Approved string    `xorm:"comment_approved"`
	Agent    string    `xorm:"comment_agent"`
	Type     string    `xorm:"comment_type"`
	Parent   int64     `xorm:"comment_parent"`
	UserID   int64     `xorm:"user_id"`
}

// TableName return the tablename
func (c *Comment) TableName() string {
	return "wp_comments"
}

type CommentMeta struct {
	ID        int64  `xorm:"meta_id pk"`
	CommentID int64  `xorm:"comment_id"`
	MetaKey   string `xorm:"meta_key"`
	MetaValue string `xorm:"meta_value"`
}

// TableName return the tablename
func (cm *CommentMeta) TableName() string {
	return "wp_commentmeta"
}
