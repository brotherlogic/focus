package main

import (
	"sort"

	pbgh "github.com/brotherlogic/githubcard/proto"
	pb "github.com/brotherlogic/tasklist/proto"

	"golang.org/x/net/context"
)

func (s Server) processTaskLists(ctx context.Context, config *pb.Config) error {
	for _, list := range config.GetLists() {
		sort.SliceStable(list.Tasks, func(i, j int) bool {
			return list.Tasks[i].GetIndex() < list.Tasks[j].GetIndex()
		})

		for _, item := range list.GetTasks() {
			if item.GetState() == pb.Task_UNKNOWN || item.GetState() == pb.Task_TASK_WAITING {
				issue, err := s.ghclient.AddIssue(ctx, &pbgh.Issue{
					Title:   item.GetTitle(),
					Service: item.GetJob(),
					Body:    item.GetTitle(),
				})
				if err != nil {
					return err
				}
				item.State = pb.Task_TASK_IN_PROGRESS
				item.IssueNumber = issue.GetNumber()
			}
		}
	}

	return nil
}
