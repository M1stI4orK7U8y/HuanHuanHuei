package service

import "testing"

func TestGetRate(t *testing.T) {
	ret := getrate("btc", "eth")
	if ret > 0 {
		t.Logf("Success Test getrate : rate = %f", ret)
	} else {
		t.Error("Fail Test getrat")
	}
}
