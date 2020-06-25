package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&Product{})

	db.DropTableIfExists(&Product{})
	// 创建
	if !db.HasTable("products") {
		/*
		db.Create(&Product{
			Code:  "L1212",
			Price: 1000,
		})
		*/
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Product{})
	}

	// 插入值
	p := &Product{
		Code:  "L1212",
		Price: 1000,
	}
	db.Create(p)

	// 修改列
	db.Model(&Product{}).ModifyColumn("price", "bigint not null")

	// 删除列
	//db.Model(&Product{}).DropColumn("price")

	// 添加外键
	// 1st param : 外键字段
	// 2nd param : 外键表(字段)
	// 3rd param : ONDELETE
	// 4th param : ONUPDATE
	// db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")

	// 索引
	db.Model(&Product{}).AddIndex("idx_product_code", "code")
	db.Model(&Product{}).AddIndex("idx_product_code_price", "code", "price")
	db.Model(&Product{}).AddUniqueIndex("idx_product_code", "code")
	db.Model(&Product{}).AddUniqueIndex("idx_product_code_price", "code", "price")
	db.Model(&Product{}).RemoveIndex("idx_product_code_price")

	// 读取
	var product Product
	//row := db.First(&product, 1) //查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	fmt.Println(product.ID)
}


