package model

import (
	"testing"
)

func init() {
	Init("root:123456@tcp(172.17.0.1:3306)/blog_super?charset=utf8")
}

func TestInit(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Log("init test ok", err)
		}
		Init("root:123456@tcp(172.17.0.1:3306)/blog_super?charset=utf8")
	}()
	Init("root:123456@tcp(1:3306)/blog_super?charset=utf8")
}

func TestUser(t *testing.T) {
	users, err := GetUsers(&User{ID: 0})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("users:", users)
}

func TestUserByID(t *testing.T) {
	users, err := GetUserWithID(0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("users:", users)
}
func TestUserByPwd(t *testing.T) {
	user, err := GetUserWithID(0)
	if err != nil {
		t.Fatal(err)
	}
	_, err = GetUserWithPwd(user.UserLogin, user.UserPass)
	if err != nil {
		// t.Fatal(err)
		t.Log("login:", err)
	}

	_, err = GetUserWithPwd(user.UserLogin, "123")
	if err == nil {
		t.Fatal("password is error but return no error")
	}
	t.Log("users:", user)
}

func TestUserMeta(t *testing.T) {
	var users []UserMeta
	err := GetData(&users)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("users meta:", users)
}
func TestPosts(t *testing.T) {
	posts, err := GetPosts(&Post{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("posts:", len(posts))
}

func TestPostsMeta(t *testing.T) {
	var postMeta []PostMeta
	err := GetData(&postMeta)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("post meta:", len(postMeta))
}

func TestComment(t *testing.T) {
	var comments []Comment
	err := GetData(&comments)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("comments:", len(comments))
}

func TestCommentMeta(t *testing.T) {
	var commentMetas []CommentMeta
	err := GetData(&commentMetas)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("comment metas:", len(commentMetas))
}

func TestTerm(t *testing.T) {
	var terms []Term
	err := GetData(&terms)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("terms:", len(terms))
}

func TestTermMeta(t *testing.T) {
	var termMetas []TermMeta
	err := GetData(&termMetas)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("term metas:", len(termMetas))
}
func TestTermRelationship(t *testing.T) {
	var termRels []TermRelationship
	err := GetData(&termRels)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("term relationships:", len(termRels))
}

func TestTermTaxonomyelationship(t *testing.T) {
	var termTax []TermTaxonomy
	err := GetData(&termTax)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("term taxonomy:", len(termTax))
}

func TestPublishPosts(t *testing.T) {
	posts, err := GetPublishPosts()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("posts len:", len(posts))
	post, err := GetPublishPost(posts[0].Name)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("post find:", post.Name, post.Title)
}

func TestRecentPublishPosts(t *testing.T) {
	posts, err := GetRecentPublishPosts(0, 10)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("post find:", len(posts))
}
