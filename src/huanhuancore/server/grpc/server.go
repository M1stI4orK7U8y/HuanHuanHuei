package grpc

import (
	"context"
	"sync"

	"gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/reply"
	huanhuan "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"
)

var instance *Server
var once sync.Once

// Server contrustor
type Server struct {
	//svc *service.Service
}

// Instance get grpc server instance
func Instance() *Server {
	once.Do(func() {
		instance = &Server{ /*svc: &service.Service{}*/ }
	})
	return instance
}

// Close close grpc server
func (s *Server) Close() {
	//s.svc.Close()
}

// DoHuanHuan requst exchange job
func (s *Server) DoHuanHuan(ctx context.Context, in *huanhuan.HuanHuanRequest) (*reply.Reply, error) {

	return nil, nil
}
