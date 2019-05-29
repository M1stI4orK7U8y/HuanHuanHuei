package rpc

import (
	"encoding/base64"

	"github.com/ybbus/jsonrpc"
	cf "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/config"
)

func btcrpc() jsonrpc.RPCClient {
	return jsonrpc.NewClientWithOpts(cf.BTCURL(), &jsonrpc.RPCClientOpts{
		CustomHeaders: map[string]string{
			"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(cf.BTCUser()+":"+cf.BTCPassword())),
		},
	})
}
