package main

import (
	"fmt"
	"sort"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/focus/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"
	pbtl "github.com/brotherlogic/tasklist/proto"
)

func (s *Server) getDeepFocus(ctx context.Context, config *pb.Config) (*pb.Focus, error) {
	if time.Now().Hour() < 20 && time.Now().Hour() != 16 && time.Now().Hour() != 17 {
		return nil, fmt.Errorf("not the time for deep focus")
	}

	resp, err := s.ghClient.GetIssues(ctx, &pbgh.GetAllRequest{})
	if err != nil {
		return nil, err
	}

	sort.SliceStable(resp.Issues, func(i, j int) bool {
		return resp.GetIssues()[i].DateAdded < resp.GetIssues()[j].DateAdded
	})

	for _, issue := range resp.Issues {
		if issue.GetService() == "recordsorganiser" {
			return &pb.Focus{
				Type:   pb.Focus_FOCUS_ON_NON_HOME_TASKS,
				Detail: issue.GetTitle(),
				Link:   fmt.Sprintf("https://github.com/brotherlogic/%v/issues/%v", issue.GetService(), issue.GetNumber()),
			}, nil
		}
	}

	tasks, err := s.tasklistClient.GetTasks(ctx, &pbtl.GetTasksRequest{Tags: []string{"gramophile"}})
	if err != nil {
		return nil, err
	}

	for _, issue := range resp.GetIssues() {
		if issue.GetService() != "home" && issue.GetService() != "recordalerting" {
			for _, task := range tasks.GetTasks() {
				if task.GetJob() == issue.GetService() && task.GetIssueNumber() == issue.GetNumber() {
					return &pb.Focus{
						Type:   pb.Focus_FOCUS_ON_NON_HOME_TASKS,
						Detail: issue.GetTitle(),
						Link:   fmt.Sprintf("https://github.com/brotherlogic/%v/issues/%v", issue.GetService(), issue.GetNumber()),
					}, nil
				}
			}

		}
	}

	return nil, fmt.Errorf("no non-home issues are remaining")
}
