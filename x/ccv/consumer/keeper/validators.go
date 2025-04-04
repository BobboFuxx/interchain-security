package keeper

import (
	"context"
	"errors"
	"time"

	"cosmossdk.io/math"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/interchain-security/v7/x/ccv/consumer/types"
)

//
// TODO: make unit tests for all of: MVP consumer, democ consumer, and pre-ccv consumer
// for previously unimplemented methods, if they're implemented to solve the above issue.
//

// ApplyCCValidatorChanges applies the given changes to the cross-chain validators states
// and returns updates to forward to tendermint.
func (k Keeper) ApplyCCValidatorChanges(ctx sdk.Context, changes []abci.ValidatorUpdate) []abci.ValidatorUpdate {
	ret := []abci.ValidatorUpdate{}
	for _, change := range changes {
		// convert TM pubkey to SDK pubkey
		pubkey, err := cryptocodec.FromCmtProtoPublicKey(change.GetPubKey())
		if err != nil {
			// An error here would indicate that the validator updates
			// received from the provider are invalid.
			panic(err)
		}
		addr := pubkey.Address()
		val, found := k.GetCCValidator(ctx, addr)

		if found {
			// update or delete an existing validator
			if change.Power < 1 {
				k.DeleteCCValidator(ctx, addr)
			} else {
				val.Power = change.Power
				k.SetCCValidator(ctx, val)
			}
		} else if 0 < change.Power {
			// create a new validator
			consAddr := sdk.ConsAddress(addr)

			ccVal, err := types.NewCCValidator(addr, change.Power, pubkey)
			if err != nil {
				// An error here would indicate that the validator updates
				// received from the provider are invalid.
				panic(err)
			}

			k.SetCCValidator(ctx, ccVal)
			err = k.AfterValidatorBonded(ctx, consAddr, nil)
			if err != nil {
				// AfterValidatorBonded is called by the Slashing module and should not return an error.
				panic(err)
			}
			// Sanity check: making sure the outstanding downtime flag is not
			// set for this new validator. This is especially useful to deal with
			// https://github.com/cosmos/interchain-security/issues/1569.
			k.DeleteOutstandingDowntime(ctx, consAddr)
		} else {
			// edge case: we received an update for 0 power
			// but the validator is already deleted. Do not forward
			// to tendermint.
			continue
		}

		ret = append(ret, change)
	}
	return ret
}

// IterateValidators - unimplemented on CCV keeper but perform a no-op in order to pass the slashing module InitGenesis.
// It is allowed since the condition verifying validator public keys in HandleValidatorSignature (x/slashing/keeper/infractions.go) is removed
// therefore it isn't required to store any validator public keys to the slashing states during genesis.
func (k Keeper) IterateValidators(context.Context, func(index int64, validator stakingtypes.ValidatorI) (stop bool)) error {
	return nil
}

// Validator - unimplemented on CCV keeper but implemented on standalone keeper
func (k Keeper) Validator(sdkCtx context.Context, addr sdk.ValAddress) (stakingtypes.ValidatorI, error) {
	ctx := sdk.UnwrapSDKContext(sdkCtx)
	if k.ChangeoverIsComplete(ctx) && k.standaloneStakingKeeper != nil {
		return k.standaloneStakingKeeper.Validator(ctx, addr)
	}

	return stakingtypes.Validator{}, errors.New("unimplemented on CCV keeper")
}

// IsJailed returns the outstanding slashing flag for the given validator address
func (k Keeper) IsValidatorJailed(goCtx context.Context, addr sdk.ConsAddress) (bool, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// if the changeover is not complete for prev standalone chain,
	// return the standalone staking keeper's jailed status
	if k.IsPrevStandaloneChain(ctx) && !k.ChangeoverIsComplete(ctx) {
		return k.standaloneStakingKeeper.IsValidatorJailed(ctx, addr)
	}
	// Otherwise, return the ccv consumer keeper's notion of a validator being jailed
	return k.OutstandingDowntime(ctx, addr), nil
}

// ValidatorByConsAddr returns an empty validator
func (k Keeper) ValidatorByConsAddr(context.Context, sdk.ConsAddress) (stakingtypes.ValidatorI, error) {
	/*
		NOTE:

		The evidence module will call this function when it handles equivocation evidence.
		The returned value must not be nil and must not have an UNBONDED validator status,
		or evidence will reject it.

		Also, the slashing module will cal lthis function when it observes downtime. In that case
		the only requirement on the returned value is that it isn't null.
	*/
	return stakingtypes.Validator{}, nil
}

