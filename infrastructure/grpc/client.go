package grpc

import (
  "context"
  "fmt"
  "time"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
)

type Config struct {
  grpcUrl     string
  DialTimeout time.Duration
}

func NewDefaultConfig(grpcUrl string) *Config {
  return &Config{
    grpcUrl:     grpcUrl,
    DialTimeout: 10 * time.Second,
  }
}

func NewClient(cfg *Config) (*grpc.ClientConn, error) {
  if cfg == nil {
    return nil, fmt.Errorf("config cannot be nil")
  }

  ctx, cancel := context.WithTimeout(context.Background(), cfg.DialTimeout)
  defer cancel()

  conn, err := grpc.DialContext(
    ctx,
    cfg.grpcUrl,
    grpc.WithTransportCredentials(insecure.NewCredentials()),
    grpc.WithBlock(),
  )
  if err != nil {
    return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
  }

  return conn, nil
}

func NewGRPCClient(grpcURL string) (GRPCClientInterface, error) {
  conn, err := NewClient(NewDefaultConfig(grpcURL))
  if err != nil {
    return nil, err
  }

  return &GRPCClient{
    conn: conn,
  }, nil
}

func (gc *GRPCClient) GetConnection() *grpc.ClientConn {
  return gc.conn
}

func (gc *GRPCClient) Close() error {
  if gc.conn != nil {
    if err := gc.conn.Close(); err != nil {
      return NewGRPCError("close connection", err)
    }
  }
  return nil
}
