package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var UserRepository = userRepository{}
var engine *xorm.Engine

// User is
type User struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"name" xorm:"'nickname'"`
}

// init ...
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "bookshelf:bookshelf@/bookshelf") //データベースのためのEngineを作る(DB, ユーザ名：パスワード/ DB名)
	if err != nil {
		panic(err)
	}
}

// NewUser ...
func NewUser(id int, username string) User {
	return User{
		ID:       id,
		Username: username,
	}
}

// UserRepository is
type userRepository struct {
}

// NewUserRepository ...
func NewUserRepository() userRepository {
	return userRepository{}
}

// GetByID ...
func (m userRepository) GetByID(id int) *User {
	var user = User{ID: id}
	has, err := engine.Get(&user)
	if err != nil {
		panic(err)
	}
	if has {
		return &user
	}
	return nil
}

// Create ...
func (m userRepository) Create(name string) {
	var user = User{Username: name}
	_, err := engine.Insert(&user)
	if err != nil {
		panic(err)
	}
}
