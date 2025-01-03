package server

import (
	"context"
	"encoding/binary"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

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
		address: "127.0.0.1:50051",
	}
	if network == "unix" {
		server.network = "unix"
		server.address = "/tmp/df.sock"
	}

	http.HandleFunc("/execute", server.ExecuteQueryHttp)

	return server
}

func (e *DataFusionExecutorServerImpl) Run() error {
	log.Println("df-go: Running server.")

	if e.network == "http" {
		http.ListenAndServe(e.address, nil)
	} else {
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

func (e *DataFusionExecutorServerImpl) ExecuteQueryHttp(w http.ResponseWriter, req *http.Request) {
	for i := 0; i < 5; i++ {
		message := pb.ExecuteQueryResponse{
			Id:        uint64(i),
			RequestId: 12345,
		}

		bytes, err := proto.Marshal(&message)
		if err != nil {
			panic(err)
		}

		var n int = len(bytes)

		var buf [2]byte
		binary.LittleEndian.PutUint16(buf[:], uint16(n))

		c, err := w.Write(buf[:])
		if c != 2 || err != nil {
			panic(err)
		}

		for n > 0 {
			c, err := w.Write(bytes)
			if err != nil {
				panic(err)
			}

			n -= c
		}
	}
}
