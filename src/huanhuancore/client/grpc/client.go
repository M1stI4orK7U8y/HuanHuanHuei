package grpc

import (
	"errors"

	agency "gitlab.com/packtumi9722/etcd-agency"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/config"
	"google.golang.org/grpc"
)

// ConnectDB connect to db
func ConnectDB() (*grpc.ClientConn, error) {
	if agency.ServiceAvailable(config.DBServiceName()) {
		worker := agency.GetServiceWorker(config.DBServiceName())
		return grpc.Dial(worker.Address(), grpc.WithInsecure())
	}
	return nil, errors.New(config.DBServiceName() + " grpc service not available")
}
