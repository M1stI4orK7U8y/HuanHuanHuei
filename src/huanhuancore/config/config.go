package config

type config struct {
	BTC      *btcconfig `json:"btc"`
	ETH      *ethconfig `json:"eth"`
	DB       *dbconfig  `json:"db"`
	Official *official  `json:"official"`
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
	Grpcurl string `json:"grpcurl"`
}

type official struct {
	Btcaddress string `json:"btcaddress"`
	Ethaddress string `json:"ethaddress"`
	Ethsecret  string `json:"ethsecret"`
}
