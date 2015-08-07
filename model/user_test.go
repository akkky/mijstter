package model

import (
	"reflect"
	"testing"
)

func TestMakePasswordHash(t *testing.T) {
	actual := makePasswordHash("abcdefghijklmn")
	expected := []byte{15, 146, 20, 247, 72, 49, 75, 250, 209, 76, 106, 93, 83, 94, 67, 202, 115, 104, 137, 57, 10, 0, 78, 120, 163, 221, 209, 253, 121, 57, 194, 46}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestUserCheckPassword(t *testing.T) {
	user := &User{
		UserName:        "Test",
		Password:        "abcdef",
		PasswordConfirm: "abcdef",
	}

	var actual, expected bool

	actual = user.CheckPassword()
	expected = true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	user.PasswordConfirm = "abcdeF"
	actual = user.CheckPassword()
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestSetPasswordHash(t *testing.T) {
	user := &User{Password: "abcdefghijklmn"}

	user.SetPasswordHash()
	actual := user.PasswordHash
	expected := makePasswordHash(user.Password)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestAuthorize(t *testing.T) {
	user := &User{Password: "abcdefghijklmn"}
	user.SetPasswordHash()

	actual := user.Authorize("abcdefghijklmn")
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	actual = user.Authorize("abcdefghijklmN")
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestIsValidUserName(t *testing.T) {
	user := &User{UserName: "0123456789abcdefgA_"}

	actual := user.IsValidUserName()
	expected := true
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	user.UserName = "*!#$%&"

	actual = user.IsValidUserName()
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}

	user.UserName = ""

	actual = user.IsValidUserName()
	expected = false
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
