package main

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dstore_client "github.com/brotherlogic/dstore/client"
	pb "github.com/brotherlogic/focus/proto"
	github_client "github.com/brotherlogic/githubcard/client"
	pbgd "github.com/brotherlogic/godiscogs/proto"
	pbrc "github.com/brotherlogic/recordcollection/proto"
	tasklist_client "github.com/brotherlogic/tasklist/client"

	recordcleaner_client "github.com/brotherlogic/recordcleaner/client"
	recordcollection_client "github.com/brotherlogic/recordcollection/client"
)

type testCheck struct{}

func (t *testCheck) home(ctx context.Context) error {
	return nil
}

func InitTestServer() *Server {
	s := Init(true)
	s.dialer = &testCheck{}
	s.test = true
	s.cleanerClient = &recordcleaner_client.RecordCleanerClient{Test: true}
	s.rccClient = &recordcollection_client.RecordCollectionClient{Test: true}
	s.ghClient = &github_client.GHClient{Test: true}
	s.dsClient = &dstore_client.DStoreClient{Test: true}
	s.tasklistClient = &tasklist_client.TasklistClient{Test: true}
	return s
}

func TestGetRecordCleanerFocusSucceeds(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getRecordCleaningFocus}
	s.rccClient.AddRecord(&pbrc.Record{Release: &pbgd.Release{InstanceId: 1234, Images: []*pbgd.Image{&pbgd.Image{Uri: "blah"}}}})

	res, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err != nil {
		t.Fatalf("Bad focus pull for cleaner: %v", err)
	}

	if res.GetFocus().GetType() != pb.Focus_FOCUS_ON_RECORD_CLEANING {
		t.Errorf("Bad focus: %v", res)
	}

	if len(res.Focus.GetLink()) == 0 {
		t.Errorf("No link: %v", res)
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

func TestGetRecordCleanerFocusWater(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getRecordCleaningFocus}
	s.cleanerClient.ErrorCode = codes.FailedPrecondition

	r, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if status.Code(err) != codes.OK {
		t.Errorf("Expected this to fail on call to record cleaner: %v -> %v", err, r)
	}
}

func TestGeRecordCleanerFocusFailsOnRecord(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getRecordCleaningFocus}
	s.rccClient.ErrorCode = codes.Unknown

	r, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if status.Code(err) != codes.OutOfRange {
		t.Errorf("Expected this to fail on call to record collection: %v", r)
	}
}
