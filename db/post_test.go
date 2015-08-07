package db

import (
	"mijstter/model"
	"reflect"
	"testing"
)

func TestPostReadWrite(t *testing.T) {
	post := &model.Post{UserId: 1, UserName: "test", Message: "This is a test.", Url: "http://example.com/"}

	err := db.WritePost(post)
	if err != nil {
		t.Errorf("post can not insert.\n%v", err)
	}
	if post.Id == 0 {
		t.Errorf("post id is not set.")
	}

	post2, err := db.ReadPost(post.Id)
	if err != nil {
		t.Errorf("post can not select.\n%v", err)
	}
	if post2 == nil {
		t.Errorf("post is not returned.")
	}

	if !reflect.DeepEqual(post, post2) {
		t.Errorf("got %v\nwant %v", post2, post)
	}
}

func TestReadPosts(t *testing.T) {
	err := db.DeleteAllPosts()
	if err != nil {
		t.Errorf("posts are not deleted.\n%v", err)
	}

	posts := []model.Post{
		model.Post{UserId: 1, UserName: "test1", Message: "This is a test.", Url: "http://example.com/"},
		model.Post{UserId: 2, UserName: "test2", Message: "This is a test.", Url: "http://example.com/"},
		model.Post{UserId: 3, UserName: "test3", Message: "This is a test.", Url: "http://example.com/"},
	}

	for _, post := range posts {
		db.WritePost(&post)
	}

	readPosts, err := db.ReadPosts(2)
	if err != nil {
		t.Errorf("posts are not read.\n%v", err)
	}

	actual := len(readPosts)
	expected := 2
	if actual != expected {
		t.Errorf("got %v posts\nwant %v posts", actual, expected)
	}
}
