## GO语言 gin-swagger

### 安装 gin-swagger

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

在工作目录创建一个`main.go`

生成swagger文档
```
swag.exe init
```

下载gin-swagger中间件
```
go get -u github.com/swaggo/gin-swagger // gin-swagger middleware
go get -u github.com/swaggo/files // swagger embed files
```



