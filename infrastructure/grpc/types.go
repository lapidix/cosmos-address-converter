package grpc

import (
  "google.golang.org/grpc"
)

type ClientInterface interface {
  Close() error
}
type GRPCClientInterface interface {
  ClientInterface
  GetConnection() *grpc.ClientConn
}

type GRPCClient struct {
  conn *grpc.ClientConn
}
