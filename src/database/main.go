package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	dbgrpc "gitlab.com/packtumi9722/huanhuanhuei/src/database/server/grpc"
	pendpro "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/pending"
	rdpro "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/record"
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
	rdpro.RegisterDatabaseServer(s, dbgrpc.Instance())
	pendpro.RegisterDatabaseServer(s, dbgrpc.Instance())
	// close grpc server
	defer dbgrpc.Instance().Close()

	// // register reflection service
	reflection.Register(s)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("%v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Printf("database grpc service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
