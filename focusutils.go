package main

import (
	"time"

	pb "github.com/brotherlogic/focus/proto"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	pbds "github.com/brotherlogic/dstore/proto"
	pbgd "github.com/brotherlogic/godiscogs"
	pbrcl "github.com/brotherlogic/recordcleaner/proto"
	pbrc "github.com/brotherlogic/recordcollection/proto"
)

var (
	CONFIG = "github.com/brotherlogic/focus/config"
)

func getImage(images []*pbgd.Image) string {
	if len(images) > 0 {
		return images[0].Uri
	}

	return ""
}

func (s *Server) save(ctx context.Context, config *pb.Config) error {
	data, _ := proto.Marshal(config)
	_, err := s.dsClient.Write(ctx, &pbds.WriteRequest{Key: CONFIG, Value: &anypb.Any{Value: data}})
	return err
}

func (s *Server) load(ctx context.Context) (*pb.Config, error) {
	data, err := s.dsClient.Read(ctx, &pbds.ReadRequest{Key: CONFIG})
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			data = &pbds.ReadResponse{Value: &anypb.Any{}}
		} else {
			return nil, err
		}
	}

	config := &pb.Config{}
	proto.Unmarshal(data.GetValue().GetValue(), config)

	datestr := time.Now().Format("01/02/06")
	if config.GetDate() != datestr || config.GetIssueCount() == nil {
		config.Date = datestr
		config.IssueCount = make(map[string]int32)
		config.IssuesSeen = make(map[string]bool)
	}

	if config.IssuesSeen == nil {
		config.IssuesSeen = make(map[string]bool)
	}

	for key, val := range config.IssueCount {
		issueCount.With(prometheus.Labels{"service": key}).Set(float64(val))
	}

	return config, nil
}

func (s *Server) getRecordCleaningFocus(ctx context.Context, _ *pb.Config) (*pb.Focus, error) {
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
