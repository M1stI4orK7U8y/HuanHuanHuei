package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var instance *config
var once sync.Once

type config struct {
	Port string      `json:"port"`
	DB   *dbconfig   `json:"db"`
	Core *coreconfig `json:"core"`
}

type dbconfig struct {
	Grpcurl string `json:"grpcurl"`
}

type coreconfig struct {
	Grpcurl string `json:"grpcurl"`
}

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
