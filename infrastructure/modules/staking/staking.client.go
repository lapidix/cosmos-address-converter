package staking

import (
	"context"
	"fmt"

	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"google.golang.org/grpc"
)

type StakingClient interface {
	GetValidator(ctx context.Context, validatorAddr string) (*stakingTypes.QueryValidatorResponse, error)
	Close() error
}

type stakingClient struct {
	conn          *grpc.ClientConn
	stakingClient stakingTypes.QueryClient
}

// NewStakingClient creates a new staking client
func NewStakingClient(conn *grpc.ClientConn) StakingClient {
	return &stakingClient{
		conn:          conn,
		stakingClient: stakingTypes.NewQueryClient(conn),
	}
}

// GetValidator retrieves validator information
func (sc *stakingClient) GetValidator(ctx context.Context, validatorAddr string) (*stakingTypes.QueryValidatorResponse, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context cannot be nil")
	}

	if validatorAddr == "" {
		return nil, fmt.Errorf("validator address cannot be empty")
	}

	resp, err := sc.stakingClient.Validator(ctx, &stakingTypes.QueryValidatorRequest{
		ValidatorAddr: validatorAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get validator: %w", err)
	}
	return resp, nil
}

// Close closes the client connection
func (sc *stakingClient) Close() error {
	if sc.conn != nil {
		return sc.conn.Close()
	}
	return nil
}
