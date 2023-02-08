package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/focus/proto"
	pbgh "github.com/brotherlogic/githubcard/proto"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	focusCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "focus_current",
	}, []string{"type"})
	issueCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "focus_issues",
	}, []string{"service"})
)

func (s *Server) GetFocus(ctx context.Context, req *pb.GetFocusRequest) (*pb.GetFocusResponse, error) {
	config, err := s.load(ctx)
	if err != nil {
		return nil, err
	}

	for _, provider := range s.foci {
		focus, err := provider(ctx, config)
		if status.Code(err) == codes.OK {
			focusCount.With(prometheus.Labels{"type": fmt.Sprintf("%v", focus.GetType())}).Inc()
			return &pb.GetFocusResponse{Focus: focus}, nil
		}
	}

	return nil, status.Errorf(codes.OutOfRange, "No focus found")
}

func (s *Server) ChangeUpdate(ctx context.Context, req *pbgh.ChangeUpdateRequest) (*pbgh.ChangeUpdateResponse, error) {
	config, err := s.load(ctx)
	if err != nil {
		return nil, err
	}

	rep := fmt.Sprintf("%v-%v", req.GetIssue().GetService(), req.GetIssue().GetNumber())
	if config.IssuesSeen[rep] {
		return nil, status.Errorf(codes.AlreadyExists, "We've already seen %v", req.GetIssue())
	}

	config.IssuesSeen[rep] = true
	if !strings.Contains(req.GetIssue().GetTitle(), "P1") {
		config.IssueCount[req.GetIssue().GetService()]++
	}
	log.Printf("CLOSED WITH %v - %v", config.IssueCount, req)

	return nil, s.save(ctx, config)
}
