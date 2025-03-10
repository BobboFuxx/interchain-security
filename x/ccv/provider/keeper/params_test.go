package keeper_test

import (
	"testing"
	"time"

	clienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v10/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/v10/modules/light-clients/07-tendermint"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testkeeper "github.com/cosmos/interchain-security/v7/testutil/keeper"
	providertypes "github.com/cosmos/interchain-security/v7/x/ccv/provider/types"
)

// TestParams tests the getting/setting of provider ccv module params.
func TestParams(t *testing.T) {
	// Construct an in-mem keeper with registered key table
	keeperParams := testkeeper.NewInMemKeeperParams(t)
	providerKeeper, ctx, ctrl, _ := testkeeper.GetProviderKeeperAndCtx(t, keeperParams)
	defer ctrl.Finish()

	defaultParams := providertypes.DefaultParams()
	providerKeeper.SetParams(ctx, defaultParams)
	params := providerKeeper.GetParams(ctx)
	require.Equal(t, defaultParams, params)

	newParams := providertypes.NewParams(
		ibctmtypes.NewClientState(
			"",
			ibctmtypes.DefaultTrustLevel,
			0,
			0,
			time.Second*40,
			clienttypes.Height{},
			commitmenttypes.GetSDKSpecs(),
			[]string{"ibc", "upgradedIBCState"},
		),
		"0.25",
		7*24*time.Hour,
		time.Hour,
		"0.4",
		sdk.Coin{
			Denom:  "stake",
			Amount: math.NewInt(10000000),
		},
		600,
		24,
		10,
	)
	providerKeeper.SetParams(ctx, newParams)
	params = providerKeeper.GetParams(ctx)
	require.Equal(t, newParams, params)
}
