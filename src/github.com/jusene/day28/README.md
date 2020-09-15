## GO语言 GORM

### 安装

```
go get -u github.com/jinzhu/gorm
```

驱动程序
```go
import _ "github.com/jinzhu/gorm/dialects/mysql"
// import _ "github.com/jinzhu/gorm/dialects/postgres"
// import _ "github.com/jinzhu/gorm/dialects/sqlite"
// import _ "github.com/jinzhu/gorm/dialects/mssql"
```

### CRUD

```go
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
```

### 模型定义

```go
type User struct {
    Name string `gorm:"<-:create"` // allow read and create
    Name string `gorm:"<-:update"` // allow read and update
    Name string `gorm:"<-"`        // allow read and write (create and update)
    Name string `gorm:"<-:false"`  // allow read, disable write perssion
    Name string `gorm:"->"`        // readonly (disable write perssion)
    Name string `gorm:"->;<-:create"` // allow read and create
    Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
    Name string `gorm:"-"`         // ignore this field when write and read
}
```

GORM 默认使用 `CreatedAt`,'UpdatedAt' 去追踪创建和更新时间，如果不同的字段名称，需要tag `autoCreateTime`,'autoUpdateTime'

```go
type User struct {
    CreatedAt time.Time // Set to current time if it is zero on creating
    UpdatedAt int  // Set to current unix seconds on updaing or if it is zero on creating
    Updated int64  `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
    Updated int64  `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
    Created int64  `gorm:"autoCreateTime"`  // Use unix second as creating time
}
```

嵌入结构

```go
type User struct {
    gorm.Model
    Name string
}

// equals
type User struct {
    ID uint   `gorm:"primaryKey"`
    CreateAt time.Time
    UpdateAt time.Time
    DeleteAt gorm.DeleteAt `gorm:"index"`
    Name string
}
```

```go
type Author struct {
    Name string
    Email string
}

type Blog struct {
    ID int
    Author Author `gorm:"embedded"`
    Upvotes int32
}

// equals
type Blog struct {
    ID int64
    Name string
    Email string
    Upvote int32
}
```

并且可以使用tag `embeddedPrefix` 去添加添加内嵌前缀字段名

```go
type Blog struct {
    ID int
    Author Author `gorm:"embedded;embeddedPrefix:author_"`
    Upvotes int32
}

// equals   
type Blog struct {
    ID int64
    AuthorName string
    AuthorEmail string
    Upvotes int32
}
```

## 字段标签

| 标签名称 | 描述 | 
| -----| ----- | 
| column | 列数据库名称 | 
| type | 列数据类型，使用兼容的通用类型，例如：bool、int、uint、float、string、time、bytes，适用于所有数据库，可以与其他标记一起使用，如not null、size、autoIncrement…在使用指定的数据库数据类型时，也支持varbinary（8）等指定的数据库数据类型，它需要是一个完整的数据库数据类型，例如：MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT | 
| size | 指定列数据的大小/长度 |
| 
