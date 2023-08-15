package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/focus/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
)

type checkHome interface {
	home(ctx context.Context) error
}

type prodCheck struct {
	dialer func(context.Context, string, string) (*grpc.ClientConn, error)
}

func (p *prodCheck) home(ctx context.Context) error {
	conn, err := p.dialer(ctx, "gobuildslave", "cd")
	if err != nil {
		return status.Errorf(codes.Unavailable, "Cannot reach home server, likely not home %v", err)
	}
	conn.Close()
	return nil
}

func (s *Server) getHomeTaskFocus(ctx context.Context, config *pb.Config) (*pb.Focus, error) {
	err := s.dialer.home(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := s.ghClient.GetIssues(ctx, &pbgh.GetAllRequest{})
	if err != nil {
		return nil, err
	}

	s.trimIssues(ctx, resp)

	sort.SliceStable(resp.Issues, func(i, j int) bool {
		return resp.GetIssues()[i].DateAdded < resp.GetIssues()[j].DateAdded
	})

	// Look for any P1s and quick return these
	for _, issue := range resp.GetIssues() {
		if strings.Contains(issue.GetTitle(), "P1") {
			return &pb.Focus{
				Type:   pb.Focus_FOCUS_ON_HOME_TASKS,
				Detail: issue.GetTitle(),
				Link:   fmt.Sprintf("https://github.com/brotherlogic/home/issues/%v", issue.GetNumber()),
			}, nil
		}
	}

	if (config.IssueCount["home"] >= 10) || ((time.Now().Weekday() != time.Saturday && time.Now().Weekday() != time.Sunday) && config.IssueCount["home"] >= 10) {
		return nil, fmt.Errorf("done enough stuff in the home today")
	}

	for _, issue := range resp.GetIssues() {
		if issue.GetService() == "home" {
			return &pb.Focus{
				Type:   pb.Focus_FOCUS_ON_HOME_TASKS,
				Detail: issue.GetTitle(),
				Link:   fmt.Sprintf("https://github.com/brotherlogic/home/issues/%v", issue.GetNumber()),
			}, nil
		}
	}

	return nil, fmt.Errorf("no home issues are remaining")
}
