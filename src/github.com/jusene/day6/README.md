## Go语言复合数据类型

### 映射（map）

Go语言中提供的映射关系容器为map，其内部使用散列表（hash）实现。

map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用

#### map定义

```go
mao[KeyType]ValueType
```

map类型的变量默认初始值为nil，需要使用make()函数来分配内存

```go
make(map[KeyType]ValueType, [cap])
```

