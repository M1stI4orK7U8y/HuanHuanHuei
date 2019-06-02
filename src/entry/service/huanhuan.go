package service

import (
	"context"

	"gitlab.com/packtumi9722/huanhuanhuei/src/database/model/reply"
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/model/token"
	"gitlab.com/packtumi9722/huanhuanhuei/src/entry/client/grpc"
	huanhuan "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"
)

// NewHuanHuanJob create new huanhuei (exchange) job
func NewHuanHuanJob(txid, receiver string, from, to token.TokenType) (*reply.Reply, error) {
	conn, err := grpc.ConnectCore()

	if err != nil {
		return nil, err
	}

	// close connection
	defer conn.Close()

	c := huanhuan.NewHuanhuanClient(conn)
	req := new(huanhuan.HuanHuanRequest)
	req.FromTxid = txid
	req.From = from
	req.Receiver = receiver
	req.To = to
	return c.DoHuanHuan(context.Background(), req)
}
