package main

import (
  "fmt"
  "log"

  "github.com/mingi3442/cosmos-key-converter/infrastructure/grpc"
  "github.com/mingi3442/cosmos-key-converter/infrastructure/modules/staking"
  "github.com/mingi3442/cosmos-key-converter/internal/address/domain"
  addressService "github.com/mingi3442/cosmos-key-converter/internal/address/service"
)

func main() {
  grpcClient, err := grpc.NewGRPCClient("cosmos-grpc.polkachu.com:14990")
  if err != nil {
    log.Fatalf("Failed to create gRPC client: %v", err)
  }
  defer grpcClient.Close()

  stakingClient := staking.NewStakingClient(grpcClient.GetConnection())
  converter := addressService.NewAddressConverter(stakingClient)

  accountAddr := address.NewAddress("cosmos1c4k24jzduc365kywrsvf5ujz4ya6mwymy8vq4q")

  validatorAddr, err := converter.ConvertToValidatorAddress(accountAddr)
  if err != nil {
    log.Fatalf("Failed to get validator address: %v", err)
  }
  fmt.Printf("Validator operator address: %s\n", validatorAddr)

  consensusAddr, err := converter.ConvertToConsensusAddress(validatorAddr, accountAddr.Prefix)
  if err != nil {
    log.Fatalf("Failed to get consensus address: %v", err)
  }
  fmt.Printf("Consensus address: %s\n", consensusAddr)
}
