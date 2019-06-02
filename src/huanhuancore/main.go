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

	huangrpc "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"
	huanserver "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/server/grpc"
)

func main() {
	grpcproc()
}

func grpcproc() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("%v", err)
	}

	// register database server
	s := grpc.NewServer()
	huangrpc.RegisterHuanhuanServer(s, huanserver.Instance())
	//pendpro.RegisterDatabaseServer(s, dbgrpc.Instance())
	// close grpc server

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
			log.Printf("huanhuan grpc service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
