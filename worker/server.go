package worker

import (
	"context"
	"crawler/proto"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/short-d/app/fw/service"
	"google.golang.org/grpc"
)

var _ proto.WorkerServer = (*Server)(nil)

// Server struct to encapsulate the Worker struct and starts a grpc server for worker
type Server struct {
	worker Worker
}

// StartExtractingLinks calls the worker to extract links
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

// Start launches grpc server for worker node
func (s Server) Start(port int) error {
	gRPCService, err := service.
		NewGRPCBuilder("Worker").
		RegisterHandler(func(server *grpc.Server) {
			proto.RegisterWorkerServer(server, &s)
		}).
		Build()
	if err != nil {
		return err
	}

	fmt.Printf("Worker started at %d\n", port)
	err = s.worker.Init("localhost", port)
	if err != nil {
		return err
	}
	gRPCService.StartAndWait(port)
	return nil
}

// NewServer creates a worker server connecting to master server running the given port and IP address
func NewServer(masterIP string, masterPort int, workerSecret string) Server {
	masterClient := newMasterClient(masterIP, masterPort)
	return Server{
		worker: newWorker(masterClient, workerSecret),
	}
}
