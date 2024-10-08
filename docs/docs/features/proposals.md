---
sidebar_position: 3
---

# ICS Provider Proposals

Interchain Security introduces the following new governance proposal types to the provider chain.

## `ConsumerAdditionProposal`

:::tip
If you are preparing a `ConsumerAdditionProposal` you can find more information in the [consumer onboarding checklist](../consumer-development/onboarding.md).
:::

`ConsumerAdditionProposal` is used to add new consumer chains. 
When proposals of this type pass governance and the `spawn_time` specified in the proposal is reached, all opted in provider validators are expected to run infrastructure (validator nodes) for the proposed consumer chain. 
Note that for TopN consumer chains, the validators in the top N% of the voting power are automatically opted in at spawn time. 

Minimal example:
```js
{
    // Time on the provider chain at which the consumer chain genesis is finalized and all validators
    // will be responsible for starting their consumer chain validator node.
    "spawn_time": "2023-02-28T20:40:00.000000Z",
    "title": "Add consumer chain",
    "description": ".md description of your chain and all other relevant information",
    "chain_id": "newchain-1",
    "initial_height" : {
        "revision_height": 0,
        "revision_number": 1,
    },
    // Unbonding period for the consumer chain.
    // It should be smaller than that of the provider.
    "unbonding_period": 86400000000000,
    // Timeout period for CCV related IBC packets.
    // Packets are considered timed-out after this interval elapses.
    "ccv_timeout_period": 259200000000000,
    "transfer_timeout_period": 1800000000000,
    "consumer_redistribution_fraction": "0.75",
    "blocks_per_distribution_transmission": 1000,
    "historical_entries": 10000,
    "genesis_hash": "d86d756e10118e66e6805e9cc476949da2e750098fcc7634fd0cc77f57a0b2b0",
    "binary_hash": "376cdbd3a222a3d5c730c9637454cd4dd925e2f9e2e0d0f3702fc922928583f1",
    "distribution_transmission_channel": "channel-123",
    "top_N": 95,
    "validators_power_cap": 0,
    "validator_set_cap": 0,
    "allowlist": [],
    "denylist": [],
    "min_stake": 0,
    "allow_inactive_vals": false
}
```

More examples can be found in the Interchain Security testnet repository [here](https://github.com/cosmos/testnets/blob/master/interchain-security/stopped/baryon-1/proposal-baryon-1.json) and [here](https://github.com/cosmos/testnets/blob/master/interchain-security/stopped/noble-1/start-proposal-noble-1.json).

## `ConsumerRemovalProposal`

`ConsumerRemovalProposal` is used to remove an existing consumer chain.
When proposals of this type pass governance, the consumer chain in question will be gracefully removed from Interchain Security and validators will no longer be required to run infrastructure for the specified chain.

Minimal example:
```js
{
    // the time on the provider chain at which all validators are responsible to stop their consumer chain validator node
    "stop_time": "2023-03-07T12:40:00.000000Z",
    // the chain-id of the consumer chain to be stopped
    "chain_id": "consumerchain-1",
    "title": "This was a great chain",
    "description": "Here is a .md formatted string specifying removal details"
}
```

After the consumer chain removal, the consumer chain is no longer secured by the provider chain.
The consumer chain might continue to produce blocks using the last validator set received from the provider.
However, its validator set can no longer be slashed for any infractions committed on the consumer.
Additional steps are required to completely offboard a consumer chain, such as re-introducing the staking module and removing the provider's validators from the active set.

## `ConsumerModificationProposal`

`ConsumerModificationProposal` is used to change the power shaping parameters of a running consumer chain, as well as to change a Top N running consumer chain to an Opt-In chain and vice versa.
When proposals of this type pass governance, the consumer chain in question would change all its parameters to the ones passed in the `ConsumerModificationProposal`.

For example, given `chain-1` is a TopN consumer chain with, if the following `ConsumerModificationProposal` passes, then `chain-1` would become an Opt-In chain with a 40% validators power cap, a maximum number of 30 validators, and one denylisted validator.
```js
{
    "title": "Modify consumer chain",
    "description": ".md description of your chain and all other relevant information",
    "chain_id": "chain-1",
    "top_N": 0,
    "validators_power_cap": 40,
    "validator_set_cap": 30,
    "allowlist": [],
    "denylist": ["cosmosvalcons1qmq08eruchr5sf5s3rwz7djpr5a25f7xw4mceq"]
}
```

:::warning
If `top_N`, `validators_power_cap`, or some other argument is not included in the proposal, then it is considered
that the default value is set for this argument. For example, if a Top 50% chain wants to only modify `validators_power_cap`
from 35 to 40, then the `ConsumerModificationProposal` would still need to include that `top_N` is 50. Otherwise
`top_N` would be set to its default value of 0, and the chain would become an Opt-In chain.

To be **safe**, always include `top_N` and all the power shaping parameters in your `ConsumerModificationProposal`.
:::

## `ChangeRewardDenomProposal`

`ChangeRewardDenomProposal` is used to update the set of denoms accepted by the provider as rewards.
Note that a `ChangeRewardDenomProposal` will only be accepted on the provider chain if at least one of the `denomsToAdd` or `denomsToRemove` fields is populated with at least one denom. Also, a denom cannot be repeated in both sets.

Minimal example:
```js
{
  "title": "Add uatom as a reward denom",
  "description": "Here is more information about the proposal",
  "denomsToAdd": ["uatom"],
  "denomsToRemove": []
}
```

Besides native provider denoms (e.g., `uatom` for the Cosmos Hub), please use the `ibc/*` denom trace format.
For example, for `untrn` transferred over the path `transfer/channel-569`, the denom trace can be queried using the following command:
```bash
> gaiad query ibc-transfer denom-hash transfer/channel-569/untrn
hash: 0025F8A87464A471E66B234C4F93AEC5B4DA3D42D7986451A059273426290DD5
```
Then use the resulting hash in the `ChangeRewardDenomProposal`, e.g., 
```js
{
  "title": "Add untrn as a reward denom",
  "description": "Here is more information about the proposal",
  "denomsToAdd": ["ibc/0025F8A87464A471E66B234C4F93AEC5B4DA3D42D7986451A059273426290DD5"],
  "denomsToRemove": []
}
```
