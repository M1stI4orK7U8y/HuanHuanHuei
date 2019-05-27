package main

import (
	"net/http"
	"testing"

	"gitlab.com/packtumi9722/huanhuanhuei/src/entry/config"
)

func TestGetRecord(t *testing.T) {
	id := "4564"
	response, err := http.Get("http://127.0.0.1:" + config.Port() + "/api/v1/record/" + id)
	if err != nil {
		t.Errorf("Failed get record: %s", err.Error())
	} else {
		t.Logf("Success get record %s : %v", id, response)
	}

	id = "123456"
	response, err = http.Get("http://127.0.0.1:" + config.Port() + "/api/v1/record/" + id)
	if err != nil {
		t.Errorf("Failed get record: %s", err.Error())
	} else {
		t.Logf("Success get record %s : %v", id, response)
	}
}
