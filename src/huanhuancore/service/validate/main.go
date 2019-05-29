package validate

import (
	"strings"

	r "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"
	t "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/model"
)

// ValidateInputTx validate input tx
func CheckInputTx(input *r.HuanHuanRequest) bool {
	if input == nil {
		return false
	}

	switch input.From {
	case t.TokenType_BTC:
		tx, err := getBtcTxDetail(input.FromTxid)

		// check tx exists
		if err != nil || tx.Confirmations < 0 { // confirmations : -1 not in mainchain
			return false
		}

		// check vout value to official is not 0
		if strings.Compare(getValueOutToOfficial(tx), "0") == 0 {
			return false
		}
		return true
	case t.TokenType_ETH:
		return true
	default:
		return false
	}
}
