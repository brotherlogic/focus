package main

import (
	"context"
	"testing"
	"time"

	pbds "github.com/brotherlogic/dstore/proto"
	pb "github.com/brotherlogic/focus/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestGetHomeTasksSucceed(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getHomeTaskFocus}
	s.ghClient.AddIssue(context.Background(), &pbgh.Issue{Title: "Test", Service: "home", DateAdded: time.Now().Unix()})
	s.ghClient.AddIssue(context.Background(), &pbgh.Issue{Title: "Test2", Service: "home", DateAdded: time.Now().Add(time.Hour).Unix()})

	res, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err != nil {
		t.Fatalf("Bad focus pull for home tasks: %v", err)
	}

	if res.GetFocus().GetType() != pb.Focus_FOCUS_ON_HOME_TASKS {
		t.Errorf("Bad focus: %v", res)
	}

	if res.GetFocus().GetDetail() != "Test" {
		t.Errorf("Bad ordering on home tasks: %v", res)
	}
}

func TestGetHomeTasksP1Succeed(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getHomeTaskFocus}
	s.ghClient.AddIssue(context.Background(), &pbgh.Issue{Title: "Test", Service: "home", DateAdded: time.Now().Unix()})
	s.ghClient.AddIssue(context.Background(), &pbgh.Issue{Title: "P1: Test2", Service: "home", DateAdded: time.Now().Add(time.Hour).Unix()})

	res, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err != nil {
		t.Fatalf("Bad focus pull for home tasks: %v", err)
	}

	if res.GetFocus().GetType() != pb.Focus_FOCUS_ON_HOME_TASKS {
		t.Errorf("Bad focus: %v", res)
	}

	if res.GetFocus().GetDetail() == "Test" {
		t.Errorf("Bad ordering on home tasks")
	}
}

func TestGetHomeTasksFailOnIssuePull(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getHomeTaskFocus}
	s.ghClient.ErrorCode = codes.DataLoss

	_, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err == nil || status.Code(err) != codes.OutOfRange {
		t.Fatalf("Bad focus pull for home tasks - should be dataloss: %v", err)
	}
}

func TestGetHomeTasksFailOnNoIssues(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getHomeTaskFocus}

	_, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err == nil {
		t.Fatalf("Bad focus pull for home tasks (expected no issues): %v", err)
	}
}

func TestGetHomeTasksFailOnHomeTasksDone(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getHomeTaskFocus}
	config := &pb.Config{IssueCount: map[string]int32{"home": 5}, Date: time.Now().Format("01/02/06")}
	data, _ := proto.Marshal(config)
	s.dsClient.Write(context.Background(), &pbds.WriteRequest{Key: CONFIG, Value: &anypb.Any{Value: data}})
	s.ghClient.AddIssue(context.Background(), &pbgh.Issue{Title: "Test", Service: "home", DateAdded: time.Now().Unix()})
	s.ghClient.AddIssue(context.Background(), &pbgh.Issue{Title: "Test2", Service: "home", DateAdded: time.Now().Add(time.Hour).Unix()})

	_, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err == nil {
		t.Fatalf("Bad focus pull for home tasks (expected tasks dones): %v", err)
	}
}
