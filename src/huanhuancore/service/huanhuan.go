package service

import (
	"errors"

	//model
	rdquest "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/record"

	//api
	huanhuan "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"

	// service
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/service/btc"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/service/grpc"
)

// DoHuanHuan do huanhuan service
func (*Service) DoHuanHuan(input *huanhuan.HuanHuanRequest) error {
	if input == nil {
		return errors.New("DoHuanHuan input nil pointer")
	}

	// get the tx detail
	intx := getTxDetail(input.From, input.FromTxid)

	// do first check
	firsterr := firstcheck(input, intx)
	if firsterr != nil {
		return firsterr
	}

	// insert to database
	req := new(rdquest.RecordDatum)
	req.Record = createRecord(input.From, input.To, input.Receiver, intx)
	grpc.UpdateRecord(req)

	// send to receiver

	return nil
}

func firstcheck(input *huanhuan.HuanHuanRequest, intx interface{}) error {
	// check input tx
	if intx != nil || btc.CheckInputTx(input.From, intx) == false {
		return errors.New("btc input tx error")
	}
	// check receiver field
	if validReceiver(input.To, input.Receiver) == false {
		return errors.New("btc receiver data error")
	}
	return nil
}
