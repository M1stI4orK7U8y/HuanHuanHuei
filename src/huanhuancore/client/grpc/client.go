package grpc

import (
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/config"
	"google.golang.org/grpc"
)

// ConnectDB connect to db
func ConnectDB() (*grpc.ClientConn, error) {
	return grpc.Dial(config.DBGrpcURL(), grpc.WithInsecure())
}
