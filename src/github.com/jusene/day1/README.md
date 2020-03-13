## Go语言

> Go语言官网：https://golang.org/dl/

> Go语言官方镜像站：https://golang.google.cn/dl/

## 系统级别的环境变量

- $GOROOT: 表示GO语言环境在计算机上的安装位置
- $GOPATH: GO语言的工作目录，可以多个，类似工作空间概念

example:
```
export GOROOT=/usr/local/go/
export GOPATH=$HOME/Workspace/Go:$HOME/other/Go
export PATH=$PATH:$GOROOT/bin:${GOPATH//://bin:}/bin
```

$GOPATH可以设置多个，但是包管理安装时默认使用第一个$GOPATH的值作为下载目录，建议$GOPATH第一个值设置为常用的全局工作目录。

## 基本命令及使用

- build： 编译源代码包和依赖
- clean： 删除对象文件
- doc： 现实Go语言环境的变量信息
- bug: bug提交程序
- fix：修复程序
- fmt: 格式化源代码中的代码
- generate: 通过扫描Go源码中的go:generate注释来识别要运行的常规命令
- get: 下载并安装指定的包与依赖
- install: 编译并安装指定的包与依赖
- list: 打印指定源代码的信息
- run: 编译并运行Go程序
- test: 测试一个源代码
- tool: 运行一个指定的go tool
- version: 打印输出Go环境版本
- vet: 检查源代码包中可能出现的错误

## GOROOT结构

```
drwxr-xr-x.  2 root root   258 Dec 14 18:36 api # 存放了包含公开的变量、常量、函数等api的列表
-rw-r--r--.  1 root root 55284 Dec 14 18:36 AUTHORS # 所有参与go语言开发的人员名单
drwxr-xr-x.  2 root root    42 Dec 14 18:48 bin # 文件夹主要用于存储标准命令执行文件
-rw-r--r--.  1 root root  1339 Dec 14 18:36 CONTRIBUTING.md # 为go贡献代码的说明
-rw-r--r--.  1 root root 71070 Dec 14 18:36 CONTRIBUTORS # 所有贡献者名单
drwxr-xr-x.  8 root root  4096 Dec 14 18:36 doc # 文件夹存放了标准库的文档，使用godoc -http=:6060启动文档服务
-rw-r--r--.  1 root root  5686 Dec 14 18:36 favicon.ico # go语言标志
drwxr-xr-x.  3 root root    18 Dec 14 18:36 lib #存放一些特殊的库文件
-rw-r--r--.  1 root root  1479 Dec 14 18:36 LICENSE # 工语言的开源协议
drwxr-xr-x. 15 root root   202 Dec 14 18:48 misc #辅助工具和说明
-rw-r--r--.  1 root root  1303 Dec 14 18:36 PATENTS # go语言专利说明
drwxr-xr-x.  9 root root   167 Dec 14 18:48 pkg # 用于存放编译go语言标准库生成的文件
-rw-r--r--.  1 root root  1607 Dec 14 18:36 README.md # 说明文件
-rw-r--r--.  1 root root    26 Dec 14 18:36 robots.txt # 禁止搜索引擎索引本地启动的go文档
drwxr-xr-x. 46 root root  4096 Dec 14 18:36 src # 存放go语言自己的源代码
drwxr-xr-x. 22 root root  8192 Dec 14 18:48 test # 存放测试相关的文件
```

## GOPATH结构

- bin: 存放go install生成的可执行文件
- pkg: 存放go编译生成的文件
- src: 存放我们开发的go源代码

## GOPROXY

Go1.14版本之后，都推荐使用go mod模式来管理依赖环境了，也不再强制我们把代码必须写在GOPATH下面的src目录了，你可以在你电脑的任意位置编写go代码。

默认GoPROXY配置是：GOPROXY=https://proxy.golang.org,direct，由于国内访问不到https://proxy.golang.org，所以我们需要换一个PROXY，这里推荐使用https://goproxy.io或https://goproxy.cn。

``` 
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

## go build

- -a: 强制重新编译所有相关的go语言源代码包
- -n: 检查执行命令过程中实际会用到的命令，只输出要执行的命令，实际上并不会执行
- -p: 构建或测试指定并行运行的程序数量，默认值是可用cpu个数
- -race: 检查数据竞争问题，并发编程中会使用，只支持amd64
- -v: 打印被编译的包
- -work: 指定编译缓存工作目录
- -x: 与上面的-n类似，打印并且执行这些编译命令
- -o: 指定go编译出的文件名

## go clean

编译后需要清理当前源代码包和关联源码包里编译生成的文件，可以使用go clean快速清理，添加-r选项可以清理导入代码包中的构建缓存

## go get

- -d: 只执行下载运动，不执行安装动作
- -f: 不检查已下载代码包的导入路径，需要与-u选项配合使用
- -fix: 下载代码包后先执行fix动作（代码修复兼容问题），然后再进行编译安装
- -insecure: 允许get命令使用不安全的http协议下载代码包
- -t: 让get命令同时下载安装指定的代码包的测试源码文件中的依赖代码包
- -u: 更新已有代码包与依赖包。在默认情况下，get不会下载本地已存在的代码包，需要使用此选项更新代码包
- -v: 打印要下载安装的代码包名称
- -x: 显示下载安装需要执行的具体命令

## gofmt

- -l: 显示需要格式化的文件
- -w: 不将格式化结果打印到标准输出，而是直接保存到文件
- -r: 添加形如“<原始内容> -> <替换内容>”的重写规则，方便批量替换
- -s: 简化文件中的代码
- -d: 显示格式化前后的不同（不写入文件），默认是false
- -e: 打印所有的语法错误，默认每行的前十个错误
- -cpuprofile: 支持调试模式，将相应的cpufile写入指定文件

## go install

go install表示安装的意思，它先编译源代码得到可执行文件，然后将可执行文件移动到GOPATH的bin目录下。因为我们的环境变量中配置了GOPATH下的bin目录，所以我们就可以在任意地方直接执行可执行文件了。

## 异构

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
``

人生苦短 Let's Go