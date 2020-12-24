## GO语言 GORM

## GORM V1
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

## GORM V2

```go
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
| primaryKey | 指定列作为主键 |
| unique | 指定列作为唯一键 |
| default | 指定列默认值 |
| precision | 指定列的精度 |
| scale | 指定列比例 |
| not null | 指定列not null |
| autoIncrement | 指定列autoincrement |
| embedded | 嵌入字段 |
| embeddedPrefix | 嵌入字段前缀 |
| autoCreateTime | 跟踪创建时的当前时间，对于int字段，它将跟踪unix秒，使用value nano/milli跟踪unix nano / milli秒，例如：autoCreateTime:nano |
| autoUpdateTime | 在创建/更新时跟踪当前时间，对于int字段，它将跟踪unix秒，使用value nano/milli跟踪unix nano / milli秒，例如：autoUpdateTime:milli |
| index | 使用选项创建索引 |
| uniqueIndex | 与相同index，但创建唯一索引 |
| check | 创建检查约束，例如check:age > 13 |
| <- | 设置字段的写权限，<-:create仅创建字段，<-:update仅更新字段，<-:false无写权限，<-创建和更新权限 |
| -> | 设置字段的读取权限，->:false没有读取权限 |
| - | 忽略这个字段, 没有读写权限 |

### 连接数据库

```go
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
```

### CRUD

#### Create