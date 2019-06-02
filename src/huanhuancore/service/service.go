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

	"github.com/ethereum/go-ethereum/core/types"
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
		if btc.CheckInputTx(intx.(*token.BTC)) == false {
			return errors.New("btc check input tx error")
		}
	case t.TokenType_ETH:
		if eth.CheckInputTx(intx.(*types.Transaction)) == false {
			return errors.New("eth check input tx error")
		}
	default:
		return errors.New("token not support")
	}
	return nil
}

func getTxDetail(tt t.TokenType, txid string) (interface{}, error) {
	switch tt {
	case t.TokenType_BTC:
		return btc.GetTxDetail(txid)
	case t.TokenType_ETH:
		ret, _, err := eth.GetTxDetail(txid)
		return ret, err
	}
	return nil, nil
}

func createRecord(from, to t.TokenType, receiver string, tx interface{}) *rd.Record {
	switch from {
	case t.TokenType_BTC:
		return createBTCRecord(to, receiver, tx.(*token.BTC))
	case t.TokenType_ETH:
		return createETHRecord(to, receiver, tx.(*types.Transaction))
	}
	return nil
}

func createBTCRecord(totoken t.TokenType, receiver string, tx *token.BTC) *rd.Record {
	tnow := time.Now()
	exrate := getrate(t.TokenType_BTC, totoken)

	ret := new(rd.Record)
	ret.Id = tx.Txid

	// from
	ret.FromToken = new(rd.TokenDetail)
	ret.FromToken.Txhash = tx.Txid
	// get vin[0] address as sender address
	vintx, _ := btc.GetTxDetail(tx.Vin[0].Txid)
	ret.FromToken.Address = vintx.Vout[tx.Vin[0].Vout].ScriptPubKey.Addresses[0]
	ret.FromToken.TokenType = t.TokenType_BTC
	ret.FromToken.TokenDecimal = token.Decimal[t.TokenType_BTC]
	ret.FromToken.TokenValue = btc.GetValueOutToOfficial(tx)

	// to
	ret.ToToken = new(rd.TokenDetail)
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

func createETHRecord(totoken t.TokenType, receiver string, tx *types.Transaction) *rd.Record {
	tnow := time.Now()
	exrate := getrate(t.TokenType_ETH, totoken)

	ret := new(rd.Record)
	ret.Id = tx.Hash().Hex()

	// from
	ret.FromToken = new(rd.TokenDetail)
	ret.FromToken.Txhash = tx.Hash().Hex()
	// sign the sender
	signer := types.NewEIP155Signer(tx.ChainId())
	sender, _ := signer.Sender(tx)
	ret.FromToken.Address = sender.Hex()
	ret.FromToken.TokenType = t.TokenType_ETH
	ret.FromToken.TokenDecimal = token.Decimal[t.TokenType_ETH]
	ret.FromToken.TokenValue = tx.Value().Text(10)

	// to
	ret.ToToken = new(rd.TokenDetail)
	ret.ToToken.TokenType = totoken
	ret.ToToken.Address = receiver
	ret.ToToken.TokenDecimal = token.Decimal[totoken]
	// calculate
	ret.ToToken.TokenValue = calculateTargetValue(exrate, t.TokenType_ETH, totoken, ret.FromToken.TokenValue)
	// to.txhash not defined

	ret.Exrate = exrate
	ret.StatusCode = rd.StatusCode_PENDING
	ret.StatusTime = &rd.StatusTime{PendingTime: tnow.UTC().Unix()}

	return ret
}

func calculateTargetValue(exrate float64, from, to t.TokenType, fromvalue string) string {
	ori, _ := new(big.Float).SetString(fromvalue)                                         // big number
	little := new(big.Float).Quo(ori, big.NewFloat(math.Pow10(int(token.Decimal[from])))) // to float
	exchangelittle := new(big.Float).Mul(little, big.NewFloat(exrate))                    // calculate
	// convert to string and use this string to generate a big.Float
	// avoid float trailing problem
	convstr, _ := new(big.Float).SetString(exchangelittle.String())
	exchangebig := new(big.Float).Mul(convstr, big.NewFloat(math.Pow10(int(token.Decimal[to])))).Text('f', 0) // to target big number
	sendValue, _ := new(big.Int).SetString(exchangebig, 10)
	return sendValue.String()
}

func sendtoreceiver(to t.TokenType, address, value string) (string, error) {
	switch to {
	case t.TokenType_BTC:
		return btc.SendToAddress(address, value)
	case t.TokenType_ETH:
		return eth.SendToAddress(address, value)
	}
	return "", errors.New("sendtoreceiver: type not support")
}
