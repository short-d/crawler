package worker

import (
	"context"
	"crawler/proto"
	"fmt"

	"google.golang.org/grpc"
)

// MasterClient to connect to master server
type MasterClient struct {
	masterIP   string
	masterPort int
	client     proto.MasterClient
}

// FinishExtractingLinks makes a grpc request to inform Finishing extracting links by this worker
func (m MasterClient) FinishExtractingLinks(workerID int, links []string) error {
	req := proto.FinishExtractingLinksRequest{
		WorkerId: int32(workerID),
		Links:    links,
	}
	_, err := m.client.FinishExtractingLinks(context.Background(), &req)
	return err
}

// RegisterWorker makes a grpc call to register the new worker to master
func (m MasterClient) RegisterWorker(workerIP string, workerPort int, secret string) (int, error) {
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

// Connect connects to the master grpc master server
func (m *MasterClient) Connect() error {
	address := fmt.Sprintf("%s:%d", m.masterIP, m.masterPort)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	m.client = proto.NewMasterClient(conn)
	return nil
}

func newMasterClient(ip string, port int) MasterClient {
	return MasterClient{
		masterIP:   ip,
		masterPort: port,
	}
}
