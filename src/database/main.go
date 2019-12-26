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

	agency "gitlab.com/packtumi9722/etcd-agency"
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

	agency.InitAgency(config.ETCDHosts(), agency.V3)
	w := new(worker.Info)
	w.Name = config.Name()
	w.ServiceName = config.ServiceName()
	w.Address = lis.Addr().String()
	w.Protocol = "grpc"

	go func() {
		sayIAmAlive := agency.RegisterService(w)
		for {
			sayIAmAlive <- w
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
