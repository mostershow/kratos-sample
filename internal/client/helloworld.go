package client

import (
	v1 "github.com/go-kratos/kratos-layout/api/helloworld/v1"
	grpc2 "google.golang.org/grpc"
)

func NewGreeterClient(conn *grpc2.ClientConn) v1.GreeterClient {
	return v1.NewGreeterClient(conn)
}
