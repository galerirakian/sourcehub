package embedded

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// acpAccountKeeper implements ACP's expected AccountKeeper.
// For the time being it stubs an account as the Msg's have no need to be signed.
type acpAccountKeeper struct{}

func (s *acpAccountKeeper) GetAccount(ctx context.Context, address sdk.AccAddress) sdk.AccountI {
	acc := &types.BaseAccount{
		Address:       address.String(),
		Sequence:      1,
		AccountNumber: 1,
	}
	return acc
}
