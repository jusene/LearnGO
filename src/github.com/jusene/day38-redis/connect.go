package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()

	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := initClient(); err != nil {
		panic(err)
	}

	err := rdb.Set("test", "2222", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get("test").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("VALUE", val)

	val, err = rdb.Get("test11").Result()
	if err == redis.Nil {
		panic("ket not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("value", val)
	}
}
