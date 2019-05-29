package service

import (
	"errors"

	huanhuan "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/service/validate"
)

// DoHuanHuan do huanhuan service
func (*Service) DoHuanHuan(input *huanhuan.HuanHuanRequest) error {
	if input == nil {
		return errors.New("DoHuanHuan input nil pointer")
	}

	// check input tx
	if validate.CheckInputTx(input) == false {
		return errors.New("validate input tx error")
	}
	// check receiver field
	if validReceiver(input.To, input.Receiver) == false {
		return errors.New("validate receiver data error")
	}

	// check all ok -> update to database
	return nil
}
