package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/tabac/df/pb"
)

var _ pb.DataFusionExecutorServer = &DataFusionExecutorServerImpl{}

type DataFusionExecutorServerImpl struct {
	network string
	address string

	pb.UnimplementedDataFusionExecutorServer
}

func New(network string) *DataFusionExecutorServerImpl {
	server := &DataFusionExecutorServerImpl{
		network: "tcp",
		address: ":50051",
	}
	if network == "unix" {
		server.network = "unix"
		server.address = "/tmp/df.sock"
	}

	return server
}

func (e *DataFusionExecutorServerImpl) Run() error {
	log.Println("df-go: Running server.")

	listener, err := net.Listen(e.network, e.address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	pb.RegisterDataFusionExecutorServer(server, e)

	err = server.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

func (e *DataFusionExecutorServerImpl) CreateSession(context.Context, *pb.CreateSessionRequest) (*pb.CreateSessionResponse, error) {
	return nil, nil
}

func (e *DataFusionExecutorServerImpl) ExecuteQuery(request *pb.ExecuteQueryRequest, stream grpc.ServerStreamingServer[pb.ExecuteQueryResponse]) error {
	for i := 0; i < 5; i++ {
		message := pb.ExecuteQueryResponse{
			Id:        uint64(i),
			RequestId: request.Id,
		}

		err := stream.Send(&message)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
