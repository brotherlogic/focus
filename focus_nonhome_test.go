package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/brotherlogic/focus/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetNonHomeTasksSucceed(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getNoHomeTaskFocus}
	s.ghClient.AddIssue(context.Background(), &pbgh.Issue{Title: "Test", Service: "home", DateAdded: time.Now().Unix()})
	s.ghClient.AddIssue(context.Background(), &pbgh.Issue{Title: "Test2", Service: "blah", DateAdded: time.Now().Add(time.Hour).Unix()})

	res, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err != nil {
		t.Fatalf("Bad focus pull for non home tasks: %v", err)
	}

	if res.GetFocus().GetType() != pb.Focus_FOCUS_ON_HOME_TASKS {
		t.Errorf("Bad focus: %v", res)
	}

	if res.GetFocus().GetDetail() != "Test2" {
		t.Errorf("Bad ordering on non home tasks")
	}
}

func TestGetNonHomeTasksFailOnIssuePull(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getNoHomeTaskFocus}
	s.ghClient.ErrorCode = codes.DataLoss

	_, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err == nil || status.Code(err) != codes.OutOfRange {
		t.Fatalf("Bad focus pull for home tasks - should be dataloss: %v", err)
	}
}

func TestGetNonHomeTasksFailOnNoIssues(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getNoHomeTaskFocus}

	_, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err == nil {
		t.Fatalf("Bad focus pull for home tasks (expected no issues): %v", err)
	}
}
