package address

import (
  "context"
  "crypto/sha256"

  sdk "github.com/cosmos/cosmos-sdk/types"
  "github.com/mingi3442/cosmos-key-converter/infrastructure/modules/staking"
  domain "github.com/mingi3442/cosmos-key-converter/internal/address/domain"
)

const (
  valoperSuffix   = "valoper"
  consensusSuffix = "valcons"
)

type addressConverter struct {
  stakingClient staking.StakingClient
}

func NewAddressConverter(stakingClient staking.StakingClient) domain.Converter {
  return &addressConverter{stakingClient: stakingClient}
}

func (ac *addressConverter) ConvertToValidatorAddress(address *domain.Address) (string, error) {
  accAddr, err := sdk.AccAddressFromBech32(address.AccAddress)
  if err != nil {
    return "", &domain.ConversionError{
      Address: address.AccAddress,
      Err:     err,
    }
  }

  valAddr := sdk.ValAddress(accAddr)
  bech32ValAddr, err := sdk.Bech32ifyAddressBytes(address.Prefix+valoperSuffix, valAddr)
  if err != nil {
    return "", &domain.ConversionError{
      Address: address.AccAddress,
      Err:     err,
    }
  }

  return bech32ValAddr, nil
}

func (ac *addressConverter) ConvertToConsensusAddress(validatorAddr string, prefix string) (string, error) {
  ctx := context.Background()

  valAddr, err := sdk.ValAddressFromBech32(validatorAddr)
  if err != nil {
    return "", &domain.ConversionError{
      Address: validatorAddr,
      Err:     err,
    }
  }

  validatorRes, err := ac.stakingClient.GetValidator(ctx, valAddr.String())
  if err != nil || validatorRes == nil || validatorRes.GetValidator().ConsensusPubkey == nil {
    return "", &domain.ConversionError{
      Address: validatorAddr,
      Err:     err,
    }
  }

  pubKeyBytes := validatorRes.GetValidator().ConsensusPubkey.Value[2:]
  hash := sha256.Sum256(pubKeyBytes)
  consAddr := hash[:20]

  bech32ConsAddr, err := sdk.Bech32ifyAddressBytes(prefix+consensusSuffix, consAddr)
  if err != nil {
    return "", &domain.ConversionError{
      Address: validatorAddr,
      Err:     err,
    }
  }

  return bech32ConsAddr, nil
}
