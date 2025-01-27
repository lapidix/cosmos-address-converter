package config

import (
  "fmt"
  "github.com/BurntSushi/toml"
)

// Config는 애플리케이션의 설정을 담는 구조체입니다.
type Config struct {
  GrpcUrl        string `toml:"grpc_url"`
  AccountAddress string `toml:"account_address"`
}

func LoadConfig(configPath string) (*Config, error) {
  var config Config
  if _, err := toml.DecodeFile(configPath, &config); err != nil {
    return nil, fmt.Errorf("failed to decode config file: %w", err)
  }

  if config.GrpcUrl == "" {
    panic("grpc_url is required in config file")
  }
  if config.AccountAddress == "" {
    panic("account_address is required in config file")
  }

  return &config, nil
}
