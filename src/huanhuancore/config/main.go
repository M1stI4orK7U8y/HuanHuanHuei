package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var instance *config
var once sync.Once

// getConfig config Instance
func getConfig() *config {
	once.Do(func() {
		viper.SetConfigName("config") // config file name w/o extension
		viper.AddConfigPath(".")      // config file path
		err := viper.ReadInConfig()   // read config
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}
		instance = new(config)
		viper.Unmarshal(instance) // load config to config object
	})
	return instance
}

// BTCURL btc rpc url
func BTCURL() string {
	return getConfig().BTC.Rpcurl
}

//BTCUser btc rpc user
func BTCUser() string {
	return getConfig().BTC.Username
}

//BTCPassword btc rpc pass
func BTCPassword() string {
	return getConfig().BTC.Password
}

// DBGrpcURL get db grpc url
func DBGrpcURL() string {
	return getConfig().DB.Grpcurl
}
