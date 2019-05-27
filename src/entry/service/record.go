package service

import (
	"context"

	"gitlab.com/packtumi9722/huanhuanhuei/src/entry/client/grpc"

	rdpro "gitlab.com/packtumi9722/huanhuanhuei/src/database/types/record"
)

// GetRecord get record by id
func GetRecord(id string) (*rdpro.RecordDatum, error) {
	conn, err := grpc.ConnectDB()

	if err != nil {
		return nil, err
	}

	// close connection
	defer conn.Close()

	c := rdpro.NewDatabaseClient(conn)
	datum := new(rdpro.RecordID)
	datum.Id = id
	return c.GetRecord(context.Background(), datum)

}

// GetRecords get records by ids
func GetRecords(ids []string) (*rdpro.RecordData, error) {
	conn, err := grpc.ConnectDB()
	if err != nil {
		return nil, err
	}

	// close connection
	defer conn.Close()

	c := rdpro.NewDatabaseClient(conn)
	datum := new(rdpro.RecordIDs)
	datum.Ids = ids
	return c.GetRecords(context.Background(), datum)
}

// UpdateRecord update record
// func UpdateRecord(in *rdpro.RecordDatum) (*rppro.Reply, error) {
// 	conn, err := grpc.ConnectDB()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// close connection
// 	defer conn.Close()
// 	c := rdpro.NewDatabaseClient(conn)
// 	return c.UpdateRecord(context.Background(), in)
// }
