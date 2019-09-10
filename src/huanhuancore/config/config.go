package config

import "time"

type config struct {
	BTC      *btcconfig `json:"btc"`
	ETH      *ethconfig `json:"eth"`
	DB       *dbconfig  `json:"db"`
	Official *official  `json:"official"`
	IP       string     `json:"ip"`
	Port     string     `json:"port"`

	Name        string        `json:"name"`
	ServiceName string        `json:"servicename"`
	ETCDHosts   []string      `josn:"etcdhosts"`
	ETCDTimeout time.Duration `json:"etcdtimeout"`
	Heartbeat   time.Duration `json:"heartbeat"`
}

type btcconfig struct {
	Rpcurl   string `json:"rpcurl"` // httpurl:port
	Username string `json:"username"`
	Password string `json:"password"`
}

type ethconfig struct {
	Rpcurl string `json:"rpcurl"` // httpurl : port
}

type dbconfig struct {
	ServiceName string `json:"servicename"`
}

type official struct {
	Btcaddress string `json:"btcaddress"`
	Ethaddress string `json:"ethaddress"`
	Ethsecret  string `json:"ethsecret"`
}
