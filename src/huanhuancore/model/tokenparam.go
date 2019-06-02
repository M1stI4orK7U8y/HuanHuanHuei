package token

import "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/token"

const (
	// Satoshi 1e8 satoshi = 1btc
	Satoshi = 1e8
)

// Fullname full name of each token
var Fullname = map[token.TokenType]string{
	token.TokenType_BTC: "bitcoin",
	token.TokenType_ETH: "ethereum",
}

// ShortName short name of each token
var ShortName = map[token.TokenType]string{
	token.TokenType_BTC: "btc",
	token.TokenType_ETH: "eth",
}

// Decimal decimals of each token
var Decimal = map[token.TokenType]int32{
	token.TokenType_BTC: 8,
	token.TokenType_ETH: 18,
}
