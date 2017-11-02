package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Article struct {
	gorm.Model
	Title string
	Body  string
	//Tags      []*Tag `gorm:"-"` // tags of post
}

type User struct {
	gorm.Model
	Username  string
	Password  string
	GithubUrl string
	IsAdmin   bool
	AvatarUrl string
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	//db, err := gorm.Open("mysql", "root:mysql@/wblog?charset=utf8&parseTime=True&loc=Asia/Shanghai")
	if err == nil {
		DB = db
		db.AutoMigrate(&Article{}, &User{})
		return db, err
	}
	return nil, err
}