// Calls SlashWithInfractionReason with Infraction_INFRACTION_UNSPECIFIED.
// ConsumerKeeper must implement StakingKeeper interface.
// This function should not be called anywhere
func (k Keeper) Slash(ctx context.Context, addr sdk.ConsAddress, infractionHeight, power int64, slashFactor math.LegacyDec) (math.Int, error) {
	return k.SlashWithInfractionReason(ctx, addr, infractionHeight, power, slashFactor, stakingtypes.Infraction_INFRACTION_UNSPECIFIED)
}

// Slash queues a slashing request for the provider chain
// All queued slashing requests will be cleared in EndBlock
// Called by Slashing keeper in SlashWithInfractionReason
func (k Keeper) SlashWithInfractionReason(goCtx context.Context, addr sdk.ConsAddress, infractionHeight, power int64, slashFactor math.LegacyDec, infraction stakingtypes.Infraction) (math.Int, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if infraction == stakingtypes.Infraction_INFRACTION_UNSPECIFIED {
		return math.ZeroInt(), nil
	}

	// If this is a previously standalone chain and infraction happened before the changeover was completed,
	// slash only on the standalone staking keeper.
	if k.IsPrevStandaloneChain(ctx) && infractionHeight < k.FirstConsumerHeight(ctx) {
		// Slash for a standalone chain does not require an infraction reason so we pass in Infraction_INFRACTION_UNSPECIFIED
		return k.standaloneStakingKeeper.SlashWithInfractionReason(ctx, addr, infractionHeight, power, slashFactor, stakingtypes.Infraction_INFRACTION_UNSPECIFIED)
	}

	// Otherwise infraction happened after the changeover was completed.
	// get VSC ID for infraction height
	vscID := k.GetHeightValsetUpdateID(ctx, uint64(infractionHeight))

	k.Logger(ctx).Debug("vscID obtained from mapped infraction height",
		"infraction height", infractionHeight,
		"vscID", vscID,
	)

	// this is the most important step in the function
	// everything else is just here to implement StakingKeeper interface
	// IBC packets are created from slash data and sent to the provider during EndBlock
	k.QueueSlashPacket(
		ctx,
		abci.Validator{
			Address: addr.Bytes(),
			Power:   power,
		},
		vscID,
		infraction,
	)

	// Only return to comply with the interface restriction
	return math.ZeroInt(), nil
}

// Jail - unimplemented on CCV keeper
//
// This method should be a no-op even during a standalone to consumer changeover.
// Once the upgrade has happened as a part of the changeover,
// the provider validator set will soon be in effect, and jailing is n/a.
func (k Keeper) Jail(context.Context, sdk.ConsAddress) error { return nil }

// Unjail is enabled for previously standalone chains and chains implementing democracy staking.
//
// This method should be a no-op for consumer chains that launched with the CCV module first.
func (k Keeper) Unjail(sdkCtx context.Context, addr sdk.ConsAddress) error {
	ctx := sdk.UnwrapSDKContext(sdkCtx)
	if k.ChangeoverIsComplete(ctx) && k.standaloneStakingKeeper != nil {
		return k.standaloneStakingKeeper.Unjail(ctx, addr)
	}
	return nil
}

// Delegation - unimplemented on CCV keeper but implemented on standalone keeper
func (k Keeper) Delegation(sdkCtx context.Context, addr sdk.AccAddress, valAddr sdk.ValAddress) (stakingtypes.DelegationI, error) {
	ctx := sdk.UnwrapSDKContext(sdkCtx)
	if k.ChangeoverIsComplete(ctx) && k.standaloneStakingKeeper != nil {
		return k.standaloneStakingKeeper.Delegation(ctx, addr, valAddr)
	}
	return stakingtypes.Delegation{}, errors.New("unimplemented on CCV keeper")
}

// MaxValidators - unimplemented on CCV keeper
func (k Keeper) MaxValidators(context.Context) (uint32, error) {
	panic("unimplemented on CCV keeper")
}

// UnbondingTime returns consumer unbonding period, satisfying the staking keeper interface
func (k Keeper) UnbondingTime(goCtx context.Context) (time.Duration, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return k.GetUnbondingPeriod(ctx), nil
}

// GetHistoricalInfo gets the historical info at a given height
func (k Keeper) GetHistoricalInfo(goCtx context.Context, height int64) (stakingtypes.HistoricalInfo, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	key := types.HistoricalInfoKey(height)

	value := store.Get(key)
	if value == nil {
		return stakingtypes.HistoricalInfo{}, stakingtypes.ErrNoHistoricalInfo
	}

	return stakingtypes.UnmarshalHistoricalInfo(k.cdc, value)
}

