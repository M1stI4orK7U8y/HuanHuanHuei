package grpc

import (
	"context"

	rdpro "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/record"
	rppro "gitlab.com/packtumi9722/huanhuanhuei/src/database/model/reply"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/client/grpc"
)

// UpdateRecord call grpc to update record
func UpdateRecord(in *rdpro.RecordDatum) (*rppro.Reply, error) {
	conn, err := grpc.ConnectDB()
	if err != nil {
		return nil, err
	}

	// close connection
	defer conn.Close()
	c := rdpro.NewDatabaseClient(conn)
	return c.UpdateRecord(context.Background(), in)
}
