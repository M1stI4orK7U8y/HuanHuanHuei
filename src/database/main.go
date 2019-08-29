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

	pendpro "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/pending"
	rdpro "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/record"
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/config"
	dbgrpc "gitlab.com/packtumi9722/huanhuanhuei/src/database/server/grpc"

	"gitlab.com/packtumi9722/etcd-agency/worker"
)

func main() {
	grpcproc()
}

func grpcproc() {
	lis, err := net.Listen("tcp", config.IP()+config.Port())
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

	w, err := worker.NewWorker(config.ETCDHosts(), config.ETCDTimeout())
	info := w.WorkerInfo()
	info.Name = config.Name()
	info.ServiceName = config.ServiceName()
	info.Address = lis.Addr().String()
	info.Protocol = "grpc"

	go func() {
		for {
			w.SayIAmAlive(info)
			time.Sleep(config.Heartbeat())
		}
	}()

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
