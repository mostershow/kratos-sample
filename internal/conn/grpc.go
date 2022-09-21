package conn

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	grpc2 "google.golang.org/grpc"
)

func NewGrpcClient() *grpc2.ClientConn {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
	)

	if err != nil {
		panic(err)
	}
	return conn
}
