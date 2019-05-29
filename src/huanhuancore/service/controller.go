package service

import (
	r "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"
	t "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/model"
)

func checktx(input *r.HuanHuanRequest) bool {
	switch input.From {
	case t.TokenType_BTC:
		return true
	case t.TokenType_ETH:
		return true
	default:
		return true
	}
}
