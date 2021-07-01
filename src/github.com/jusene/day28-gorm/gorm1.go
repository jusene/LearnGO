package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 默认情况下，GORM在事务中执行单个创建、更新、删除操作，以确保数据库数据的完整性
		// 可以通过将“SkipDefaultTransaction”设置为true来禁用它
		SkipDefaultTransaction: false,
		// 命名策略表，列命名策略
		NamingStrategy: nil,
		// Logger
		Logger: nil,
		// NowFunc创建新时间戳时要使用的函数
		NowFunc: nil,
		// DryRun生成sql而不执行
		DryRun: false,
		// PrepareStmt在缓存语句中执行给定的查询
		PrepareStmt: false,
		// 禁止自動ping
		DisableAutomaticPing: false,
		// 迁移时禁用ForeignKeyConstraintWhen迁移
		DisableForeignKeyConstraintWhenMigrating: false,
		// 允許全局更新
		AllowGlobalUpdate: false,
		// 子句生成器
		ClauseBuilders: nil,
		// 连接池db连接池
		ConnPool:  nil,
		Dialector: nil,
		Plugins:   nil,
	})

	if err != nil {
		panic("failed to connect database")
	}

	// 迁移架构
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "D42", Price: 100})

	// 读取
	var product Product
	// SELECT * FROM `products` WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1
	db.First(&product, 1)
	fmt.Println(product.Code)
	db.First(&product, "code = ?", "D42")
	fmt.Println(product.ID)

	// 更新
	db.Model(&product).Where("code = ?", "F43").Update("Price", 200)
	db.Model(&product).Updates(Product{Price: 300, Code: "F43"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 400})

	// 删除
	db.Delete(&product, 1)
}
