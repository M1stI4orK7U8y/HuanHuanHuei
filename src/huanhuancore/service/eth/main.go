package eth

import (
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/config"
)

// CheckInputTx validate input tx
func CheckInputTx(input *types.Transaction) bool {
	if strings.Compare(strings.ToUpper(input.To().Hex()), strings.ToUpper(config.ETHOfficial())) != 0 {
		return false
	}
	return true
}
