package grpc

import "fmt"

type GRPCError struct {
  Operation string
  Err       error
}

func (e *GRPCError) Error() string {
  return fmt.Sprintf("gRPC operation '%s' failed: %v", e.Operation, e.Err)
}

func NewGRPCError(operation string, err error) error {
  return &GRPCError{
    Operation: operation,
    Err:       err,
  }
}
