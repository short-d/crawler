package master

import (
	"context"
	"crawler/proto"
	"fmt"

	"github.com/short-d/app/fw/rpc"
)

// Worker encapsulates grpc api for worker server
type Worker struct {
	id     int
	ip     string
	port   int
	secret string
	client proto.WorkerClient
}

// Connect to worker server running on paticular ip and port
func (w *Worker) Connect() error {
	conn, err := rpc.
		NewClientConnBuilder(w.ip, w.port).
		Build()
	if err != nil {
		return err
	}

	w.client = proto.NewWorkerClient(conn)
	return nil
}

// FetchLinks starts fetching a link asynchronously 
func (w Worker) FetchLinks(ctx context.Context, source string) error {
	req := proto.StartExtractingLinksRequest{
		Secret: w.secret,
		Url:    source,
	}
	_, err := w.client.StartExtractingLinks(ctx, &req)
	return err
}

func newWorker(ip string, port int, secret string) Worker {
	return Worker{
		ip:     ip,
		port:   port,
		secret: secret,
	}
}
