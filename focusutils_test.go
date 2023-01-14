package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/focus/proto"
	"github.com/brotherlogic/godiscogs"
	"google.golang.org/grpc/codes"
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
