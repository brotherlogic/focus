package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/focus/proto"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	focusCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "focus_current",
	}, []string{"type"})
)

func (s *Server) GetFocus(ctx context.Context, req *pb.GetFocusRequest) (*pb.GetFocusResponse, error) {
	for _, provider := range s.foci {
		focus, err := provider(ctx)
		if status.Code(err) == codes.OK {
			focusCount.With(prometheus.Labels{"type": fmt.Sprintf("%v", focus.GetType())}).Inc()
			return &pb.GetFocusResponse{Focus: focus}, nil
		}
	}

	log.Printf("Reached here")

	return nil, status.Errorf(codes.OutOfRange, "No focus found")
}
