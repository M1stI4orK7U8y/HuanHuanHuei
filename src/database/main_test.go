package main

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	dbgrpc "gitlab.com/packtumi9722/huanhuanhuei/src/database/server/grpc"
	pdpro "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/pending"
	rdpro "gitlab.com/packtumi9722/huanhuanhuei/src/database/api/grpc/record"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/test/bufconn"

	"github.com/golang/protobuf/ptypes/empty"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	rdpro.RegisterDatabaseServer(s, dbgrpc.Instance())
	pdpro.RegisterDatabaseServer(s, dbgrpc.Instance())

	reflection.Register(s)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(string, time.Duration) (net.Conn, error) {
	return lis.Dial()
}

func TestUpdateRecord(t *testing.T) {

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("Failed to dial bufnet: %s", err.Error())
	}
	defer conn.Close()
	c := rdpro.NewDatabaseClient(conn)
	datum := new(rdpro.RecordDatum)
	datum.Record = &rdpro.Record{Id: "123456"}
	_, err = c.UpdateRecord(context.Background(), datum)
	if err != nil {
		t.Errorf("UpdateRecord failed: %s", err.Error())
	} else {
		t.Logf("UpdateRecord success")
	}
	// Test for output here.
}

func TestGetRecord(t *testing.T) {

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("Failed to dial bufnet: %s", err.Error())
	}
	defer conn.Close()
	c := rdpro.NewDatabaseClient(conn)
	datum := new(rdpro.RecordID)
	datum.Id = "123456"
	_, err = c.GetRecord(context.Background(), datum)
	if err != nil {
		t.Errorf("GetRecord failed: %s", err.Error())
	} else {
		t.Logf("GetRecord success")
	}
	// Test for output here.
}

func TestGetRecords(t *testing.T) {

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("Failed to dial bufnet: %s", err.Error())
	}
	defer conn.Close()
	c := rdpro.NewDatabaseClient(conn)
	datum := new(rdpro.RecordIDs)
	datum.Ids = []string{"123456", "7891011"}
	_, err = c.GetRecords(context.Background(), datum)
	if err != nil {
		t.Errorf("GetRecords failed: %s", err.Error())
	} else {
		t.Logf("GetRecords success")
	}
	// Test for output here.
}

func TestUpdatePending(t *testing.T) {

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("Failed to dial bufnet: %s", err.Error())
	}
	defer conn.Close()
	c := pdpro.NewDatabaseClient(conn)
	datum := new(pdpro.PendingItem)
	datum.Item = &pdpro.Pending{Id: "123456", PendingTime: time.Now().UTC().Unix()}
	_, err = c.UpdatePending(context.Background(), datum)
	if err != nil {
		t.Errorf("UpdatePending failed: %s", err.Error())
	} else {
		t.Logf("UpdatePending success")
	}
	// Test for output here.
}

func TestGetPendings(t *testing.T) {

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("Failed to dial bufnet: %s", err.Error())
	}
	defer conn.Close()
	c := pdpro.NewDatabaseClient(conn)
	_, err = c.GetPendings(context.Background(), &empty.Empty{})
	if err != nil {
		t.Errorf("GetRecord failed: %s", err.Error())
	} else {
		t.Logf("GetRecord success")
	}
	// Test for output here.
}

func TestDeletePending(t *testing.T) {

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("Failed to dial bufnet: %s", err.Error())
	}
	defer conn.Close()
	c := pdpro.NewDatabaseClient(conn)
	datum := new(pdpro.ItemID)
	datum.Id = "123456"
	_, err = c.DeletePending(context.Background(), datum)
	if err != nil {
		t.Errorf("DeletePending failed: %s", err.Error())
	} else {
		t.Logf("DeletePending success")
	}
	// Test for output here.
}
