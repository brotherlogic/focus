package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/focus/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
)

func (s *Server) getHomeTaskFocus(ctx context.Context, config *pb.Config) (*pb.Focus, error) {
	resp, err := s.ghClient.GetIssues(ctx, &pbgh.GetAllRequest{})
	if err != nil {
		return nil, err
	}

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

	if (config.IssueCount["home"] >= 4) || ((time.Now().Weekday() != time.Saturday && time.Now().Weekday() != time.Sunday) && config.IssueCount["home"] >= 2) {
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
