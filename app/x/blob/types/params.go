package types

import (
	"fmt"

	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMinSquareSize            = []byte("MinSquareSize")
	DefaultMinSquareSize uint32 = appconsts.DefaultMinSquareSize
)

var (
	KeyMaxSquareSize            = []byte("MaxSquareSize")
	DefaultMaxSquareSize uint32 = appconsts.DefaultMaxSquareSize
)

var (
	KeyGasPerBlobByte            = []byte("GasPerBlobByte")
	DefaultGasPerBlobByte uint32 = appconsts.DefaultGasPerBlobByte
)

// ParamKeyTable returns the param key table for the blob module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	minSquareSize uint32,
	maxSquareSize uint32,
	GasPerBlobByte uint32,
) Params {
	return Params{
		MinSquareSize:  minSquareSize,
		MaxSquareSize:  maxSquareSize,
		GasPerBlobByte: GasPerBlobByte,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinSquareSize,
		DefaultMaxSquareSize,
		DefaultGasPerBlobByte,
	)
}

// ParamSetPairs gets the list of param key-value pairs
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinSquareSize, &p.MinSquareSize, validateMinSquareSize),
		paramtypes.NewParamSetPair(KeyMaxSquareSize, &p.MaxSquareSize, validateMaxSquareSize),
		paramtypes.NewParamSetPair(KeyGasPerBlobByte, &p.GasPerBlobByte, validateGasPerBlobByte),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMinSquareSize(p.MinSquareSize); err != nil {
		return err
	}

	if err := validateMaxSquareSize(p.MaxSquareSize); err != nil {
		return err
	}

	if err := validateMinMaxSquareSizeOrder(p.MinSquareSize, p.MaxSquareSize); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMinSquareSize validates the MinSquareSize param
func validateMinSquareSize(v interface{}) error {
	minSquareSize, ok := v.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if minSquareSize == 0 {
		return fmt.Errorf("min square size cannot be 0")
	}

	return nil
}

// validateMaxSquareSize validates the MaxSquareSize param
func validateMaxSquareSize(v interface{}) error {
	maxSquareSize, ok := v.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if maxSquareSize == 0 {
		return fmt.Errorf("max square size cannot be 0")
	}

	return nil
}

// validateMinMaxSquareSizeOrder validates that minSquareSize is less than or equal to maxSquareSize
func validateMinMaxSquareSizeOrder(minSquareSize, maxSquareSize uint32) error {
	if minSquareSize > maxSquareSize {
		return fmt.Errorf("min square size cannot be greater than max square size")
	}
	return nil
}

// validateGasPerBlobByte validates the GasPerBlobByte param
func validateGasPerBlobByte(v interface{}) error {
	GasPerBlobByte, ok := v.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	if GasPerBlobByte == 0 {
		return fmt.Errorf("gas per blob byte cannot be 0")
	}

	return nil
}
