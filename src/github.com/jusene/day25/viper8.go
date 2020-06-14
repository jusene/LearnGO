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
