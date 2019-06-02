package rpc

import (
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	cf "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/config"
)

var instance *ethclient.Client
var once sync.Once

// GetEthInstance : getEthConn Instance
func GetEthInstance() *ethclient.Client {
	once.Do(func() {
		conn, _ := Connect()
		instance = &jfbasetype.EthConnection{Client: conn}
	})
	return instance
}

func connect()(*ethclient.Client, error) {{
	conn, err := ethclient.Dial(cf.ETHURL())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
