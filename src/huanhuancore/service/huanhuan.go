package service

import (
	"errors"

	//model
	rdquest "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/record"

	//api
	huanhuan "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"

	// service
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/service/grpc"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/service/validate"
)

// DoHuanHuan do huanhuan service
func (*Service) DoHuanHuan(input *huanhuan.HuanHuanRequest) error {
	if input == nil {
		return errors.New("DoHuanHuan input nil pointer")
	}

	intx := getTxDetail(input.From, input.FromTxid)

	// check input tx
	if intx != nil || validate.CheckInputTx(input.From, intx) == false {
		return errors.New("validate input tx error")
	}
	// check receiver field
	if validReceiver(input.To, input.Receiver) == false {
		return errors.New("validate receiver data error")
	}

	// insert to database
	req := new(rdquest.RecordDatum)
	req.Record = createRecord(input.From, input.To, input.Receiver, intx)
	grpc.UpdateRecord(req)

	return nil
}
