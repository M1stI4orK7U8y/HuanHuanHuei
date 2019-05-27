package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var instance *config
var once sync.Once

type config struct {
	DB *dbconfig `json:"db"`
}

type dbconfig struct {
	Grpcurl string `json:"grpcurl"`
}

// getConfig config Instance
func getConfig() *config {
	once.Do(func() {
		viper.SetConfigName("config") // 设置配置文件名 (不带后缀)
		viper.AddConfigPath(".")      // 第一个搜索路径
		err := viper.ReadInConfig()   // 读取配置数据
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}
		instance = new(config)
		viper.Unmarshal(instance) // 将配置信息绑定到结构体上
	})
	return instance
}

// DBGrpcURL get db grpc url
func DBGrpcURL() string {
	return getConfig().DB.Grpcurl
}
