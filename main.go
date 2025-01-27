package main

import (
  "fmt"
  "log"

  "github.com/mingi3442/cosmos-key-converter/config"
  "github.com/mingi3442/cosmos-key-converter/infrastructure/grpc"
  "github.com/mingi3442/cosmos-key-converter/infrastructure/modules/staking"
  "github.com/mingi3442/cosmos-key-converter/internal/address/domain"
  addressService "github.com/mingi3442/cosmos-key-converter/internal/address/service"
)

func main() {
  cfg, err := config.LoadConfig("config.toml")
  if err != nil {
    log.Fatalf("Failed to load config: %v", err)
  }

  grpcClient, err := grpc.NewGRPCClient(cfg.GrpcUrl)
  if err != nil {
    log.Fatalf("Failed to create gRPC client: %v", err)
  }
  defer grpcClient.Close()

  stakingClient := staking.NewStakingClient(grpcClient.GetConnection())
  converter := addressService.NewAddressConverter(stakingClient)

  accountAddr := address.NewAddress(cfg.AccountAddress)

  validatorAddr, err := converter.ConvertToValidatorAddress(accountAddr)
  if err != nil {
    log.Fatalf("Failed to get validator address: %v", err)
  }

  consensusAddr, err := converter.ConvertToConsensusAddress(validatorAddr, accountAddr.Prefix)
  if err != nil {
    log.Fatalf("Failed to get consensus address: %v", err)
  }

  fmt.Printf("Account address:             %s\n", accountAddr.AccAddress)
  fmt.Printf("Validator address:           %s\n", validatorAddr)
  fmt.Printf("Consensus address:           %s\n", consensusAddr)
}
