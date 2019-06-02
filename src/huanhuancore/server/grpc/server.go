package grpc

import (
	"context"
	"sync"

	"gitlab.com/packtumi9722/huanhuanhuei/src/database/model/reply"
	huanhuan "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/service"
)

var instance *Server
var once sync.Once

// Server contrustor
type Server struct {
	svc *service.Service
}

// Instance get grpc server instance
func Instance() *Server {
	once.Do(func() {
		instance = &Server{svc: &service.Service{}}
	})
	return instance
}

// Close close grpc server
func (s *Server) Close() {
	//s.svc.Close()
}

// DoHuanHuan requst exchange job
func (s *Server) DoHuanHuan(ctx context.Context, in *huanhuan.HuanHuanRequest) (*reply.Reply, error) {
	err := s.svc.DoHuanHuan(in)
	if err != nil {
		return &reply.Reply{Success: false, Message: "", Error: err.Error()}, err
	}
	return &reply.Reply{Success: true, Message: "", Error: ""}, nil
}
