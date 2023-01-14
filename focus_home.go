package main

import (
	"fmt"
	"sort"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/focus/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
)

func (s *Server) getHomeTaskFocus(ctx context.Context, config *pb.Config) (*pb.Focus, error) {

	if config.HomeCount > 2 {
		return nil, fmt.Errorf("done enough stuff in the home today")
	}

	resp, err := s.ghClient.GetIssues(ctx, &pbgh.GetAllRequest{})
	if err != nil {
		return nil, err
	}

	sort.SliceStable(resp.Issues, func(i, j int) bool {
		return resp.GetIssues()[i].DateAdded < resp.GetIssues()[j].DateAdded
	})

	for _, issue := range resp.GetIssues() {
		if issue.GetService() == "home" {
			return &pb.Focus{
				Type:   pb.Focus_FOCUS_ON_HOME_TASKS,
				Detail: issue.GetTitle(),
			}, nil
		}
	}

	return nil, fmt.Errorf("no home issues are remaining")
}
