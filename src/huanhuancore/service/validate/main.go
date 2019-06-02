package validate

import (
	"strings"

	t "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/token"
	token "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/model"
)

// CheckInputTx validate input tx
func CheckInputTx(fromtoken t.TokenType, input interface{}) bool {
	if input == nil {
		return false
	}

	switch fromtoken {
	case t.TokenType_BTC:

		// check tx exists
		if input.(*token.BTC).Confirmations < 0 { // confirmations : -1 not in mainchain
			return false
		}

		// check vout value to official is not 0
		if strings.Compare(GetValueOutToOfficial(input.(*token.BTC)), "0") == 0 {
			return false
		}
		return true
	case t.TokenType_ETH:
		return true
	default:
		return false
	}
}
