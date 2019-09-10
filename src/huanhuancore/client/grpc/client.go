package grpc

import (
	agency "gitlab.com/packtumi9722/etcd-agency"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/config"
	"google.golang.org/grpc"
)

// ConnectDB connect to db
func ConnectDB() (*grpc.ClientConn, error) {
	worker := agency.GetServiceWorker(config.DBServiceName())
	return grpc.Dial(worker.Address, grpc.WithInsecure())
}
