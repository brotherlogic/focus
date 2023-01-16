package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/focus/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
	"github.com/brotherlogic/godiscogs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestEmptyImages(t *testing.T) {
	res := getImage([]*godiscogs.Image{})
	if res != "" {
		t.Errorf("Should have been blank: %v", res)
	}
}

func TestBadLoad(t *testing.T) {
	s := InitTestServer()
	s.dsClient.ErrorCode = make(map[string]codes.Code)
	s.dsClient.ErrorCode[CONFIG] = codes.InvalidArgument

	config, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err != nil {
		t.Errorf("Should not have failed: %v, %v", config, err)
	}
}

func TestBadLoadActual(t *testing.T) {
	s := InitTestServer()
	s.dsClient.ErrorCode = make(map[string]codes.Code)
	s.dsClient.ErrorCode[CONFIG] = codes.DataLoss

	config, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err == nil {
		t.Errorf("Should not have failed: %v, %v", config, err)
	}
}

func TestUpdateSuccess(t *testing.T) {
	s := InitTestServer()

	_, err := s.ChangeUpdate(context.Background(), &pbgh.ChangeUpdateRequest{Issue: &pbgh.Issue{Service: "home"}})
	if err != nil {
		t.Errorf("This should succeed: %v", err)
	}
}

func TestUpdateFailOnLoad(t *testing.T) {
	s := InitTestServer()
	s.dsClient.ErrorCode = map[string]codes.Code{CONFIG: codes.DataLoss}

	_, err := s.ChangeUpdate(context.Background(), &pbgh.ChangeUpdateRequest{Issue: &pbgh.Issue{Service: "home"}})
	if status.Code(err) != codes.DataLoss {
		t.Errorf("Should have failed with data loss: %v", err)
	}
}
