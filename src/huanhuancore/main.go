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

	"gitlab.com/packtumi9722/etcd-agency/agency"
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

	// register grpc service to etcd
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

	// init agency
	agency.InitAgency(config.ETCDHosts())
	agency.Instance().SubscribeService([]string{config.DBServiceName()})

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
