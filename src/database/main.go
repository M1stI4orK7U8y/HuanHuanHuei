package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	dbgrpc "gitlab.com/packtumi9722/huanhuanhuei/src/database/server/grpc"
	pendpro "gitlab.com/packtumi9722/huanhuanhuei/src/database/types/pending"
	rdpro "gitlab.com/packtumi9722/huanhuanhuei/src/database/types/record"
)

func main() {
	grpcproc()
}

func grpcproc() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("%v", err)
	}

	// register database server
	s := grpc.NewServer()
	rdpro.RegisterDatabaseServer(s, dbgrpc.NewServer())
	pendpro.RegisterDatabaseServer(s, dbgrpc.NewServer())

	// // register reflection service
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("%v", err)
	}
}
