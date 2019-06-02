package btc

import (
	"strings"

	token "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/model"
)

// CheckInputTx validate input tx
func CheckInputTx(input *token.BTC) bool {
	if input == nil {
		return false
	}

	// check tx exists
	if input.Confirmations < 0 { // confirmations : -1 not in mainchain
		return false
	}

	// check vout value to official is not 0
	if strings.Compare(GetValueOutToOfficial(input), "0") == 0 {
		return false
	}
	return true
}
