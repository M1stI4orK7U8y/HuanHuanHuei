package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var instance *config
var once sync.Once

type config struct {
	IP   string      `json:"ip"`
	Port string      `json:"port"`
	DB   *dbconfig   `json:"db"`
	Core *coreconfig `json:"core"`

	ETCDHosts   []string      `josn:"etcdhosts"`
	ETCDTimeout time.Duration `json:"etcdtimeout"`
	Heartbeat   time.Duration `json:"heartbeat"`
}

type dbconfig struct {
	ServiceName string `json:"servicename"`
}

type coreconfig struct {
	ServiceName string `json:"servicename"`
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
