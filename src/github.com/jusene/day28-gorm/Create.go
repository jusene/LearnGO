package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Dsn struct {
	Host string
	User string
	Pass string
	Port int
	DB   string
}

func main() {
	var dsn = &Dsn{
		Host: "127.0.0.1",
		User: "root",
		Pass: "123456",
		Port: 3306,
		DB:   "user",
	}

	var ds = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dsn.User, dsn.Pass,
		dsn.Host, dsn.Port, dsn.DB)

	sqlDB, _ := sql.Open("mysql", ds)
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	type User struct {
		gorm.Model
		Name     string
		Age      int
		Birthday time.Time
	}

	gormDB.AutoMigrate(&User{})

	user := User{
		Name:     "Jusene",
		Age:      33,
		Birthday: time.Now(),
	}

	result := gormDB.Create(&user)

	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)

	var u User
	gormDB.Debug().First(&u, "id = ?", 1)
	fmt.Println(u.ID)

}
