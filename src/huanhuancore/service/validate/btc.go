package validate

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/client/rpc"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/config"
	t "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/model"
)

const (
	satoshi = 1e8
)

func getBtcTxDetail(txid string) (*t.BTC, error) {
	btcrpc := rpc.BTCRPC()

	args := make([]interface{}, 0)
	args = append(args, txid)
	args = append(args, 1)

	// btc rpc : getrawtransaction <txid> 1
	response, err := btcrpc.Call("getrawtransaction", args)
	if response != nil {
		if response.Error != nil {
			return nil, errors.New(response.Error.Message)
		}
	} else if err != nil {
		return nil, err
	}
	jsonRes, _ := json.Marshal(response.Result)
	ret := new(t.BTC)
	json.Unmarshal(jsonRes, ret)

	return ret, err
}

// getValueOutToOfficial calculate the btc to official in satoshi (string)
func getValueOutToOfficial(_txdata *t.BTC) string {

	v := int64(0)
	for _, vout := range _txdata.Vout {

		if strings.Compare(vout.ScriptPubKey.Addresses[0], config.BTCOfficial()) == 0 {
			v += int64(vout.Value * satoshi) // convert btc to satoshi
		}
	}
	return strconv.FormatInt(v, 10)
}
