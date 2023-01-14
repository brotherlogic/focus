package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/focus/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
)

func TestGetHomeTasksSucceed(t *testing.T) {
	s := InitTestServer()
	s.foci = []FocusBuilder{s.getHomeTaskFocus}
	s.ghClient.AddIssue(context.Background(), &pbgh.Issue{Title: "Test", Service: "home"})

	res, err := s.GetFocus(context.Background(), &pb.GetFocusRequest{})
	if err != nil {
		t.Fatalf("Bad focus pull for home tasks: %v", err)
	}

	if res.GetFocus().GetType() != pb.Focus_FOCUS_ON_HOME_TASKS {
		t.Errorf("Bad focus: %v", res)
	}
}
