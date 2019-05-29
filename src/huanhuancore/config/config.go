package config

type config struct {
	BTC *btcconfig `json:"btc"`
	DB  *dbconfig  `json:"db"`
}

type btcconfig struct {
	Rpcurl   string `json:"rpcurl"` // ip:port
	Username string `json:"username"`
	Password string `json:"password"`
}

type dbconfig struct {
	Grpcurl string `json:"grpcurl"`
}

type official struct {
	Btcaddress string `json:"btcaddress"`
	Ethaddress string `json:"ethaddress"`
}
