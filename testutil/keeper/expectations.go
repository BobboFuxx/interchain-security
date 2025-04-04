package keeper

import (
	time "time"

	clienttypes "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v10/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v10/modules/core/04-channel/types"
	ibctmtypes "github.com/cosmos/ibc-go/v10/modules/light-clients/07-tendermint"
	"github.com/golang/mock/gomock"

	math "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	providertypes "github.com/cosmos/interchain-security/v7/x/ccv/provider/types"
	"github.com/cosmos/interchain-security/v7/x/ccv/types"
)

//
// A file containing groups of commonly used mock expectations.
// Note: Each group of mock expectations is associated with a single method
// that may be called during unit tests.
//

// GetMocksForCreateConsumerClient returns mock expectations needed to call CreateConsumerClient().
func GetMocksForCreateConsumerClient(ctx sdk.Context, mocks *MockedKeepers,
	expectedChainID string, expectedLatestHeight clienttypes.Height,
) []*gomock.Call {
	return []*gomock.Call{
		mocks.MockClientKeeper.EXPECT().CreateClient(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return("clientID", nil).Times(1),
	}
}

// GetMocksForMakeConsumerGenesis returns mock expectations needed to call MakeConsumerGenesis().
func GetMocksForMakeConsumerGenesis(ctx sdk.Context, mocks *MockedKeepers,
	unbondingTimeToInject time.Duration, revisionHeight int64,
) []*gomock.Call {
	return []*gomock.Call{
		mocks.MockStakingKeeper.EXPECT().UnbondingTime(gomock.Any()).Return(unbondingTimeToInject, nil).Times(1),
		mocks.MockStakingKeeper.EXPECT().GetHistoricalInfo(gomock.Any(), revisionHeight).Times(1),
	}
}

// GetMocksForSetConsumerChain returns mock expectations needed to call SetConsumerChain().
func GetMocksForSetConsumerChain(ctx sdk.Context, mocks *MockedKeepers,
	chainIDToInject string,
) []*gomock.Call {
	return []*gomock.Call{
		mocks.MockChannelKeeper.EXPECT().GetChannel(ctx, types.ProviderPortID, gomock.Any()).Return(
			channeltypes.Channel{
				State:          channeltypes.OPEN,
				ConnectionHops: []string{"connectionID"},
			},
			true,
		).Times(1),
		mocks.MockConnectionKeeper.EXPECT().GetConnection(ctx, "connectionID").Return(
			conntypes.ConnectionEnd{ClientId: "clientID"}, true,
		).Times(1),
		mocks.MockClientKeeper.EXPECT().GetClientState(ctx, "clientID").Return(
			&ibctmtypes.ClientState{ChainId: chainIDToInject}, true,
		).Times(1),
	}
}

// GetMocksForDeleteConsumerChain returns mock expectations needed to call `DeleteConsumerChain`
func GetMocksForDeleteConsumerChain(ctx sdk.Context, mocks *MockedKeepers) []*gomock.Call {
	return []*gomock.Call{
		mocks.MockChannelKeeper.EXPECT().GetChannel(gomock.Any(), types.ProviderPortID, "channelID").Return(
			channeltypes.Channel{State: channeltypes.OPEN}, true,
		).Times(1),
		mocks.MockChannelKeeper.EXPECT().ChanCloseInit(gomock.Any(), types.ProviderPortID, "channelID").Times(1),
	}
}

func GetMocksForHandleSlashPacket(ctx sdk.Context, mocks MockedKeepers,
	expectedProviderValConsAddr providertypes.ProviderConsAddress,
	valToReturn stakingtypes.Validator, expectJailing bool,
) []*gomock.Call {
	// These first two calls are always made.
	calls := []*gomock.Call{
		mocks.MockStakingKeeper.EXPECT().GetValidatorByConsAddr(
			ctx, expectedProviderValConsAddr.ToSdkConsAddr()).Return(
			valToReturn, nil,
		).Times(1),

		mocks.MockSlashingKeeper.EXPECT().IsTombstoned(ctx,
			expectedProviderValConsAddr.ToSdkConsAddr()).Return(false).Times(1),
	}

	if expectJailing {
		// slash
		calls = append(calls, mocks.MockStakingKeeper.EXPECT().SlashWithInfractionReason(ctx, expectedProviderValConsAddr.ToSdkConsAddr(), gomock.Any(),
			gomock.Any(), gomock.Any(), gomock.Any()).Return(math.NewInt(0), nil).Times(1))
		// jail
		calls = append(calls, mocks.MockStakingKeeper.EXPECT().Jail(
			gomock.Eq(ctx),
			gomock.Eq(expectedProviderValConsAddr.ToSdkConsAddr()),
		).Return(nil))

		// JailUntil is set in this code path.
		calls = append(calls, mocks.MockSlashingKeeper.EXPECT().JailUntil(ctx,
			expectedProviderValConsAddr.ToSdkConsAddr(), gomock.Any()).Return(nil).Times(1))
	}

	return calls
}

func ExpectLatestConsensusStateMock(ctx sdk.Context, mocks MockedKeepers, clientID string, consState *ibctmtypes.ConsensusState) *gomock.Call {
	return mocks.MockClientKeeper.EXPECT().
		GetLatestClientConsensusState(ctx, clientID).Return(consState, true).Times(1)
}

func ExpectCreateClientMock(ctx sdk.Context, mocks MockedKeepers, clientType, clientID string,
	clientState, consState []byte,
) *gomock.Call {
	return mocks.MockClientKeeper.EXPECT().CreateClient(ctx, clientType, clientState, consState).Return(clientID,
		nil).Times(1)
}

func GetMocksForSendIBCPacket(ctx sdk.Context, mocks MockedKeepers, channelID string, times int) []*gomock.Call {
	return []*gomock.Call{
		mocks.MockChannelKeeper.EXPECT().GetChannel(ctx, types.ConsumerPortID,
			"consumerCCVChannelID").Return(channeltypes.Channel{}, true).Times(times),
		mocks.MockChannelKeeper.EXPECT().SendPacket(ctx,
			types.ConsumerPortID,
			"consumerCCVChannelID",
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(uint64(888), nil).Times(times),
	}
}

func GetMocksForSlashValidator(
	ctx sdk.Context,
	mocks MockedKeepers,
	validator stakingtypes.Validator,
	consAddr sdk.ConsAddress,
	undelegations []stakingtypes.UnbondingDelegation,
	redelegations []stakingtypes.Redelegation,
	powerReduction math.Int,
	slashFraction math.LegacyDec,
	currentPower,
	expectedInfractionHeight,
	expectedSlashPower int64,
) []*gomock.Call {
	return []*gomock.Call{
		mocks.MockStakingKeeper.EXPECT().
			GetUnbondingDelegationsFromValidator(ctx, validator.GetOperator()).
			Return(undelegations),
		mocks.MockStakingKeeper.EXPECT().
			GetRedelegationsFromSrcValidator(ctx, validator.GetOperator()).
			Return(redelegations),
		mocks.MockStakingKeeper.EXPECT().
			GetLastValidatorPower(ctx, validator.GetOperator()).
			Return(currentPower),
		mocks.MockStakingKeeper.EXPECT().
			PowerReduction(ctx).
			Return(powerReduction),
		mocks.MockStakingKeeper.EXPECT().
			SlashUnbondingDelegation(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			DoAndReturn(
				func(_ sdk.Context, undelegation stakingtypes.UnbondingDelegation, _ int64, _ math.LegacyDec) math.Int {
					sum := math.NewInt(0)
					for _, r := range undelegation.Entries {
						if r.IsMature(ctx.BlockTime()) {
							continue
						}
						sum = sum.Add(math.NewInt(r.InitialBalance.Int64()))
					}
					return sum
				}).AnyTimes(),
		mocks.MockStakingKeeper.EXPECT().
			SlashRedelegation(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
			DoAndReturn(
				func(_ sdk.Context, _ stakingtypes.Validator, redelegation stakingtypes.Redelegation, _ int64, _ math.LegacyDec) math.Int {
					sum := math.NewInt(0)
					for _, r := range redelegation.Entries {
						if r.IsMature(ctx.BlockTime()) {
							continue
						}
						sum = sum.Add(math.NewInt(r.InitialBalance.Int64()))
					}
					return sum
				}).AnyTimes(),
		mocks.MockSlashingKeeper.EXPECT().
			SlashFractionDoubleSign(ctx).
			Return(slashFraction),
		mocks.MockStakingKeeper.EXPECT().
			SlashWithInfractionReason(ctx, consAddr, expectedInfractionHeight, expectedSlashPower, slashFraction, stakingtypes.Infraction_INFRACTION_DOUBLE_SIGN).
			Times(1),
	}
}

// SetupMocksForLastBondedValidatorsExpectation sets up the expectation for the `GetBondedValidatorsByPower` and `MaxValidators` methods of the `mockStakingKeeper` object.
// These are needed in particular when calling `GetLastBondedValidators` from the provider keeper.
// Times is the number of times the expectation should be called. Provide -1 for `AnyTimes“.
func SetupMocksForLastBondedValidatorsExpectation(mockStakingKeeper *MockStakingKeeper, maxValidators uint32, vals []stakingtypes.Validator, times int) {
	validatorsCall := mockStakingKeeper.EXPECT().GetBondedValidatorsByPower(gomock.Any()).Return(vals, nil)
	maxValidatorsCall := mockStakingKeeper.EXPECT().MaxValidators(gomock.Any()).Return(maxValidators, nil)

	if times == -1 {
		validatorsCall.AnyTimes()
		maxValidatorsCall.AnyTimes()
	} else {
		validatorsCall.Times(times)
		maxValidatorsCall.Times(times)
	}
}
