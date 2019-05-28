package main

import (
	"net/http"
	"strings"
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

func TestGetRecords(t *testing.T) {
	response, err := http.Post("http://127.0.0.1:"+config.Port()+"/api/v1/record",
		"application/x-www-form-urlencoded",
		strings.NewReader("ids=123456&ids=456789"))

	if err != nil {
		t.Errorf("Failed get records: %s", err.Error())
	} else {
		t.Logf("Success get records %v", response)
	}
}
