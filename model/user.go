package model

import (
	"crypto/sha256"
	"reflect"
	"regexp"
)

type User struct {
	Model
	UserName        string `form:"user_name" json:"user_name"`
	Password        string `form:"password" json:"password,omitempty"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm,omitempty"`
	PasswordHash    []byte `json:"-"`
}

const hashRepeatCount = 10

var (
	regstr = regexp.MustCompile("^[a-zA-Z0-9_]+$")
)

// パスワードのハッシュを生成します。
// ハッシュの生成はSHA256で10回繰り返します。
func makePasswordHash(password string) []byte {
	hash := []byte(password)
	for i := 0; i < hashRepeatCount; i++ {
		tmp := sha256.Sum256(hash)
		hash = tmp[:]
	}
	return hash
}

// パスワードが確認用パスワードと一致しているかどうかチェックします。
func (user User) CheckPassword() bool {
	return user.Password == user.PasswordConfirm
}

// パスワードからパスワードハッシュを生成してセットします。
func (user *User) SetPasswordHash() {
	user.PasswordHash = makePasswordHash(user.Password)
}

// 指定されたパスワードがユーザーのパスワードと一致するかパスワードハッシュの比較を行います。
func (user *User) Authorize(password string) bool {
	hash := makePasswordHash(password)
	return reflect.DeepEqual(hash, user.PasswordHash)
}

// ユーザ名に不正な文字列が入力されていないか検証を行います
func (user *User) IsValidUserName() bool {
	return regstr.MatchString(user.UserName)
}
