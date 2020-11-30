package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Model struct {
	ID       uint `gorm:"primaryKey"`
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

type Model2 struct {
	ID uint
	gorm.Model
}

type User1 struct {
	Name  string `gorm:"<-:create"`          // allow read and create
	Name1 string `gorm:"<-:update"`          // allow read and update
	Name2 string `gorm:"<-"`                 // allow read and write (create and update)
	Name3 string `gorm:"<-:false"`           // allow read, disable write permission
	Name4 string `gorm:"->"`                 // readonly (disable write permission unless it configured )
	Name5 string `gorm:"->;<-:create"`       // allow read and create
	Name6 string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
	Name7 string `gorm:"-"`                  // ignore this field when write and read
}

type User2 struct {
	CreatedAt time.Time // Set to current time if it is zero on creating
	UpdatedAt int       // Set to current unix seconds on updaing or if it is zero on creating
	Updated   int64     `gorm:"autoUpdateTime:nano"`  // Use unix nano seconds as updating time
	Updated1  int64     `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
	Created   int64     `gorm:"autoCreateTime"`       // Use unix seconds as creating time
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      uint
	Author  Author `gorm:"embedded;embeddedPrefix:author_"`
	Upvotes int32
}

func main() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&Model{})
	//db.AutoMigrate(&Model2{})
	//db.AutoMigrate(&User2{})
	db.AutoMigrate(&Blog{})
	//db.Create(&User1{Name: "jusene"})
}
