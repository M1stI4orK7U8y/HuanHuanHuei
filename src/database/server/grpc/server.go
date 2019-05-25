package grpc

import (
	"context"

	"gitlab.com/packtumi9722/huanhuanhuei/src/database/types/pending"
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/types/record"
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/types/reply"
)

// Server contrustor
type Server struct{}

// GetPendings get all pending items
func (s *Server) GetPendings(ctx context.Context) (*pending.PendingItems, error) {
	return nil, nil
}

// UpdatePending update pending item
func (s *Server) UpdatePending(ctx context.Context, items *pending.PendingItem) (*reply.Reply, error) {
	return &reply.Reply{Success: true, Message: "", Error: ""}, nil
}

// DeletePending delete pending item
func (s *Server) DeletePending(ctx context.Context, id *pending.ItemID) (*reply.Reply, error) {
	return &reply.Reply{Success: true, Message: "", Error: ""}, nil
}

// GetRecord get single record
func (s *Server) GetRecord(ctx context.Context, id *record.RecordID) (*record.RecordDatum, error) {
	return nil, nil
}

// GetRecords get multiple records
func (s *Server) GetRecords(ctx context.Context, ids *record.RecordIDs) (*record.RecordData, error) {
	return nil, nil
}

// UpdateRecord update/insert a record
func (s *Server) UpdateRecord(ctx context.Context, datum *record.RecordDatum) (*reply.Reply, error) {
	return &reply.Reply{Success: true, Message: "", Error: ""}, nil
}
