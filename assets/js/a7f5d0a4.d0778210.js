"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[7051],{7314:(e,t,o)=>{o.r(t),o.d(t,{assets:()=>d,contentTitle:()=>a,default:()=>l,frontMatter:()=>s,metadata:()=>r,toc:()=>c});var n=o(5893),i=o(1151);const s={sidebar_position:16,title:"Partial Set Security"},a="ADR 015: Partial Set Security",r={id:"adrs/adr-015-partial-set-security",title:"Partial Set Security",description:"Changelog",source:"@site/versioned_docs/version-v6.1.0/adrs/adr-015-partial-set-security.md",sourceDirName:"adrs",slug:"/adrs/adr-015-partial-set-security",permalink:"/interchain-security/v6.1.0/adrs/adr-015-partial-set-security",draft:!1,unlisted:!1,tags:[],version:"v6.1.0",sidebarPosition:16,frontMatter:{sidebar_position:16,title:"Partial Set Security"},sidebar:"tutorialSidebar",previous:{title:"Epochs",permalink:"/interchain-security/v6.1.0/adrs/adr-014-epochs"},next:{title:"Security aggregation",permalink:"/interchain-security/v6.1.0/adrs/adr-016-securityaggregation"}},d={},c=[{value:"Changelog",id:"changelog",level:2},{value:"Status",id:"status",level:2},{value:"Context",id:"context",level:2},{value:"Decision",id:"decision",level:2},{value:"How do consumer chains join?",id:"how-do-consumer-chains-join",level:3},{value:"State &amp; Query",id:"state--query",level:4},{value:"How do validators opt in?",id:"how-do-validators-opt-in",level:3},{value:"State &amp; Query",id:"state--query-1",level:4},{value:"When do validators opt in?",id:"when-do-validators-opt-in",level:4},{value:"How do validators opt out?",id:"how-do-validators-opt-out",level:3},{value:"State &amp; Query",id:"state--query-2",level:4},{value:"When does a consumer chain start?",id:"when-does-a-consumer-chain-start",level:3},{value:"How do we send the partial validator sets to the consumer chains?",id:"how-do-we-send-the-partial-validator-sets-to-the-consumer-chains",level:3},{value:"How do we distribute rewards?",id:"how-do-we-distribute-rewards",level:3},{value:"Misbehaviour",id:"misbehaviour",level:3},{value:"Fraud votes",id:"fraud-votes",level:4},{value:"Double signing",id:"double-signing",level:4},{value:"Downtime",id:"downtime",level:4},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"References",id:"references",level:2}];function h(e){const t={a:"a",code:"code",em:"em",h1:"h1",h2:"h2",h3:"h3",h4:"h4",li:"li",p:"p",pre:"pre",ul:"ul",...(0,i.a)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(t.h1,{id:"adr-015-partial-set-security",children:"ADR 015: Partial Set Security"}),"\n",(0,n.jsx)(t.h2,{id:"changelog",children:"Changelog"}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsx)(t.li,{children:"2024-01-22: Proposed, first draft of ADR."}),"\n"]}),"\n",(0,n.jsx)(t.h2,{id:"status",children:"Status"}),"\n",(0,n.jsx)(t.p,{children:"Accepted"}),"\n",(0,n.jsx)(t.h2,{id:"context",children:"Context"}),"\n",(0,n.jsxs)(t.p,{children:["Currently, in ",(0,n.jsx)(t.em,{children:"Replicated Security"}),", the entire validator set of the provider chain is used to secure consumer chains. There are at least three concerns with this approach.\nFirst, a large number of validators might be forced to validate consumer chains they are not interested in securing.\nSecond, it is costly for small validators to secure additional chains. This concern is only partially addressed through ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/main/docs/docs/adrs/adr-009-soft-opt-out.md",children:"soft opt-out"})," that allows small validators to opt out from validating consumer chains.\nThird and for the above reasons, it is challenging for a new consumer chain to join Replicated Security."]}),"\n",(0,n.jsxs)(t.p,{children:["As a solution, we present ",(0,n.jsx)(t.em,{children:"Partial Set Security"})," (PSS). As the name suggests, PSS allows for every consumer chain to be secured by only a subset of the provider validator set.\nIn what follows we propose the exact steps we need to take to implement PSS. This is a first iteration of PSS, and therefore we present the most minimal solution that make PSS possible."]}),"\n",(0,n.jsx)(t.h2,{id:"decision",children:"Decision"}),"\n",(0,n.jsxs)(t.p,{children:["In Replicated Security, all the provider validators have to secure every consumer chain (with the exception of those validators allowed to opt out through the ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/main/docs/docs/adrs/adr-009-soft-opt-out.md",children:"soft opt-out"})," feature)."]}),"\n",(0,n.jsxs)(t.p,{children:["In PSS, we allow validators to opt in and out of validating any given consumer chain.\nThis has one exception:  we introduce a parameter ",(0,n.jsx)(t.code,{children:"N"})," for each consumer chain and require that the validators in top ",(0,n.jsx)(t.code,{children:"N%"})," of the provider's voting power have to secure the consumer chain.\nValidators outside of the top ",(0,n.jsx)(t.code,{children:"N%"})," can dynamically opt in if they want to validate on the consumer chain."]}),"\n",(0,n.jsxs)(t.p,{children:["For example, if a consumer chain has ",(0,n.jsx)(t.code,{children:"N = 95%"}),", then it ultimately receives the same security it receives today with Replicated Security (with a default ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/main/docs/docs/adrs/adr-009-soft-opt-out.md",children:"SoftOptOutThreshold"})," of 5%).\nOn the other hand, if a consumer chain has ",(0,n.jsx)(t.code,{children:"N = 0%"}),", then no validator is forced to validate the chain, but validators can opt in to do so instead."]}),"\n",(0,n.jsxs)(t.p,{children:["For the remainder of this ADR, we call a consumer chain ",(0,n.jsx)(t.em,{children:"Top N"})," if it has joined as a Top N chain with ",(0,n.jsx)(t.code,{children:"N > 0"})," and ",(0,n.jsx)(t.em,{children:"Opt In"})," chain otherwise.  An Opt In consumer chain is secured only by the validators that have opted in to secure that chain."]}),"\n",(0,n.jsxs)(t.p,{children:["We intend to implement PSS using a feature branch off ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/tree/v4.0.0",children:"v4.0.0 interchain security"}),"."]}),"\n",(0,n.jsx)(t.h3,{id:"how-do-consumer-chains-join",children:"How do consumer chains join?"}),"\n",(0,n.jsxs)(t.p,{children:["As a simplification and to avoid ",(0,n.jsx)(t.a,{href:"https://forum.cosmos.network/t/pss-permissionless-vs-premissioned-lite-opt-in-consumer-chains/12984/17",children:"chain id squatting"}),", a consumer chain can only join PSS through a governance proposal and not in a permissionless way."]}),"\n",(0,n.jsx)(t.p,{children:'However, this proposal type will be modified so that it requires a lower quorum percentage than normal proposal, and every validator who voted "YES" on the proposal will form the consumer chain\'s initial validator set.'}),"\n",(0,n.jsxs)(t.p,{children:["Consumer chains join PSS the same way chains now join Replicated Security, namely through a ",(0,n.jsx)(t.code,{children:"ConsumerAdditionProposal"})," proposal.\nWe extend ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/proto/interchain_security/ccv/provider/v1/provider.proto#L27",children:(0,n.jsx)(t.code,{children:"ConsumerAdditionProposal"})})," with one optional field:"]}),"\n",(0,n.jsxs)(t.p,{children:[(0,n.jsx)(t.code,{children:"uint32 top_N"}),": Corresponds to the percentage of validators that join under the Top N case.\nFor example, ",(0,n.jsx)(t.code,{children:"53"})," corresponds to a Top 53% chain, meaning that the top ",(0,n.jsx)(t.code,{children:"53%"})," provider validators have to validate the proposed consumer chain.\n",(0,n.jsx)(t.code,{children:"top_N"}),"  can be ",(0,n.jsx)(t.code,{children:"0"})," or include any value in ",(0,n.jsx)(t.code,{children:"[50, 100]"}),". A chain can join with ",(0,n.jsx)(t.code,{children:"top_N == 0"})," as an Opt In, or with ",(0,n.jsx)(t.code,{children:"top_N \u2208 [50, 100]"})," as a Top N chain."]}),"\n",(0,n.jsxs)(t.p,{children:["In case of a Top N chain, we restrict the possible values of ",(0,n.jsx)(t.code,{children:"top_N"})," from ",(0,n.jsx)(t.code,{children:"(0, 100]"})," to ",(0,n.jsx)(t.code,{children:"[50, 100]"}),".\nBy having ",(0,n.jsx)(t.code,{children:"top_N >= 50"})," we can guarantee that we cannot have a successful attack, assuming that at most ",(0,n.jsx)(t.code,{children:"1/3"})," of provider validators can be malicious.\nThis is because, a Top N chain with ",(0,n.jsx)(t.code,{children:"N >= 50%"})," would have at least ",(0,n.jsx)(t.code,{children:"1/3"})," honest validators, which is sufficient to stop attacks.\nAdditionally, by having ",(0,n.jsx)(t.code,{children:"N >= 50%"})," (and hence ",(0,n.jsx)(t.code,{children:"N > (VetoThreshold = 33.4%)"}),") we enable the top N validators to ",(0,n.jsx)(t.code,{children:"Veto"})," any ",(0,n.jsx)(t.code,{children:"ConsumerAdditionProposal"})," for consumer chains they do not want to validate."]}),"\n",(0,n.jsxs)(t.p,{children:["If a proposal has the ",(0,n.jsx)(t.code,{children:"top_N"})," argument wrongly set, it should get rejected in [ValidateBasic] (",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/x/ccv/provider/types/proposal.go#L86",children:"https://github.com/cosmos/interchain-security/blob/v4.0.0/x/ccv/provider/types/proposal.go#L86"}),")."]}),"\n",(0,n.jsxs)(t.p,{children:["In the code, we distinguish whether a chain is ",(0,n.jsx)(t.em,{children:"Top N"})," or ",(0,n.jsx)(t.em,{children:"Opt In"})," by checking whether ",(0,n.jsx)(t.code,{children:"top_N"})," is zero or not."]}),"\n",(0,n.jsxs)(t.p,{children:["In a future version of PSS, we intend to introduce a ",(0,n.jsx)(t.code,{children:"ConsumerModificationProposal"})," so that we can modify the parameters of a consumer chain, e.g, a chain that is ",(0,n.jsx)(t.em,{children:"Opt In"})," to become ",(0,n.jsx)(t.em,{children:"Top N"}),", etc."]}),"\n",(0,n.jsx)(t.h4,{id:"state--query",children:"State & Query"}),"\n",(0,n.jsxs)(t.p,{children:["We augment the provider module\u2019s state to keep track of the ",(0,n.jsx)(t.code,{children:"top_N"})," value for each consumer chain. The key to store this information would be:"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{children:"topNBytePrefix | len(chainID) | chainID\n"})}),"\n",(0,n.jsxs)(t.p,{children:["To create the above key, we can use ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/x/ccv/provider/types/keys.go#L418",children:(0,n.jsx)(t.code,{children:"ChainIdWithLenKey"})}),"."]}),"\n",(0,n.jsxs)(t.p,{children:["Then in the ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/x/ccv/provider/keeper/keeper.go",children:"keeper"})," we introduce methods as follows:"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-golang",children:"func (k Keeper) SetTopN(ctx sdk.Context, chainID string, topN uint32)\nfunc (k Keeper) IsTopN(ctx sdk.Context, chainID string) bool\nfunc (k Keeper) IsOptIn(ctx sdk.Context, chainID string) bool\n\n// returns the N if Top N chain, otherwise an error\nfunc (k Keeper) GetTopN(ctx sdk.Context, chainID string) (uint32, error)\n"})}),"\n",(0,n.jsxs)(t.p,{children:["We also extend the ",(0,n.jsx)(t.code,{children:"interchain-security-pd query provider list-consumer-chains"}),' query to return information on whether a consumer chain is an Opt In or a Top N chain and with what N.\nThis way, block explorers can present informative messages such as "This chain is secured by N% of the provider chain" for consumer chains.']}),"\n",(0,n.jsx)(t.h3,{id:"how-do-validators-opt-in",children:"How do validators opt in?"}),"\n",(0,n.jsxs)(t.p,{children:["A validator can opt in by sending a new type of message that we introduce in ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/proto/interchain_security/ccv/provider/v1/tx.proto#L1",children:"tx.proto"}),"."]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-protobuf",children:"message MsgOptIn {\n    // the chain id of the consumer chain to opt in to\n    string chainID = 1;\n    // the provider address of the validator\n    string providerAddr = 2;\n    // (optional) the consensus public key to use on the consumer\n    optional string consumerKey = 3;\n}\n"})}),"\n",(0,n.jsxs)(t.p,{children:["Note that in a Top N consumer chain, the top ",(0,n.jsx)(t.code,{children:"N%"})," provider validators have to validate the consumer chain.\nNevertheless, validators in the bottom ",(0,n.jsx)(t.code,{children:"(100 - N)%"})," can opt in to validate as well.\nProvider validators that belong or enter the top ",(0,n.jsx)(t.code,{children:"N%"})," validators are ",(0,n.jsx)(t.em,{children:"automatically"})," opted in to validate a Top N consumer chain.\nThis means that if a validator ",(0,n.jsx)(t.code,{children:"V"})," belongs to the top ",(0,n.jsx)(t.code,{children:"N%"})," validators but later falls (e.g., due to undelegations) to the bottom ",(0,n.jsx)(t.code,{children:"(100 - N)%"}),", ",(0,n.jsx)(t.code,{children:"V"})," is still considered opted in and has to validate unless ",(0,n.jsx)(t.code,{children:"V"})," sends a ",(0,n.jsx)(t.code,{children:"MsgOptOut"})," message (see below).\nBy automatically opting in validators when they enter the top ",(0,n.jsx)(t.code,{children:"N%"})," validators and by forcing top ",(0,n.jsx)(t.code,{children:"N%"})," validators to explicitly opt out in case they fall to the ",(0,n.jsx)(t.code,{children:"(100 - N)%"})," bottom validators we simplify the design of PSS."]}),"\n",(0,n.jsxs)(t.p,{children:["Note that a validator can send a ",(0,n.jsx)(t.code,{children:"MsgOptIn"})," message even if the consumer chain is not yet running. To do this we reuse the ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/x/ccv/provider/keeper/key_assignment.go#L644",children:(0,n.jsx)(t.code,{children:"IsConsumerProposedOrRegistered"})}),". If the ",(0,n.jsx)(t.code,{children:"chainID"})," does not exist, the ",(0,n.jsx)(t.code,{children:"MsgOptIn"})," should fail, as well as if the provider address does not exist."]}),"\n",(0,n.jsxs)(t.p,{children:["Optionally, a validator that opts in can provide a ",(0,n.jsx)(t.code,{children:"consumerKey"})," so that it assigns a different consumer key (from the provider) to the consumer chain.\nNaturally, a validator can always change the consumer key on a consumer chain by sending a ",(0,n.jsx)(t.code,{children:"MsgAssignConsumerKey"})," message at a later point in time, as is done in Replicated Security."]}),"\n",(0,n.jsx)(t.h4,{id:"state--query-1",children:"State & Query"}),"\n",(0,n.jsxs)(t.p,{children:["For each validator, we store a pair ",(0,n.jsx)(t.code,{children:"(blockHeight, isOptedIn)"})," that contains the block height the validator opted in and whether the validator is currently opted in or not, under the key:"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{children:"optedInBytePrefix | len(chainID) | chainID | addr\n"})}),"\n",(0,n.jsxs)(t.p,{children:["By using a prefix iterator on ",(0,n.jsx)(t.code,{children:"optedInBytePrefix | len(chainID) | chainID"})," we retrieve all the opted in validators."]}),"\n",(0,n.jsxs)(t.p,{children:["We introduce the following ",(0,n.jsx)(t.code,{children:"Keeper"})," methods."]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-golang",children:"// returns all the validators that have opted in on chain `chainID`\nfunc (k Keeper) GetOptedInValidators(ctx sdk.Context, chainID string) []Validators\n\nfunc (k Keeper) IsValidatorOptedIn(ctx sdk.Context, chainID string, val Validator) bool\n"})}),"\n",(0,n.jsx)(t.p,{children:"We introduce the following two queries:"}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-bash",children:"interchain-security-pd query provider optedInValidators $chainID\ninterchain-security-pd query provider hasToValidate $providerAddr\n"})}),"\n",(0,n.jsx)(t.p,{children:"One query to retrieve the validators that are opted in and hence the validators that need to validate the consumer chain and one query that given a validator's address returns all the chains this validator has to validate."}),"\n",(0,n.jsx)(t.h4,{id:"when-do-validators-opt-in",children:"When do validators opt in?"}),"\n",(0,n.jsxs)(t.p,{children:["As described earlier, validators can manually opt in by sending a ",(0,n.jsx)(t.code,{children:"MsgOptIn"})," message.\nAdditionally, in a Top N chain, a validator is automatically opted in when it moves from the bottom ",(0,n.jsx)(t.code,{children:"(100 - N)%"})," to the top ",(0,n.jsx)(t.code,{children:"N%"})," validators."]}),"\n",(0,n.jsxs)(t.p,{children:["Lastly, validators can also opt in if they vote ",(0,n.jsx)(t.code,{children:"Yes"})," during the ",(0,n.jsx)(t.code,{children:"ConsumerAdditionProposal"})," that introduces a consumer chain.\nThis simplifies validators operations because they do not have to send an additional message to opt in."]}),"\n",(0,n.jsxs)(t.p,{children:["Because the ",(0,n.jsx)(t.code,{children:"Tally"})," method ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.7/x/gov/keeper/tally.go#L71",children:"deletes the votes"})," after reading them, we cannot check the votes of the validators after the votes have been tallied.\nTo circumvent this, we introduce a hook for ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.7/x/gov/keeper/vote.go#L35",children:(0,n.jsx)(t.code,{children:"AfterProposalVote"})})," and keep track of all the votes cast by a validator.\nIf a validator casts more than one vote, we only consider the latest vote.\nFinally, we only consider a validator has opted in if it casts a 100% ",(0,n.jsx)(t.code,{children:"Yes"})," vote in case of a ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/cosmos-sdk/blob/main/docs/architecture/adr-037-gov-split-vote.md",children:"weighted vote"}),"."]}),"\n",(0,n.jsx)(t.h3,{id:"how-do-validators-opt-out",children:"How do validators opt out?"}),"\n",(0,n.jsx)(t.p,{children:"Validators that have opted in on a chain can opt out by sending the following message:"}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-protobuf",children:"message MsgOptOut {\n    // the chain id of the consumer chain to opt out from\n    string chainID = 1;\n    // the provider address of the validator\n    string providerAddr = 2;\n}\n"})}),"\n",(0,n.jsxs)(t.p,{children:["Validators can only opt out after a consumer chain has started and hence the above message returns an error if the chain with ",(0,n.jsx)(t.code,{children:"chainID"})," is not running.\nAdditionally, a validator that belongs to the top ",(0,n.jsx)(t.code,{children:"N%"})," validators cannot opt out from a Top N chain and hence a ",(0,n.jsx)(t.code,{children:"MsgOptOut"})," would error in such a case."]}),"\n",(0,n.jsx)(t.h4,{id:"state--query-2",children:"State & Query"}),"\n",(0,n.jsx)(t.p,{children:"We also update the state of the opted-in validators when a validator has opted out by removing the opted-out validator."}),"\n",(0,n.jsxs)(t.p,{children:["Note that only opted-in validators can be punished for downtime on a consumer chain.\nFor this, we use historical info of all the validators that have opted in; We can examine the ",(0,n.jsx)(t.code,{children:"blockHeight"})," stored under the key ",(0,n.jsx)(t.code,{children:"optedInBytePrefix | len(chainID) | chainID | addr"})," to see if a validator was opted in.\nThis way we can jail validators for downtime knowing that indeed the validators have opted in at some point in the past.\nOtherwise, we can think of a scenario where a validator ",(0,n.jsx)(t.code,{children:"V"})," is down for a period of time, but before ",(0,n.jsx)(t.code,{children:"V"})," gets punished for downtime, validator ",(0,n.jsx)(t.code,{children:"V"})," opts out, and then we do not know whether ",(0,n.jsx)(t.code,{children:"V"})," should be punished or not."]}),"\n",(0,n.jsx)(t.h3,{id:"when-does-a-consumer-chain-start",children:"When does a consumer chain start?"}),"\n",(0,n.jsxs)(t.p,{children:["A Top N consumer chain always starts at the specified date (",(0,n.jsx)(t.code,{children:"spawn_time"}),") if the ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/proto/interchain_security/ccv/provider/v1/provider.proto#L27",children:(0,n.jsx)(t.code,{children:"ConsumerAdditionProposal"})})," has passed.\nAn Opt In consumer chain only starts if at least one validator has opted in. We check this in ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/x/ccv/provider/keeper/proposal.go#L357",children:"BeginBlockInit"}),":"]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-golang",children:'func (k Keeper) BeginBlockInit(ctx sdk.Context) {\n    propsToExecute := k.GetConsumerAdditionPropsToExecute(ctx)\n\n    for _, prop := range propsToExecute {\n        chainID := prop.ChainId\n        if !k.IsTopN(ctx, chainID) && len(k.GetOptedInValidators(ctx, chainID)) == 0 {\n            // drop the proposal\n            ctx.Logger().Info("could not start chain because no validator has opted in")\n            continue\n        }   \n        ...\n'})}),"\n",(0,n.jsx)(t.h3,{id:"how-do-we-send-the-partial-validator-sets-to-the-consumer-chains",children:"How do we send the partial validator sets to the consumer chains?"}),"\n",(0,n.jsxs)(t.p,{children:["A consumer chain should only be validated by opted in validators.\nWe introduce logic to do this when we ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/x/ccv/provider/keeper/relay.go#L213",children:"queue"})," the ",(0,n.jsx)(t.code,{children:"VSCPacket"}),"s.\nThe logic behind this, is not as straightforward as it seems because CometBFT does not receive the validator set that has to validate a chain, but rather a delta of ",(0,n.jsx)(t.a,{href:"https://docs.cometbft.com/v0.37/spec/abci/abci++_methods#validatorupdate",children:"validator updates"}),".\nFor example, to remove an opted-out validator from a consumer chain, we have to send a validator update with a ",(0,n.jsx)(t.code,{children:"power"})," of ",(0,n.jsx)(t.code,{children:"0"}),", similarly to what is done in the ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/x/ccv/provider/keeper/key_assignment.go#L525",children:"assignment of consumer keys"}),".\nWe intend to update this ADR at a later stage on how exactly we intend to implement this logic."]}),"\n",(0,n.jsx)(t.h3,{id:"how-do-we-distribute-rewards",children:"How do we distribute rewards?"}),"\n",(0,n.jsxs)(t.p,{children:["Currently, rewards are distributed as follows: The consumer ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/x/ccv/consumer/keeper/distribution.go#L148",children:"periodically sends rewards"})," on the provider ",(0,n.jsx)(t.code,{children:"ConsumerRewardsPool"})," address.\nThe provider then ",(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/blob/v4.0.0/x/ccv/provider/keeper/distribution.go#L77",children:"transfers those rewards to the fee collector address"})," and those transferred rewards are distributed to validators and delegators."]}),"\n",(0,n.jsx)(t.p,{children:"In PSS, we distribute rewards only to validators that actually validate the consumer chain.\nTo do this, we have a pool associated with each consumer chain and consumers IBC transfer the rewards to this pool.\nWe then extract the rewards from each consumer pool and distribute them to the opted in validators."}),"\n",(0,n.jsx)(t.p,{children:"Note that we only distribute rewards to validators that have been opted in for some time (e.g., 10000 blocks) to avoid cases where validators opt in just to receive rewards and then opt out immediately afterward."}),"\n",(0,n.jsx)(t.h3,{id:"misbehaviour",children:"Misbehaviour"}),"\n",(0,n.jsx)(t.h4,{id:"fraud-votes",children:"Fraud votes"}),"\n",(0,n.jsxs)(t.p,{children:["In an Opt In chain, a set of validators might attempt to perform an attack. To deter such potential attacks, PSS allows for the use of fraud votes.\nA ",(0,n.jsx)(t.em,{children:"fraud vote"})," is a governance proposal that enables the slashing of validators that performed an attack.\nDue to their inherent complexity, we intend to introduce fraud votes in a different ADR and at a future iteration of PSS."]}),"\n",(0,n.jsx)(t.h4,{id:"double-signing",children:"Double signing"}),"\n",(0,n.jsx)(t.p,{children:"We do not change the way slashing for double signing and light client attacks functions.\nIf a validator misbehaves on a consumer, then we slash that validator on the provider."}),"\n",(0,n.jsx)(t.h4,{id:"downtime",children:"Downtime"}),"\n",(0,n.jsx)(t.p,{children:"We do not change the way downtime jailing functions.\nIf a validator is down on a consumer chain for an adequate amount of time, we jail this validator on the provider but only if the validator was opted in on this consumer chain in the recent past."}),"\n",(0,n.jsx)(t.h2,{id:"consequences",children:"Consequences"}),"\n",(0,n.jsx)(t.h3,{id:"positive",children:"Positive"}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsxs)(t.li,{children:["\n",(0,n.jsx)(t.p,{children:"Easier for new consumer chains to consume the provider's chain economic security because proposals are more likely to pass if not everyone is forced to validate."}),"\n"]}),"\n",(0,n.jsxs)(t.li,{children:["\n",(0,n.jsx)(t.p,{children:"Smaller validators are not forced to validate chains anymore if they do not want to."}),"\n"]}),"\n",(0,n.jsxs)(t.li,{children:["\n",(0,n.jsx)(t.p,{children:"We can deprecate the soft opt-out implementation."}),"\n"]}),"\n"]}),"\n",(0,n.jsx)(t.h3,{id:"negative",children:"Negative"}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsxs)(t.li,{children:["A consumer chain does not receive the same economic security as with Replicated Security (assuming the value of ",(0,n.jsx)(t.code,{children:"SoftOptOutThreshold"})," is ",(0,n.jsx)(t.code,{children:"5%"}),"), unless it is a Top N chain with ",(0,n.jsx)(t.code,{children:"N >= 95%"}),"."]}),"\n"]}),"\n",(0,n.jsx)(t.h2,{id:"references",children:"References"}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsx)(t.li,{children:(0,n.jsx)(t.a,{href:"https://forum.cosmos.network/t/pss-permissionless-vs-premissioned-lite-opt-in-consumer-chains/12984",children:"PSS: Permissionless vs premissioned-lite opt-in consumer chains"})}),"\n",(0,n.jsx)(t.li,{children:(0,n.jsx)(t.a,{href:"https://forum.cosmos.network/t/chips-discussion-phase-partial-set-security-updated/11775",children:"CHIPs discussion phase: Partial Set Security (updated)"})}),"\n",(0,n.jsx)(t.li,{children:(0,n.jsx)(t.a,{href:"https://forum.cosmos.network/t/pss-exclusive-vs-inclusive-top-n/13058",children:"PSS: Exclusive vs Inclusive Top-N"})}),"\n",(0,n.jsx)(t.li,{children:(0,n.jsx)(t.a,{href:"https://github.com/cosmos/interchain-security/pull/1518",children:"Initial PSS ADR and notes #1518"})}),"\n",(0,n.jsx)(t.li,{children:(0,n.jsx)(t.a,{href:"https://informal.systems/blog/replicated-vs-mesh-security",children:"Replicated vs. Mesh Security"})}),"\n"]})]})}function l(e={}){const{wrapper:t}={...(0,i.a)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(h,{...e})}):h(e)}},1151:(e,t,o)=>{o.d(t,{Z:()=>r,a:()=>a});var n=o(7294);const i={},s=n.createContext(i);function a(e){const t=n.useContext(s);return n.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function r(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:a(e.components),n.createElement(s.Provider,{value:t},e.children)}}}]);