package master

import (
	"context"
	"crawler/proto"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/short-d/app/fw/service"
	"google.golang.org/grpc"
)

var _ proto.MasterServer = (*Server)(nil)

type Server struct {
	master *Master
}

func (s Server) ExploreWebsite(ctx context.Context, request *proto.ExploreWebsiteRequest) (*empty.Empty, error) {
	fmt.Printf("Start exploring %s\n", request.Url)
	s.master.ExploreWebsite(ctx, request.Url)
	return &empty.Empty{}, nil
}

func (s Server) RegisterWorker(ctx context.Context, request *proto.RegisterWorkerRequest) (*proto.RegisterWorkerResponse, error) {
	workerClient := newWorkerClient(request.Ip, int(request.Port), request.Secret)
	err := workerClient.Connect()
	if err != nil {
		return nil, err
	}
	workerID := s.master.RegisterWorker(workerClient)
	fmt.Printf("WorkerClient registed: ID(%d) IP(%s) PORT(%d) SECRET(%s)\n", workerID, request.Ip, int(request.Port), request.Secret)
	return &proto.RegisterWorkerResponse{
		WorkerId: int32(workerID),
	}, nil
}

func (s Server) FinishExtractingLinks(ctx context.Context, request *proto.FinishExtractingLinksRequest) (*empty.Empty, error) {
	s.master.FinishExtractingLinks(int(request.WorkerId), request.Links)
	return &empty.Empty{}, nil
}

func (s Server) Start(port int) error {
	gRPCService, err := service.
		NewGRPCBuilder("Master").
		RegisterHandler(func(server *grpc.Server) {
			proto.RegisterMasterServer(server, &s)
		}).
		Build()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Master started at %d\n", port)
	gRPCService.StartAndWait(port)
	return nil
}

func NewServer() Server {
	return Server{master: newMaster()}
}
