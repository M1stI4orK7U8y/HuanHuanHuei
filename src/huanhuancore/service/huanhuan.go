package service

import (
	"errors"
	"time"

	//model
	rdquest "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/record"
	rd "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/record"

	//api
	huanhuan "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"

	// service
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/service/grpc"
)

// DoHuanHuan do huanhuan service
func (*Service) DoHuanHuan(input *huanhuan.HuanHuanRequest) error {
	if input == nil {
		return errors.New("DoHuanHuan input nil pointer")
	}

	// get the tx detail
	intx, _ := getTxDetail(input.From, input.FromTxid)

	// do first check
	firsterr := checkInputTx(input, intx)
	if firsterr != nil {
		return firsterr
	}

	// insert to database
	req := new(rdquest.RecordDatum)
	req.Record = createRecord(input.From, input.To, input.Receiver, intx)
	grpc.UpdateRecord(req)

	// check pending
	// to do: not so important

	// send to receiver
	senthash, err := sendtoreceiver(input.To, input.Receiver, req.Record.ToToken.TokenValue)
	if err != nil {
		req.Record.StatusCode = rd.StatusCode_FAIL
		req.Record.StatusTime.FailedTime = time.Now().UTC().Unix()
	} else {
		req.Record.ToToken.Txhash = senthash
		req.Record.StatusCode = rd.StatusCode_FINISH
		req.Record.StatusTime.FinishedTime = time.Now().UTC().Unix()
	}
	// update database to final state
	grpc.UpdateRecord(req)

	return nil
}
