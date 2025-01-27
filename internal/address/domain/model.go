package address

import "strings"

type Address struct {
  AccAddress string
  Prefix     string
}

type Converter interface {
  ConvertToValidatorAddress(address *Address) (string, error)
  ConvertToConsensusAddress(validatorAddr string, prefix string) (string, error)
}

func NewAddress(accAddress string) *Address {
  return &Address{
    AccAddress: accAddress,
    Prefix:     getPrefix(accAddress),
  }
}

func getPrefix(accAddress string) string {
  if accAddress == "" {
    return ""
  }
  parts := strings.Split(accAddress, "1")
  if len(parts) < 2 {
    return ""
  }
  return parts[0]
}
