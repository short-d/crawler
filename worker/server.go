package worker

import (
	"context"
	"crawler/proto"
	"fmt"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var _ proto.WorkerServer = (*Server)(nil)

type Server struct {
	worker Worker
}

func (s Server) StartExtractingLinks(ctx context.Context, request *proto.StartExtractingLinksRequest) (*empty.Empty, error) {
	fmt.Printf("Start extracting links from %s\n", request.Url)
	go func() {
		err := s.worker.ExtractLinks(request.Secret, request.Url)
		if err != nil {
			fmt.Println(err)
		}
	}()
	return &empty.Empty{}, nil
}

func (s Server) Start(port int) error {
	address := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	sv := grpc.NewServer()
	proto.RegisterWorkerServer(sv, &s)
	fmt.Printf("Worker started at %d\n", port)

	err = s.worker.Init("localhost", port)
	if err != nil {
		return err
	}
	return sv.Serve(lis)
}

func NewServer(masterIP string, masterPort int, workerSecret string) Server {
	master := newMaster(masterIP, masterPort)
	return Server{
		worker: newWorker(master, workerSecret),
	}
}
