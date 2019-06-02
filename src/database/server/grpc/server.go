package grpc

import (
	"context"
	"sync"

	// api
	pgrpc "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/pending"
	rgrpc "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/record"

	// model
	pmodel "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/pending"
	rmodel "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/record"
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/model/reply"

	"gitlab.com/packtumi9722/huanhuanhuei/src/database/service"

	"github.com/golang/protobuf/ptypes/empty"
)

var instance *Server
var once sync.Once

// Server contrustor
type Server struct {
	svc *service.Service
}

// Instance get grpc server instance
func Instance() *Server {
	once.Do(func() {
		instance = &Server{svc: &service.Service{}}
	})
	return instance
}

// Close close grpc server
func (s *Server) Close() {
	s.svc.Close()
}

// GetPendings get all pending items
func (s *Server) GetPendings(ctx context.Context, _ *empty.Empty) (*pgrpc.PendingItems, error) {
	list, err := s.svc.GetPendings()

	if err != nil {
		return nil, err
	}

	ret := new(pgrpc.PendingItems)
	for _, v := range list {
		ret.Items = append(ret.Items, pmodel.NewPending(v))
	}

	return ret, nil
}

// GetRecord get single record
func (s *Server) GetRecord(ctx context.Context, id *rgrpc.RecordID) (*rgrpc.RecordDatum, error) {
	r, err := s.svc.GetRecord(id.Id)

	if err != nil {
		return nil, err
	}
	return &rgrpc.RecordDatum{Record: rmodel.NewRecord(r)}, nil
}

// GetRecords get multiple records
func (s *Server) GetRecords(ctx context.Context, ids *rgrpc.RecordIDs) (*rgrpc.RecordData, error) {
	ret := new(rgrpc.RecordData)
	for _, v := range ids.Ids {
		r, _ := s.svc.GetRecord(v)
		ret.Records = append(ret.Records, rmodel.NewRecord(r))
	}
	return ret, nil
}

// UpdatePending update pending item
func (s *Server) UpdatePending(ctx context.Context, item *pgrpc.PendingItem) (*reply.Reply, error) {
	err := s.svc.UpdatePending(item.Item)
	if err != nil {
		return &reply.Reply{Success: false, Message: "", Error: err.Error()}, err
	}
	return &reply.Reply{Success: true, Message: "", Error: ""}, nil
}

// UpdateRecord update/insert a record
func (s *Server) UpdateRecord(ctx context.Context, datum *rgrpc.RecordDatum) (*reply.Reply, error) {
	err := s.svc.UpdateRecord(datum.Record)
	if err != nil {
		return &reply.Reply{Success: false, Message: "", Error: err.Error()}, err
	}
	return &reply.Reply{Success: true, Message: "", Error: ""}, nil
}

// DeletePending delete pending item
func (s *Server) DeletePending(ctx context.Context, id *pgrpc.ItemID) (*reply.Reply, error) {
	err := s.svc.DeletePending(id.Id)
	if err != nil {
		return &reply.Reply{Success: false, Message: "", Error: err.Error()}, err
	}
	return &reply.Reply{Success: true, Message: "", Error: ""}, nil
}
