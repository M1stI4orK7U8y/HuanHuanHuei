package service

import (
	"testing"

	"gitlab.com/packtumi9722/huanhuanhuei/src/database/model/token"
)

func TestGetRate(t *testing.T) {
	ret := getrate(token.TokenType_BTC, token.TokenType_ETH)
	if ret > 0 {
		t.Logf("Success Test getrate : rate = %f", ret)
	} else {
		t.Error("Fail Test getrat")
	}
}
