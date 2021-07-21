package ecocredit

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	// TODO: Decide a sensible default value
	DefaultCreditClassFeeTokens  = sdk.NewInt(10000)
	KeyCreditClassFee            = []byte("CreditClassFee")
	KeyAllowlistedCreditCreators = []byte("AllowlistCreditCreators")
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCreditClassFee, &p.CreditClassFee, validateCreditClassFee),
		paramtypes.NewParamSetPair(KeyAllowlistedCreditCreators, &p.AllowedClassCreatorAddresses, validateAllowlistCreditCreators),
	}
}

func validateCreditClassFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if err := v.Validate(); err != nil {
		return err
	}

	return nil
}

func validateAllowlistCreditCreators(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	for _, sAddr := range v {
		_, err := sdk.AccAddressFromBech32(sAddr)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewParams(creditClassFee sdk.Coins, allowlist []string) Params {
	return Params{
		CreditClassFee:               creditClassFee,
		AllowedClassCreatorAddresses: allowlist,
	}
}

// TODO(tyler): what to put for default addresses???
func DefaultParams() Params {
	return NewParams(sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultCreditClassFeeTokens)), []string{})
}
