package master

import (
	"context"
	"crawler/proto"
	"fmt"
	"net"

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
	worker := newWorker(request.Ip, int(request.Port), request.Secret)
	err := worker.Connect()
	if err != nil {
		return nil, err
	}
	workerID := s.master.RegisterWorker(worker)
	fmt.Printf("Worker registed: ID(%d) IP(%s) PORT(%d) SECRET(%s)\n", workerID, request.Ip, int(request.Port), request.Secret)
	return &proto.RegisterWorkerResponse{
		WorkerId: int32(workerID),
	}, nil
}

func (s Server) FinishExtractingLinks(ctx context.Context, request *proto.FinishExtractingLinksRequest) (*empty.Empty, error) {
	s.master.FinishExtractingLinks(int(request.WorkerId), request.Links)
	return &empty.Empty{}, nil
}

func (s Server) Start(port int) error {
	address := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	sv := grpc.NewServer()
	proto.RegisterMasterServer(sv, &s)
	fmt.Printf("Master started at %d\n", port)
	return sv.Serve(lis)
}

func NewServer() Server {
	return Server{master: newMaster()}
}

func main() {
    gRPCService, err := service.
            NewGRPCBuilder("Master").
            RegisterHandler(func(server *grpc.NewServer) {
                    sv := grpc.Server{}
                    proto.RegisterMasterServer(server, sv)
            }).
            Build()
    if err != nil {
        panic(err)
    }
    gRPCService.StartAndWait(8080)
}
