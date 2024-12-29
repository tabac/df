package client

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/tabac/df/pb"
)

type DataFusionExecutorClientImpl struct {
	conn   *grpc.ClientConn
	client pb.DataFusionExecutorClient
}

func New(network string) (*DataFusionExecutorClientImpl, error) {
	var target string = "127.0.0.1:50051"
	if network == "unix" {
		target = "unix:///tmp/df.sock"
	}

	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewDataFusionExecutorClient(conn)

	return &DataFusionExecutorClientImpl{
		conn:   conn,
		client: client,
	}, nil
}

func (c *DataFusionExecutorClientImpl) Run(id int) error {
	log.Println("df-go: Running client.")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := pb.ExecuteQueryRequest{Id: 123}

	stream, err := c.client.ExecuteQuery(ctx, &request)
	if err != nil {
		return err
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("df-go: Client got response: (%d, %d, %d)\n", id, response.Id, response.RequestId)
	}

	return nil
}
