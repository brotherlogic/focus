package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/focus/proto"
	recordcleaner_client "github.com/brotherlogic/recordcleaner/client"
	recordcollection_client "github.com/brotherlogic/recordcollection/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InitTestServer() *Server {
	s := Init()
	s.cleanerClient = &recordcleaner_client.RecordCleanerClient{Test: true}
	s.rccClient = &recordcollection_client.RecordCollectionClient{Test: true}
	return s
}

func TestGetRecordCleanerFocusSucceeds(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getRecordCleaningFocus}

	res, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err != nil {
		t.Fatalf("Bad focus pull for cleaner: %v", err)
	}

	if res.GetFocus().GetType() != pb.Focus_FOCUS_ON_RECORD_CLEANING {
		t.Errorf("Bad focus: %v", res)
	}
}

func TestGeRecordCleanerFocusFailsOnClean(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getRecordCleaningFocus}
	s.cleanerClient.ErrorCode = codes.Unknown

	r, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if status.Code(err) != codes.OutOfRange {
		t.Errorf("Expected this to fail on call to record cleaner: %v", r)
	}
}
