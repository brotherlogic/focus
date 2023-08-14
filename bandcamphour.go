package main

import (
	"fmt"
	"sort"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/focus/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
)

func (s *Server) getBandcampHour(ctx context.Context, config *pb.Config) (*pb.Focus, error) {
	if time.Now().Hour() >= 16 {
		return nil, fmt.Errorf("not the time for bandcamp")
	}

	resp, err := s.ghClient.GetIssues(ctx, &pbgh.GetAllRequest{})
	if err != nil {
		return nil, err
	}

	sort.SliceStable(resp.Issues, func(i, j int) bool {
		return resp.GetIssues()[i].DateAdded < resp.GetIssues()[j].DateAdded
	})

	for _, issue := range resp.GetIssues() {
		if issue.GetService() == "bandcampserver" {
			return &pb.Focus{
				Type:   pb.Focus_FOCUS_ON_NON_HOME_TASKS,
				Detail: issue.GetTitle(),
				Link:   fmt.Sprintf("https://github.com/brotherlogic/%v/issues/%v", issue.GetService(), issue.GetNumber()),
			}, nil
		}
	}

	return nil, fmt.Errorf("no bandcamp issues are remaining")
}
