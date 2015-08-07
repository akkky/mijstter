package db

import (
	"mijstter/model"
	"reflect"
	"testing"
)

func TestUserReadWrite(t *testing.T) {
	user := &model.User{UserName: "test", Password: "abcdefghijklmn"}
	user.SetPasswordHash()

	err := db.WriteUser(user)
	if err != nil {
		t.Errorf("user can not insert.\n%v", err)
	}
	if user.Id == 0 {
		t.Errorf("user id is not set.")
	}

	user2, err := db.ReadUser(user.Id)
	if err != nil {
		t.Errorf("user can not select.\n%v", err)
	}
	if user2 == nil {
		t.Errorf("user is not returned.")
	}

	user.Password = ""
	user.PasswordConfirm = ""
	if !reflect.DeepEqual(user, user2) {
		t.Errorf("got %v\nwant %v", user2, user)
	}
}

func TestReadUsers(t *testing.T) {
	err := db.DeleteAllUsers()
	if err != nil {
		t.Errorf("users are not deleted.\n%v", err)
	}

	users := []model.User{
		model.User{UserName: "test1", Password: "abcdefg"},
		model.User{UserName: "test2", Password: "abcdefg"},
		model.User{UserName: "test3", Password: "abcdefg"},
	}

	for _, user := range users {
		user.SetPasswordHash()
		db.WriteUser(&user)
	}

	readUsers, err := db.ReadUsers(2)
	if err != nil {
		t.Errorf("users are not read.\n%v", err)
	}

	actual := len(readUsers)
	expected := 2
	if actual != expected {
		t.Errorf("got %v users\nwant %v users", actual, expected)
	}
}

func IsUserNameExist(t *testing.T) {
	err := db.DeleteAllUsers()
	if err != nil {
		t.Errorf("users are not deleted.\n%v", err)
	}

	users := []model.User{
		model.User{UserName: "test1", Password: "abcdefg"},
		model.User{UserName: "test2", Password: "abcdefg"},
		model.User{UserName: "test3", Password: "abcdefg"},
	}

	for _, user := range users {
		user.SetPasswordHash()
		db.WriteUser(&user)
	}

	expected := true
	actual, err := db.IsUserNameExist("test2")
	if err != nil {
		t.Errorf("user can not select.\n%v", err)
	}

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	expected = false
	actual, err = db.IsUserNameExist("test4")
	if err != nil {
		t.Errorf("user can not select.\n%v", err)
	}

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
