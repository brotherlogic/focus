package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	dspb "github.com/brotherlogic/dstore/proto"
	pb "github.com/brotherlogic/tasklist/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	CONFIG_KEY = "github.com/brotherlogic/tasklist/config"

	lists = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "tasklist_lists",
	})
)

func (s *Server) metrics(config *pb.Config) {
	lists.Set(float64(len(config.GetLists())))
}

func (s *Server) readConfig(ctx context.Context) (*pb.Config, error) {
	data, err := s.dclient.Read(ctx, &dspb.ReadRequest{Key: CONFIG_KEY})
	if err != nil {
		if status.Code(err) != codes.InvalidArgument {
			return nil, err
		}
		data = &dspb.ReadResponse{}
	}

	config := &pb.Config{}
	err = proto.Unmarshal(data.GetValue().GetValue(), config)

	if err == nil {
		s.metrics(config)
	}

	return config, err
}

func (s *Server) saveConfig(ctx context.Context, config *pb.Config) error {
	data, _ := proto.Marshal(config)
	_, err := s.dclient.Write(ctx, &dspb.WriteRequest{Key: CONFIG_KEY, Value: &anypb.Any{Value: data}})
	return err
}

func (s *Server) GetTaskLists(ctx context.Context, req *pb.GetTaskListsRequest) (*pb.GetTaskListsResponse, error) {
	config, err := s.readConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetTaskListsResponse{Lists: config.GetLists()}, nil
}

func (s *Server) AddTaskList(ctx context.Context, req *pb.AddTaskListRequest) (*pb.AddTaskListResponse, error) {
	config, err := s.readConfig(ctx)
	if err != nil {
		return nil, err
	}

	found := false
	for _, list := range config.GetLists() {
		if list.GetName() == req.GetAdd().GetName() {
			found = true
			break
		}
	}

	if found {
		return nil, status.Errorf(codes.AlreadyExists, "%v already exists", req.GetAdd().GetName())
	}

	config.Lists = append(config.Lists, req.GetAdd())

	err = s.processTaskLists(ctx, config)
	if err != nil {
		return nil, err
	}

	return &pb.AddTaskListResponse{}, s.saveConfig(ctx, config)
}
