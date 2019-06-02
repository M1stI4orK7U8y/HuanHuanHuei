package grpc

import (
	"gitlab.com/packtumi9722/huanhuanhuei/src/entry/config"
	"google.golang.org/grpc"
)

// ConnectDB connect to db
func ConnectDB() (*grpc.ClientConn, error) {
	return grpc.Dial(config.DBGrpcURL(), grpc.WithInsecure())
}

// ConnectCore connect to core
func ConnectCore() (*grpc.ClientConn, error) {
	return grpc.Dial(config.CoreGrpcURL(), grpc.WithInsecure())
}
