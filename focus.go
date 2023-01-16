package main

import (
	"fmt"
	"time"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/focus/proto"
	pbg "github.com/brotherlogic/goserver/proto"

	dstore_client "github.com/brotherlogic/dstore/client"
	github_client "github.com/brotherlogic/githubcard/client"
	recordcleaner_client "github.com/brotherlogic/recordcleaner/client"
	recordcollection_client "github.com/brotherlogic/recordcollection/client"
)

type FocusBuilder = func(context.Context, *pb.Config) (*pb.Focus, error)

// Server main server type
type Server struct {
	*goserver.GoServer
	foci          []FocusBuilder
	cleanerClient *recordcleaner_client.RecordCleanerClient
	rccClient     *recordcollection_client.RecordCollectionClient
	ghClient      *github_client.GHClient
	dsClient      *dstore_client.DStoreClient
	test          bool
}

// Init builds the server
func Init(test bool) *Server {
	s := &Server{
		GoServer: &goserver.GoServer{},
	}
	s.cleanerClient = &recordcleaner_client.RecordCleanerClient{Gs: s.GoServer}
	s.rccClient = &recordcollection_client.RecordCollectionClient{Gs: s.GoServer}
	s.ghClient = &github_client.GHClient{Gs: s.GoServer}
	s.dsClient = &dstore_client.DStoreClient{Gs: s.GoServer}

	if test || time.Now().After(time.Date(2023, time.January, 23, 17, 30, 0, 0, time.Now().Location())) {
		s.foci = []FocusBuilder{s.getRecordCleaningFocus, s.getHomeTaskFocus}
	} else {
		s.foci = []FocusBuilder{s.getHomeTaskFocus}
	}
	return s
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterFocusServiceServer(server, s)
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{}
}

func main() {
	server := Init(false)
	server.PrepServer("focus")
	server.Register = server

	err := server.RegisterServerV2(false)
	if err != nil {
		return
	}

	fmt.Printf("%v", server.Serve())
}
