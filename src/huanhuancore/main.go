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
	"gitlab.com/packtumi9722/etcd-agency/worker"
	huangrpc "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/api/grpc"
	huanserver "gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/server/grpc"

	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/config"
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
	huangrpc.RegisterHuanhuanServer(s, huanserver.Instance())

	// // register reflection service
	reflection.Register(s)

	// init agency
	agency.InitAgency(config.ETCDHosts(), agency.V3)
	// register grpc service to etcd
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

	// subscribe grpc service
	agency.SubscribeService([]string{config.DBServiceName()})

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
