package main

import (
	"fmt"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	dstore_client "github.com/brotherlogic/dstore/client"
	github_client "github.com/brotherlogic/githubcard/client"

	pbg "github.com/brotherlogic/goserver/proto"
)

// Server main server type
type Server struct {
	*goserver.GoServer
	dclient  *dstore_client.DStoreClient
	ghclient *github_client.GHClient
}

// Init builds the server
func Init() *Server {
	s := &Server{
		GoServer: &goserver.GoServer{},
	}
	s.dclient = &dstore_client.DStoreClient{Gs: s.GoServer}
	s.ghclient = &github_client.GHClient{Gs: s.GoServer}
	return s
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {

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
	return []*pbg.State{
		&pbg.State{Key: "magic", Value: int64(12345)},
	}
}

func main() {
	server := Init()
	server.PrepServer("focus")
	server.Register = server

	err := server.RegisterServerV2(false)
	if err != nil {
		return
	}

	fmt.Printf("%v", server.Serve())
}
