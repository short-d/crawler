package master

import (
	"context"
	"crawler/proto"
	"fmt"

	"google.golang.org/grpc"
)

type Worker struct {
	id     int
	ip     string
	port   int
	secret string
	client proto.WorkerClient
}

func (w *Worker) Connect() error {
	address := fmt.Sprintf("%s:%d", w.ip, w.port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	w.client = proto.NewWorkerClient(conn)
	return nil
}

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
