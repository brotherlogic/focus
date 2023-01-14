package main

import (
	pb "github.com/brotherlogic/focus/proto"
	"golang.org/x/net/context"

	pbrcl "github.com/brotherlogic/recordcleaner/proto"
	pbrc "github.com/brotherlogic/recordcollection/proto"
)

func (s *Server) getRecordCleaningFocus(ctx context.Context) (*pb.Focus, error) {
	toclean, err := s.cleanerClient.GetClean(ctx, &pbrcl.GetCleanRequest{})
	if err != nil {
		return nil, err
	}

	record, err := s.rccClient.GetRecord(ctx, &pbrc.GetRecordRequest{InstanceId: toclean.GetInstanceId()})
	if err != nil {
		return nil, err
	}

	return &pb.Focus{
		Type:   pb.Focus_FOCUS_ON_RECORD_CLEANING,
		Detail: record.GetRecord().GetRelease().GetTitle(),
		Link:   record.GetRecord().GetRelease().GetImages()[0].GetUri(),
	}, nil
}
