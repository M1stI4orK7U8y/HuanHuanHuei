package grpc

import (
	"gitlab.com/packtumi9722/etcd-agency/agency"
	"gitlab.com/packtumi9722/huanhuanhuei/src/entry/config"
	"google.golang.org/grpc"
)

// ConnectDB connect to db
func ConnectDB() (*grpc.ClientConn, error) {
	worker := agency.Instance().GetServiceWorker(config.DBServiceName())
	return grpc.Dial(worker.WorkerInfo().Address, grpc.WithInsecure())
}

// ConnectCore connect to core
func ConnectCore() (*grpc.ClientConn, error) {
	worker := agency.Instance().GetServiceWorker(config.CoreServiceName())
	return grpc.Dial(worker.WorkerInfo().Address, grpc.WithInsecure())
}
