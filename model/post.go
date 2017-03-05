package model

import "time"

type Post struct {
	ID              int64     `xorm:"ID pk"`
	Author          int64     `xorm:"post_author"`
	PostDate        time.Time `xorm:"post_date"`
	PostDateGmt     time.Time `xorm:"post_date_gmt"`
	Name            string    `xorm:"post_name"`
	Title           string    `xorm:"post_title"`
	Content         string    `xorm:"post_content"`
	Excerpt         string    `xorm:"post_excerpt"`
	PostStatus      string    `xorm:"post_status"`
	PostType        string    `xorm:"post_type"`
	PostMimeType    string    `xorm:"post_mime_type"`
	CommentStatus   string    `xorm:"comment_status"`
	CommentCount    int       `xorm:"comment_count"`
	PingStatus      string    `xorm:"ping_status"`
	PostPassword    string
	ToPing          string
	Pinged          string
	Modified        time.Time `xorm:"post_modified"`
	ModifiedGmt     time.Time `xorm:"post_modified_gmt"`
	ContentFiltered string    `xorm:"post_content_filtered"`
	Parent          int64     `xorm:"post_parent"`
	GUID            string    `xorm:"guid"`
	MenuOrder       int       `xorm:"menu_order"`
}

// TableName return the tablename
func (u *Post) TableName() string {
	return "wp_posts"
}

type PostMeta struct {
	ID        int64  `xorm:"meta_id pk"`
	PostID    int64  `xorm:"post_id"`
	MetaKey   string `xorm:"meta_key"`
	MetaValue string `xorm:"meta_value"`
}

// TableName return the tablename
func (u *PostMeta) TableName() string {
	return "wp_postmeta"
}
