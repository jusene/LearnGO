package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func main() {
	/*
		dsn := "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
		db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		fmt.Println(db.Name())
	*/

	/*
		db, _ = gorm.Open(mysql.New(mysql.Config{
			DSN: "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8&parseTime=True&loc=Local", // data source name
			DefaultStringSize: 256, // default size for string fields
			DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		}), &gorm.Config{})
		fmt.Println(db.Name())
	*/

	// 已经存在的数据库连接
	sqlDB, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local")
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	// 连接池
	// 设置最大空闲连接时间
	sqlDB.SetConnMaxIdleTime(10)
	// 在通过MySQL服务器，OS或其他中间件关闭连接之前，需要确保驱动程序安全地关闭连接。由于某些中间件会在5分钟之前关闭空闲连接，因此我们建议超时时间应少于5分钟。此设置也有助于负载平衡和更改系统变量。
	sqlDB.SetConnMaxLifetime(4 * time.Minute)
	// 强烈建议您限制应用程序使用的连接数。没有建议的限制数，因为它取决于应用程序和MySQL服务器。
	sqlDB.SetMaxOpenConns(1000)
	// 设置最大空闲连接
	sqlDB.SetMaxIdleConns(10)
	fmt.Println(gormDB.Name())
}
