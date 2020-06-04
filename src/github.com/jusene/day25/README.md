## GO语言 Viper配置管理

Viper(蝰蛇)是适用于Go应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。

```
go get github.com/spf13/viper
```

Viper是适用于Go应用程序（包括Twelve-Factor App）的完整配置解决方案。

- 设置默认值
- 从JSON,TOML,YAML,HCL,envfile和JAVA properties格式的配置文件读取配置信息
- 实时监控和重新读取配置文件（可选）
- 从环境变量中读取
- 从远程配置系统（etcd或consul）读取并监控配置变化
- 从命令行参数读取配置
- 从buffer读取配置
- 显式配置值

Viper会按照下面的优先级，每个项目的优先级都高于它下面的项目:
- 显式调用Set设置值
- 命令行参数（flag）
- 环境变量
- 配置文件
- key/value存储
- 默认值

`注意`：目前Viper配置的键是大小写不敏感的。

### 设置默认值

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("Tag", map[string]string{"tag": "tags", "info": "infos"})

	fmt.Println(viper.Get("ContentDir"))
	fmt.Println(viper.Get("Tag"))
}
```

### 读取配置文件

Viper需要最少知道在哪里查找配置文件的配置。Viper支持JSON、TOML、YAML、HCL、envfile和Java properties格式的配置文件。Viper可以搜索多个路径，但目前单个Viper实例只支持单个配置文件。Viper不默认任何配置搜索路径，将默认决策留给应用程序。

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("/etc/appname") // 查找配置文件所在的路径
	viper.AddConfigPath(".") // 多次调用以添加多个搜索路径
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	} // 查找并读取配置文件
	fmt.Println(viper.Get("apiVersion"))
}
```

### 写入配置文件

- WriteConfig - 将当前的viper配置写入预定义的路径并覆盖（如果存在的话）。如果没有预定义的路径，则报错。
- SafeWriteConfig - 将当前的viper配置写入预定义的路径。如果没有预定义的路径，则报错。如果存在，将不会覆盖当前的配置文件。
- WriteConfigAs - 将当前的viper配置写入给定的文件路径。将覆盖给定的文件(如果它存在的话)。
- SafeWriteConfigAs - 将当前的viper配置写入给定的文件路径。不会覆盖给定的文件(如果它存在的话)。

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.Set("KEY", "golang")
	viper.AddConfigPath(".")
	viper.SetConfigName("test")
	viper.SetConfigType("yaml")

	// 文件不存在，会报错  会覆盖
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("%v\n", err)
	}

	// 文件不存在，会创建 不会覆盖
	if err := viper.SafeWriteConfig(); err != nil {
		fmt.Printf("%v\n", err)
	}

	// 文件不存在，会创建   会覆盖
	if err := viper.WriteConfigAs("./.config.yaml"); err != nil {
		fmt.Printf("%v\n", err)
	}

	// 文件不存在，会创建  不会覆盖
	if err := viper.SafeWriteConfigAs("./.config.yaml"); err != nil {
		fmt.Printf("%v\n", err)
	}
}
```

### 监控并重写读取配置文件

```go
package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./config.yaml")
	viper.WatchConfig()
	for {
		viper.OnConfigChange(func(in fsnotify.Event) {
			// 配置文件发生变化之后调用回调函数
			fmt.Println("Config file changed: ", in.Name, in.Op, in.String())
		})
	}
}
```

### 从io.Reader读取配置

viper预先定义了许多配置源，如文件、环境变量、标志和远程K/V存储。

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml") // 或者 viper.SetConfigType("YAML")

	// 任何需要将此配置添加到程序中的方法。
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	fmt.Println(viper.Get("name")) // 这里会得到 "steve"
}
```

### 覆盖设置

```
viper.Set("Verbose", true)
```

### 注册和使用别名

```
viper.RegisterAlias("loud", "Verbose")  // 注册别名（此处loud和Verbose建立了别名）

viper.Set("verbose", true) // 结果与下一行相同
viper.Set("loud", true)   // 结果与前一行相同

viper.GetBool("loud") // true
viper.GetBool("verbose") // true
```

### 使用环境变量





