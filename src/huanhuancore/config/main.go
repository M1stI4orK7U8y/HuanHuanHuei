package config

import (
	"fmt"
	"sync"
	"time"

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

// DBServiceName get db grpc url
func DBServiceName() string {
	return getConfig().DB.ServiceName
}

// BTCOfficial official btc address
func BTCOfficial() string {
	return getConfig().Official.Btcaddress
}

// ETHOfficial official eth address
func ETHOfficial() string {
	return getConfig().Official.Ethaddress
}

// ETHSecret secret of official eth address
func ETHSecret() string {
	return getConfig().Official.Ethsecret
}

// ETHURL eth rpc url
func ETHURL() string {
	return getConfig().ETH.Rpcurl
}

// Port port number
func Port() string {
	return getConfig().Port
}

// IP IP
func IP() string {
	return getConfig().IP
}

// Name returns worker name
func Name() string {
	return getConfig().Name
}

// ServiceName returns service name
func ServiceName() string {
	return getConfig().ServiceName
}

// ETCDHosts returns all etcd hosts address
func ETCDHosts() []string {
	return getConfig().ETCDHosts
}

// ETCDTimeout returns etcd connection timeout
func ETCDTimeout() time.Duration {
	return getConfig().ETCDTimeout * time.Second
}

// Heartbeat returns the heartbeat time to say i am alive
func Heartbeat() time.Duration {
	return getConfig().Heartbeat * time.Second
}
