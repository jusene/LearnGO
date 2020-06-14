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

viper支持环境变量，有五种方法与ENV协作：
- AutomaticEnv()
- BindEnv(string...) : error
- SetEnvPrefix(string)
- SetEnvKeyReplacer(string) *strings.Replacer
- AllowEmptyEnv(bool)

Viper提供了一种机制来确保ENV变量是惟一的，通过`SetEnvPrefix`,可以告诉Viper在读取环境变量时使用前缀。`BindEnv`和`AutomaticEnv`都将使用这个前缀。

`BindEnv`使用一个或两个参数，第一个参数是键的名称，第二个是环境变量的名称。环境变量区分大小写，如果没有提供ENV变量名，那么Viper将自动假设ENV变量与以下格式匹配：前缀+"_"+键名全部大小写。当你显式提供ENV变量名（第二个参数），不会自动添加前缀。

在使用ENV变量时，需要注意的一件重要事情是，每次访问该值得时候都将读取它，Viper在调用`BindEnv`时不固定该值。

`AutomaricEnv`与`SetEnvPrefix`结合使用时，调用时，Viper会在发出`viper.Get`请求时随时检查环境变量，它将应用以下规则，将检查环境变量的名称是否与键匹配（如果设置了EnvPrefix）。

`SetEnvKeyReplacer`允许你使用`strings.Replacer`对象在一定程度上重写ENV键。如果你希望在`Get()`调用使用`-`或者其他符号，但是环境变量里使用`_`分隔符，那么这个功能非常有用，或者可以使用带有`NewWithOptions`工厂函数`EnvKeyReplacer`,与SetEnvKeyReplacer不同，它接受`StringReplacer`函数。

默认情况下，空环境变量被认定是未设置的，并将返回到下一个配置源，若要将空环境变量视为已设置，使用`AllowEmptyEnv`方法。

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.SetEnvPrefix("spf")
	viper.BindEnv("id")
	viper.BindEnv("name", "name")

	os.Setenv("SPF_ID", "12")
	os.Setenv("NAME", "JUSENE")
	fmt.Println(viper.Get("id"))
	fmt.Println(viper.Get("name"))
}
```

### 使用Flags

```go
package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	flag.Int("flagname", 1234, "help message for flagname")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	i := viper.GetInt("flagname") // 从 viper 检索值
	fmt.Println(i)
}
```

### 远程Key/Value存储支持

在viper中启用远程支持，需要在代码中匿名导入`viper/remote`。

```
import _ "github.com/spf13/viper/remote"
```

viper将读取从Key/Value存储中路径检索到的配置字符串，viper加载配置的优先级：磁盘上的配置文件>命令行标志位>环境变量>远程key/value>默认值。

viper使用crypt从K/V存储中检索配置，这意味着如果你有正确的gpg密匙，你可以将配置值加密存储并自动解密。加密是可选的。

crypt有一个命令行助手，你可以使用它将配置放入K/V存储中。crypt默认使用在http://127.0.0.1:4001的etcd
```
go get github.com/bketelsen/crypt/bin/crypt
crypt set -plaintext /config/hugo.json /Users/hugo/settings/config.json
```

```
crypt get -plaintext /config/hugo.json
```

#### 远程key/value存储示例-未加密

- etcd
```
viper.AddRemoteProvider("etcd", "http://127.0.0.1:4100", "/config/hugo.json")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
err := viper.ReadRemoteConfig()
```

- consul
consul key/value存储中设置一个key保存所需要的JSON值
```json
{
    "port": 8080,
    "hostname": "jusene.com",
}
```

```
viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
viper.SetConfigType("json")
err := viper.ReadRemoteConfig()

fmt.Println(viper.Get("port"))
fmt.Println(viper.Get("hostname"))
```

- Firestore
```
viper.AddRemoteProvider("firestore", "google-cloud-project-id", "collection/document")
viper.SetConfigType("json") // 配置的格式: "json", "toml", "yaml", "yml"
err := viper.ReadRemoteConfig()
```

#### 远程key/value存储示例-加密

```
viper.AddSecureRemoteProvider("etcd", "http://127.0.0.1:4100", "/config/hugo.json", "/etc/secrets/mykeyring.gpg")
viper.SetConfigType("json") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
err := viper.ReadRemoteConfig()
```

#### 监控etcd中的更改-未加密

```go
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type conf struct {
	host string
	port int
}

func main() {
	var runtime_viper = viper.New()

	runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
	runtime_viper.SetConfigType("yaml") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"

	// 第一次从远程读取配置
	err := runtime_viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}

	runtime_conf := new(conf)

	// 反序列化
	runtime_viper.Unmarshal(&runtime_conf)

	// 开启一个单独goroutine一直监控远程的变更
	go func() {
		for {
			time.Sleep(time.Second * 5)

			// 目前只支持了etcd支持
			err := runtime_viper.WatchRemoteConfig()
			if err != nil {
				log.Errorf("unable to read remote config: %v", err)
				continue
			}

			// 将新配置反序列化到我们运行时的配置结构体中。你还可以借助channel实现一个通知系统更改的信号
			runtime_viper.Unmarshal(&runtime_conf)
		}
	}()
}
```

### 从viper获取值

根据类型获取值，存在以下功能和方法:
- Get(key string): interface{}
- GetBool(key string): bool
- GetFloat64(key string): float64
- GetInt(key string): int
- GetIntSlice(key string): []int
- GetString(key string): string
- GetStringMap(key string): map[string]interface{}
- GetStringMapString(key string): map[string]string
- GetStringSlice(key string): []string
- GetTime(key string): time.Time
- GetDuration(key string): time.Duration
- IsSet(key string): bool
- AllSettings(): map[string]interface{}

需要认识到的一件重要事情是，每一个Get方法在找不到值的时候都会返回零值。为了检查给定的键是否存在，提供了IsSet()方法。

