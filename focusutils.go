package main

import (
	pb "github.com/brotherlogic/focus/proto"
	"golang.org/x/net/context"

	pbrcl "github.com/brotherlogic/recordcleaner/proto"
)

func (s *Server) getRecordCleaningFocus(ctx context.Context) (*pb.Focus, error) {
	toclean, err := s.cleanerClient.GetClean(ctx, &pbrcl.GetCleanRequest{})
	if err != nil {
		return nil, err
	}

	record := s.rccclient.GetRecord(ctx, &pbrc.GetRecordRequest{InstanceId: toclean.GetInstanceId()})

	return &pb.Focus{Type: pb.Focus_FOCUS_ON_RECORD_CLEANING}
}
