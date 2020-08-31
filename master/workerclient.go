package master

import (
	"context"
	"crawler/proto"
	"fmt"

	"google.golang.org/grpc"
)

type WorkerClient struct {
	id     int
	ip     string
	port   int
	secret string
	client proto.WorkerClient
}

func (w *WorkerClient) Connect() error {
	address := fmt.Sprintf("%s:%d", w.ip, w.port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	w.client = proto.NewWorkerClient(conn)
	return nil
}

func (w WorkerClient) FetchLinks(ctx context.Context, source string) error {
	req := proto.StartExtractingLinksRequest{
		Secret: w.secret,
		Url:    source,
	}
	_, err := w.client.StartExtractingLinks(ctx, &req)
	return err
}

func newWorkerClient(ip string, port int, secret string) WorkerClient {
	return WorkerClient{
		ip:     ip,
		port:   port,
		secret: secret,
	}
}
