package grpc

import (
	"context"
	"sync"

	"gitlab.com/packtumi9722/huanhuanhuei/src/database/service"
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/types/pending"
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/types/record"
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/types/reply"

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
func (s *Server) GetPendings(ctx context.Context, _ *empty.Empty) (*pending.PendingItems, error) {
	list, err := s.svc.GetPendings()

	if err != nil {
		return nil, err
	}

	ret := new(pending.PendingItems)
	for _, v := range list {
		ret.Items = append(ret.Items, pending.NewPending(v))
	}

	return ret, nil
}

// GetRecord get single record
func (s *Server) GetRecord(ctx context.Context, id *record.RecordID) (*record.RecordDatum, error) {
	r, err := s.svc.GetRecord(id.Txhash)

	if err != nil {
		return nil, err
	}
	return &record.RecordDatum{Record: record.NewRecord(r)}, nil
}

// GetRecords get multiple records
func (s *Server) GetRecords(ctx context.Context, ids *record.RecordIDs) (*record.RecordData, error) {
	ret := new(record.RecordData)
	for _, v := range ids.Txhashes {
		r, _ := s.svc.GetRecord(v)
		ret.Records = append(ret.Records, record.NewRecord(r))
	}
	return ret, nil
}

// UpdatePending update pending item
func (s *Server) UpdatePending(ctx context.Context, item *pending.PendingItem) (*reply.Reply, error) {
	err := s.svc.UpdatePending(item.Item)
	if err != nil {
		return &reply.Reply{Success: false, Message: "", Error: err.Error()}, nil
	}
	return &reply.Reply{Success: true, Message: "", Error: ""}, nil
}

// UpdateRecord update/insert a record
func (s *Server) UpdateRecord(ctx context.Context, datum *record.RecordDatum) (*reply.Reply, error) {
	err := s.svc.UpdateRecord(datum.Record)
	if err != nil {
		return &reply.Reply{Success: false, Message: "", Error: err.Error()}, nil
	}
	return &reply.Reply{Success: true, Message: "", Error: ""}, nil
}

// DeletePending delete pending item
func (s *Server) DeletePending(ctx context.Context, id *pending.ItemID) (*reply.Reply, error) {
	err := s.svc.DeletePending(id.Id)
	if err != nil {
		return &reply.Reply{Success: false, Message: "", Error: err.Error()}, nil
	}
	return &reply.Reply{Success: true, Message: "", Error: ""}, nil
}
