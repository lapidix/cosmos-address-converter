package address

import (
  "context"
)

// ConversionError represents an error that occurred during address conversion
type ConversionError struct {
  Address string
  Err     error
}

func (e *ConversionError) Error() string {
  return "failed to convert address " + e.Address + ": " + e.Err.Error()
}

// ConverterOption represents options for address conversion
type ConverterOption struct {
  Context context.Context
}

// NewConverterOption creates a new ConverterOption with default values
func NewConverterOption() *ConverterOption {
  return &ConverterOption{
    Context: context.Background(),
  }
}
