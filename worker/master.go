package worker

import (
	"context"
	"crawler/proto"
	"fmt"

	"google.golang.org/grpc"
)

type Master struct {
	masterIP   string
	masterPort int
	client     proto.MasterClient
}

func (m Master) FinishExtractingLinks(workerID int, links []string) error {
	req := proto.FinishExtractingLinksRequest{
		WorkerId: int32(workerID),
		Links:    links,
	}
	_, err := m.client.FinishExtractingLinks(context.Background(), &req)
	return err
}

func (m Master) RegisterWorker(workerIP string, workerPort int, secret string) (int, error) {
	req := proto.RegisterWorkerRequest{
		Ip:     workerIP,
		Port:   int32(workerPort),
		Secret: secret,
	}
	res, err := m.client.RegisterWorker(context.Background(), &req)
	if err != nil {
		return 0, err
	}
	return int(res.WorkerId), nil
}

func (m *Master) Connect() error {
	address := fmt.Sprintf("%s:%d", m.masterIP, m.masterPort)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	m.client = proto.NewMasterClient(conn)
	return nil
}

func newMaster(ip string, port int) Master {
	return Master{
		masterIP:   ip,
		masterPort: port,
	}
}
