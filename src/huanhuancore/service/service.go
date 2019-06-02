package service

import (
	"encoding/json"
	"errors"
	"math"
	"math/big"
	"net/http"
	"time"

	//model
	rd "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/record"
	t "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/token"
	huanhuan "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"
	token "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/model"

	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/service/btc"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/service/eth"
)

// Service service for db operation
type Service struct{}

// using coingecko api
// rateapi: https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=btc"
const apiurl = `https://api.coingecko.com/api/v3/simple/price?`

func getrate(from, to t.TokenType) float64 {
	path := apiurl + "ids=" + token.Fullname[from] + "&vs_currencies=" + token.ShortName[to]
	response, err := http.Get(path)
	if err != nil {
		return 0
	}
	var target interface{}

	json.NewDecoder(response.Body).Decode(&target)

	// response format
	// {"ethereum":{"btc":0.0312016}}
	// get the value of vs currency
	return target.(map[string]interface{})[token.Fullname[from]].(map[string]interface{})[token.ShortName[to]].(float64)
}

func checkInputTx(input *huanhuan.HuanHuanRequest, intx interface{}) error {
	switch input.From {
	case t.TokenType_BTC:
		if btc.CheckInputTx(input.From, intx) == false {
			return errors.New("check input tx error")
		}
	case t.TokenType_ETH:
	default:
		return errors.New("token not support")
	}
	return nil
}

func validReceiver(tt t.TokenType, address string) bool {

	if len(address) == 0 {
		return false
	}

	switch tt {
	case t.TokenType_BTC:
		fallthrough
	case t.TokenType_ETH:
	default:
		return false
	}

	return true
}

func getTxDetail(tt t.TokenType, txid string) interface{} {
	switch tt {
	case t.TokenType_BTC:
		ret, err := btc.GetBtcTxDetail(txid)
		if err != nil {
			return ret
		}
	case t.TokenType_ETH:
	}
	return nil
}

func createRecord(from, to t.TokenType, receiver string, tx interface{}) *rd.Record {
	switch from {
	case t.TokenType_BTC:
		return createBTCRecord(to, receiver, tx.(*token.BTC))
	}
	return nil
}

func createBTCRecord(totoken t.TokenType, receiver string, tx *token.BTC) *rd.Record {

	tnow := time.Now()
	exrate := getrate(t.TokenType_BTC, totoken)

	ret := new(rd.Record)
	ret.Id = tx.Txid

	// from
	ret.FromToken.Txhash = tx.Txid
	// get vin[0] address as sender address
	vintx, _ := btc.GetBtcTxDetail(tx.Vin[0].Txid)
	ret.FromToken.Address = vintx.Vout[tx.Vin[0].Vout].ScriptPubKey.Addresses[0]
	ret.FromToken.TokenType = t.TokenType_BTC
	ret.FromToken.TokenDecimal = token.Decimal[t.TokenType_BTC]
	ret.FromToken.TokenValue = btc.GetValueOutToOfficial(tx)

	// to
	ret.ToToken.TokenType = totoken
	ret.ToToken.Address = receiver
	ret.ToToken.TokenDecimal = token.Decimal[totoken]
	// calculate
	ret.ToToken.TokenValue = calculateTargetValue(exrate, t.TokenType_BTC, totoken, ret.FromToken.TokenValue)
	// to.txhash not defined

	ret.Exrate = exrate
	ret.StatusCode = rd.StatusCode_PENDING
	ret.StatusTime = &rd.StatusTime{PendingTime: tnow.UTC().Unix()}

	return ret

}

func calculateTargetValue(exrate float64, from, to t.TokenType, fromvalue string) string {
	ori, _ := new(big.Float).SetString(fromvalue)                                                            // big number
	little := new(big.Float).Quo(ori, big.NewFloat(math.Pow10(int(token.Decimal[from]))))                    // to float
	exchangelittle := new(big.Float).Mul(little, big.NewFloat(exrate))                                       // calculate
	return new(big.Float).Mul(exchangelittle, big.NewFloat(math.Pow10(int(token.Decimal[to])))).Text('f', 0) // to target big number
}

func sendtoreceiver(to t.TokenType, address, value string) (string, error) {
	switch to {
	case t.TokenType_BTC:
	case t.TokenType_ETH:
		return eth.SendToAddress(address, value)
	}
	return "", errors.New("sendtoreceiver: type not support")
}
