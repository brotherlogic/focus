package main

import (
	"context"
	"testing"

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
	s.dsClient.ErrorCode[CONFIG] = codes.DataLoss

	config, err := s.load(context.Background())
	if err == nil {
		t.Errorf("Should have failed: %v", config)
	}
}
