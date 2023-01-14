package main

import (
	pb "github.com/brotherlogic/focus/proto"
	"golang.org/x/net/context"

	pbgd "github.com/brotherlogic/godiscogs"
	pbrcl "github.com/brotherlogic/recordcleaner/proto"
	pbrc "github.com/brotherlogic/recordcollection/proto"
)

func getImage(images []*pbgd.Image) string {
	if len(images) > 0 {
		return images[0].Uri
	}

	return ""
}

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
		Link:   getImage(record.GetRecord().GetRelease().GetImages()),
	}, nil
}
