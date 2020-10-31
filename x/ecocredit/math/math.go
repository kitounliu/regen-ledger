package math

import (
	"fmt"
	"github.com/cockroachdb/apd/v2"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// IEEE 754-2008 decimal128
var StrictDecima128Context = apd.Context{
	Precision:   34,
	MaxExponent: 6144,
	MinExponent: -6143,
	Traps:       apd.DefaultTraps,
	Rounding:    apd.RoundHalfEven,
}

func IsPositive(x *apd.Decimal) bool {
	return x.Sign() > 0 && !x.IsZero()
}

func IsNegative(x *apd.Decimal) bool {
	return x.Sign() < 0 && !x.IsZero()
}

func MustParsePositiveDecimal(x string) (*apd.Decimal, error) {
	res, _, err := StrictDecima128Context.NewFromString(x)
	if err != nil {
		return nil, errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected a positive decimal, got %s", x))
	}

	if !IsPositive(res) {
		return nil, errors.Wrap(errors.ErrInvalidRequest, fmt.Sprintf("expected a positive decimal, got %s", x))
	}

	return res, nil
}
