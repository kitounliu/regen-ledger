package ecocredit

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	//"reflect"
	"testing"
)

func TestDefaultParams(t *testing.T) {
	expected := Params{
		CreditClassFee:               sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultCreditClassFeeTokens)),
		AllowedClassCreatorAddresses: []string{},
		AllowlistEnabled:             true,
	}
	df := DefaultParams()

	require.Equal(t, df.String(), expected.String())
}

func Test_validateAllowlistCreditCreators(t *testing.T) {

	genAddrs := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		genAddrs = append(genAddrs, sdk.MustBech32ifyAddressBytes(sdk.Bech32MainPrefix, []byte(fmt.Sprintf("testaddr-%d", i))))
	}

	tests := []struct {
		name    string
		args    interface{}
		wantErr bool
	}{
		{
			name:    "valid creator list",
			args:    genAddrs,
			wantErr: false,
		},
		{
			name:    "invalid creator list",
			args:    []string{"bogus", "superbogus", "megabogus"},
			wantErr: true,
		},
		{
			name:    "mix of invalid/valid",
			args:    []string{"bogus", genAddrs[0]},
			wantErr: true,
		},
		{
			name:    "wrong type",
			args:    []int{1, 2, 3, 4, 5},
			wantErr: true,
		},
		{
			name:    "not even an array",
			args:    "bruh",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateAllowlistCreditCreators(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("validateAllowlistCreditCreators() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateAllowlistEnabled(t *testing.T) {

	tests := []struct {
		name    string
		args    interface{}
		wantErr bool
	}{
		{
			name:    "valid boolean value",
			args:    true,
			wantErr: false,
		},
		{
			name:    "valid boolean v2",
			args:    false,
			wantErr: false,
		},
		{
			name:    "invalid type",
			args:    []string{"lol", "rofl"},
			wantErr: true,
		},
		{
			name:    "seems valid but its not",
			args:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateAllowlistEnabled(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("validateAllowlistEnabled() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateCreditClassFee(t *testing.T) {
	tests := []struct {
		name    string
		args    interface{}
		wantErr bool
	}{
		{
			name:    "valid credit fee",
			args:    sdk.NewCoins(sdk.NewCoin("regen", sdk.NewInt(45)), sdk.NewCoin("osmo", sdk.NewInt(32))),
			wantErr: false,
		},
		{
			name:    "no fee is valid",
			args:    sdk.NewCoins(),
			wantErr: false,
		},
		{
			name:    "invalid type",
			args:    45,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreditClassFee(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("validateCreditClassFee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
