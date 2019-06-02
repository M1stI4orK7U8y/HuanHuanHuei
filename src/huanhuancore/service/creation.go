package service

import (
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	rd "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/record"
	t "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/token"
	token "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/model"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/service/btc"
)

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