// SetHistoricalInfo sets the historical info at a given height
func (k Keeper) SetHistoricalInfo(goCtx context.Context, height int64, hi *stakingtypes.HistoricalInfo) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	key := types.HistoricalInfoKey(height)
	value := k.cdc.MustMarshal(hi)

	store.Set(key, value)
}

// DeleteHistoricalInfo deletes the historical info at a given height
func (k Keeper) DeleteHistoricalInfo(goCtx context.Context, height int64) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	key := types.HistoricalInfoKey(height)

	store.Delete(key)
	return nil
}

// TrackHistoricalInfo saves the latest historical-info and deletes the oldest
// heights that are below pruning height
func (k Keeper) TrackHistoricalInfo(goCtx context.Context) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	numHistoricalEntries := k.GetHistoricalEntries(ctx)

	// Prune store to ensure we only have parameter-defined historical entries.
	// In most cases, this will involve removing a single historical entry.
	// In the rare scenario when the historical entries gets reduced to a lower value k'
	// from the original value k. k - k' entries must be deleted from the store.
	// Since the entries to be deleted are always in a continuous range, we can iterate
	// over the historical entries starting from the most recent version to be pruned
	// and then return at the first empty entry.
	for i := ctx.BlockHeight() - numHistoricalEntries; i >= 0; i-- {
		_, err := k.GetHistoricalInfo(ctx, i)
		if err != nil {
			if errors.Is(err, stakingtypes.ErrNoHistoricalInfo) {
				break
			}
			return err
		}
		if err = k.DeleteHistoricalInfo(ctx, i); err != nil {
			return err
		}
	}

	// if there is no need to persist historicalInfo, return
	if numHistoricalEntries == 0 {
		return nil
	}

	// Create HistoricalInfo struct
	lastVals := []stakingtypes.Validator{}
	for _, v := range k.GetAllCCValidator(ctx) {
		pk, err := v.ConsPubKey()
		if err != nil {
			// This should never happen as the pubkey is assumed
			// to be stored correctly in ApplyCCValidatorChanges.
			panic(err)
		}

		val, err := stakingtypes.NewValidator(sdk.ValAddress(pk.Address()).String(), pk, stakingtypes.Description{})
		if err != nil {
			// This should never happen as the pubkey is assumed
			// to be stored correctly in ApplyCCValidatorChanges.
			panic(err)
		}

		// Set validator to bonded status
		val.Status = stakingtypes.Bonded
		// Compute tokens from voting power
		val.Tokens = sdk.TokensFromConsensusPower(v.Power, sdk.DefaultPowerReduction)
		lastVals = append(lastVals, val)
	}

	// Create historical info entry which sorts the validator set by voting power
	historicalEntry := stakingtypes.NewHistoricalInfo(ctx.BlockHeader(), stakingtypes.Validators{Validators: lastVals, ValidatorCodec: k.validatorAddressCodec}, sdk.DefaultPowerReduction)

	// Set latest HistoricalInfo at current height
	k.SetHistoricalInfo(ctx, ctx.BlockHeight(), &historicalEntry)
	return nil
}

// MustGetCurrentValidatorsAsABCIUpdates gets all cross-chain validators converted
// to the ABCI validator update type. It panics in case of failure.
func (k Keeper) MustGetCurrentValidatorsAsABCIUpdates(ctx sdk.Context) []abci.ValidatorUpdate {
	vals := k.GetAllCCValidator(ctx)
	valUpdates := make([]abci.ValidatorUpdate, 0, len(vals))
	for _, v := range vals {
		pk, err := v.ConsPubKey()
		if err != nil {
			// This should never happen as the pubkey is assumed
			// to be stored correctly in ApplyCCValidatorChanges.
			panic(err)
		}
		tmPK, err := cryptocodec.ToCmtProtoPublicKey(pk)
		if err != nil {
			// This should never happen as the pubkey is assumed
			// to be stored correctly in ApplyCCValidatorChanges.
			panic(err)
		}
		valUpdates = append(valUpdates, abci.ValidatorUpdate{PubKey: tmPK, Power: v.Power})
	}
	return valUpdates
}

// implement interface method needed for x/genutil in sdk v47
// returns empty updates and err
func (k Keeper) ApplyAndReturnValidatorSetUpdates(context.Context) (updates []abci.ValidatorUpdate, err error) {
	return
}

// GetAllValidators is needed to implement StakingKeeper as expected by the Slashing module since cosmos-sdk/v0.47.x.
// Use GetAllCCValidator in places where access to all cross-chain validators is needed.
func (k Keeper) GetAllValidators(ctx context.Context) ([]stakingtypes.Validator, error) {
	return []stakingtypes.Validator{}, nil
}
