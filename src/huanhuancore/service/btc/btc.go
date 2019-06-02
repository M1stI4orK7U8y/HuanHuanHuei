package btc

import (
	"encoding/json"
	"errors"
	"math"
	"math/big"
	"strconv"
	"strings"

	t "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/token"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/client/rpc"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/config"
	token "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/model"
)

// GetTxDetail get btc tx detail
func GetTxDetail(txid string) (*token.BTC, error) {
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
	ret := new(token.BTC)
	json.Unmarshal(jsonRes, ret)

	return ret, err
}

// GetValueOutToOfficial calculate the btc to official in satoshi (string)
func GetValueOutToOfficial(_txdata *token.BTC) string {

	v := int64(0)
	for _, vout := range _txdata.Vout {

		if strings.Compare(vout.ScriptPubKey.Addresses[0], config.BTCOfficial()) == 0 {
			v += int64(vout.Value * token.Satoshi) // convert btc to satoshi
		}
	}
	return strconv.FormatInt(v, 10)
}

// SendToAddress send btc to address
func SendToAddress(address string, amount string) (string, error) {
	btcrpc := rpc.BTCRPC()

	args := make([]interface{}, 0)
	args = append(args, address)

	satoshi, _ := new(big.Float).SetString(amount)
	btc, _ := new(big.Float).Quo(satoshi, big.NewFloat(math.Pow10(int(token.Decimal[t.TokenType_BTC])))).Float64()

	args = append(args, btc)

	response, err := btcrpc.Call("sendtoaddress", args)
	if response != nil {
		if response.Error != nil {
			return "", errors.New(response.Error.Message)
		}
	} else if err != nil {
		return "", err
	}
	return response.Result.(string), nil
}
