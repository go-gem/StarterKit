package models

import (
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"github.com/go-gem/StarterKit/src/advanced/common/utils"
	"time"
)

func init() {
	gob.Register(User{})
}

type User struct {
	ID           int       `json:"id" gorm:"primary_key;column:id"`
	Email        string    `json:"email" gorm:"column:email"`
	Username     string    `json:"username" gorm:"column:username"`
	PasswordHash string    `json:"-" gorm:"column:password_hash"`
	PasswordSalt string    `json:"-" gorm:"column:password_salt"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (u *User) TableName() string {
	return "user"
}

func (u User) JoinAt() string {
	return u.CreatedAt.Format("2006-01-02 03:04")
}

func (u *User) GeneratePassword(password string) {
	u.PasswordSalt = utils.RandomString(8)
	h := md5.New()
	h.Write([]byte(password + u.PasswordSalt))
	u.PasswordHash = hex.EncodeToString(h.Sum(nil))
}

func (u User) ValidatePassword(password string) bool {
	h := md5.New()
	h.Write([]byte(password + u.PasswordSalt))
	return u.PasswordHash == hex.EncodeToString(h.Sum(nil))
}
