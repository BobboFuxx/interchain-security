"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[426],{4968:(e,n,i)=>{i.r(n),i.d(n,{assets:()=>l,contentTitle:()=>a,default:()=>c,frontMatter:()=>s,metadata:()=>r,toc:()=>h});var t=i(5893),o=i(1151);const s={sidebar_position:21,title:"Customizable Slashing and Jailing"},a="ADR 020: Customizable Slashing and Jailing",r={id:"adrs/adr-020-cutomizable_slashing_and_jailing",title:"Customizable Slashing and Jailing",description:"Changelog",source:"@site/versioned_docs/version-v4.5.0/adrs/adr-020-cutomizable_slashing_and_jailing.md",sourceDirName:"adrs",slug:"/adrs/adr-020-cutomizable_slashing_and_jailing",permalink:"/interchain-security/v4.5.0/adrs/adr-020-cutomizable_slashing_and_jailing",draft:!1,unlisted:!1,tags:[],version:"v4.5.0",sidebarPosition:21,frontMatter:{sidebar_position:21,title:"Customizable Slashing and Jailing"},sidebar:"tutorialSidebar",previous:{title:"Permissionless ICS",permalink:"/interchain-security/v4.5.0/adrs/adr-019-permissionless-ics"},next:{title:"Consumer Chain Clients",permalink:"/interchain-security/v4.5.0/adrs/adr-021-consumer-chain-clients"}},l={},h=[{value:"Changelog",id:"changelog",level:2},{value:"Status",id:"status",level:2},{value:"Context",id:"context",level:2},{value:"The Cost of PoS",id:"the-cost-of-pos",level:3},{value:"Decision",id:"decision",level:2},{value:"Implementation",id:"implementation",level:3},{value:"Consequences",id:"consequences",level:2},{value:"Positive",id:"positive",level:3},{value:"Negative",id:"negative",level:3},{value:"Economic Security Model without Slashing",id:"economic-security-model-without-slashing",level:4},{value:"Neutral",id:"neutral",level:3},{value:"References",id:"references",level:2}];function d(e){const n={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",h4:"h4",li:"li",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,o.a)(),...e.components};return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)(n.h1,{id:"adr-020-customizable-slashing-and-jailing",children:"ADR 020: Customizable Slashing and Jailing"}),"\n",(0,t.jsx)(n.h2,{id:"changelog",children:"Changelog"}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsx)(n.li,{children:"2024-07-19: Initial draft of ADR"}),"\n",(0,t.jsx)(n.li,{children:"2024-08-23: Generalize ADR to make slashing and jailing customizable"}),"\n"]}),"\n",(0,t.jsx)(n.h2,{id:"status",children:"Status"}),"\n",(0,t.jsx)(n.p,{children:"Accepted"}),"\n",(0,t.jsx)(n.h2,{id:"context",children:"Context"}),"\n",(0,t.jsx)(n.p,{children:"Interchain Security (ICS) is a cross-chain staking protocol -- it uses the stake on the provider chain as collateral for the Proof of Stake (PoS) on the consumer chains.\nThis means that the voting power of validators validating (aka producing blocks) on the consumer chains is a function of their stake on the provider.\nMoreover, if these validators misbehave on the consumer chains, they get punished on the provider chain.\nICS is currently differentiating between two types of infractions -- equivocation and downtime.\nDepending on the infraction type, the misbehaving validators might be jailed (i.e., removed from the provider validator set) and / or slashed (i.e., a portion of their stake on the provider is being burned).\nFor example, validators double signing on consumer chains get slashed and are permanently jailed,\nwhile validators not validating sufficient blocks are temporarily jailed."}),"\n",(0,t.jsx)(n.p,{children:"This means that ICS consumer chains get their economical security from the provider.\nHowever, this might come at a high cost."}),"\n",(0,t.jsx)(n.h3,{id:"the-cost-of-pos",children:"The Cost of PoS"}),"\n",(0,t.jsx)(n.p,{children:"One of the cost of validating on the consumer chains is operational -- validators need to run and monitor full nodes of every consumer chain they opt in for.\nAlthough this cost varies from validator team to validator team (depending on how efficiently they can run their infrastructure), it doesn't depend on the total stake (or voting power) of the validators, so we can think of it as constant.\nThe other cost of validating comes from the risk of getting slashed or jailed."}),"\n",(0,t.jsx)(n.p,{children:"Most chains in Cosmos (including the Cosmos Hub) use delegated PoS -- users delegate their tokens to validators, which stake them in return for voting power.\nTherefore, validators act as representatives chosen by their delegators to represent their interests.\nHowever, delegators share the risk of their validators getting slashed or jailed:"}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsx)(n.li,{children:"When validators get slashed, a portion of their stake is being burned, including a portion of the tokens delegated by users.\nAs validators don't need to have their own stake, it is possible that delegators take all the risk of validators misbehaving."}),"\n",(0,t.jsx)(n.li,{children:"When validators get jailed, they no longer receive block rewards (neither from the provider nor from the consumer chains).\nThis also applies to their delegators.\nAs a result, delegators might choose to restake their tokens with another validator.\nThe longer the validators are jailed, the more likely is that delegators will restake.\nThus, by getting jailed, validators risk damaging their reputation."}),"\n"]}),"\n",(0,t.jsx)(n.p,{children:"Misbehaviors don't need to be malicious, e.g., most cases of double signing infractions are due to misconfiguration.\nThis means that, by opting in on multiple consumer chains, validators and their delegators incur a higher risk.\nAs a result, validators and their delegators want to be compensated for this additional risk, which makes the current design of ICS expensive."}),"\n",(0,t.jsx)(n.p,{children:"This ADR addresses the high cost of ICS by allowing consumer chains to customize the slashing and jailing conditions.\nBasically, every consumer chain can decide the punishment for every type of infraction.\nThis enables consumer chains to tradeoff economical security against cost."}),"\n",(0,t.jsx)(n.h2,{id:"decision",children:"Decision"}),"\n",(0,t.jsx)(n.p,{children:"To reduce the cost of ICS, consumer chains will be able to customize the slashing and jailing for every type of infraction.\nAs a result, consumer chains can decide on the amount of economic security they want and validators (and their delegators) can decide on the amount of additional risk they want to incur."}),"\n",(0,t.jsx)(n.p,{children:"For every consumer chain, we introduce the following slashing and jailing parameters:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-proto",children:'message InfractionParameters {\n  SlashJailParameters double_sign = 1;\n  SlashJailParameters downtime = 2;\n}\n\nmessage SlashJailParameters {\n  bytes slash_fraction = 1 [\n    (cosmos_proto.scalar)  = "cosmos.Dec",\n    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",\n    (gogoproto.nullable)   = false,\n    (amino.dont_omitempty) = true\n  ];\n  // use time.Unix(253402300799, 0) for permanent jailing\n  google.protobuf.Duration jail_duration = 2;\n}\n'})}),"\n",(0,t.jsx)(n.p,{children:"Currently, we consider only two infraction types -- double signing and downtime."}),"\n",(0,t.jsx)(n.p,{children:"By default, every consumer chain will have the following slashing and jailing parameters:"}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-go",children:"double_sign.slash_fraction: 0.05 // same as on the provider\ndouble_sign.jail_duration: time.Unix(253402300799, 0) // permanent jailing, same as on the provider\ndowntime.slash_fraction: 0 // no slashing for downtime on the consumer\ndowntime.jail_duration: 600s // same as on the provider\n"})}),"\n",(0,t.jsxs)(n.p,{children:["These parameters can be updated by the consumer chain owner at any given time (via ",(0,t.jsx)(n.code,{children:"MsgCreateConsumer"})," or ",(0,t.jsx)(n.code,{children:"MsgUpdateConsumer"}),").\nHowever, the changes will come into effect after a period equal to the staking module's unbonding period elapses.\nThis will allow validators that don't agree with the changes to opt out and not be affected by them.\nAlso, to avoid malicious chains attacking the provider validator set, these params will be bounded by the values on the provider chain:"]}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-go",children:"double_sign.slash_fraction <= 0.05 // 5%\ndowntime.slash_fraction <= 0.0001 // 0.1%\ndowntime.jail_duration <= 600s // 10 minutes\n"})}),"\n",(0,t.jsx)(n.p,{children:"Although consumer chains can set any values to these parameters (within the allowed bounds), we recommend the following settings, depending on the type of consumer chain."}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsxs)(n.li,{children:["\n",(0,t.jsxs)(n.p,{children:[(0,t.jsx)(n.strong,{children:"Proof-of-Stake (PoS) Consumer Chains."})," These are chains that have the full economical security of the provider validators that opted in. This means that all slashing and jailing parameters are the same as on the provider."]}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-go",children:"double_sign.slash_fraction: 0.05 \ndouble_sign.jail_duration: time.Unix(253402300799, 0)\ndowntime.slash_fraction: 0.0001 \ndowntime.jail_duration: 600s\n"})}),"\n"]}),"\n",(0,t.jsxs)(n.li,{children:["\n",(0,t.jsx)(n.p,{children:(0,t.jsx)(n.strong,{children:"Proof-of-Reputation (PoR) Consumer Chains."})}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-go",children:"double_sign.slash_fraction: 0 // no slashing\ndouble_sign.jail_duration: time.Unix(253402300799, 0)\ndowntime.slash_fraction: 0 // no slashing\ndowntime.jail_duration: 600s\n"})}),"\n",(0,t.jsx)(n.p,{children:"This means that when validators that opt in misbehave on PoR consumer chains, their stake on the provider is not being slashed, instead they are just jailed on the provider.\nAs a result, delegators incur (almost) no risk if their validators opt in on multiple PoR consumer chains.\nIf their validators are jailed, then the delegators can redelegate to other validators.\nNote though that delegators cannot redelegate multiple times, which means that if the new validators also get permanently jailed, the delegators need to wait for the unbonding period to elapse."}),"\n"]}),"\n",(0,t.jsxs)(n.li,{children:["\n",(0,t.jsx)(n.p,{children:(0,t.jsx)(n.strong,{children:"Testnet Consumer Chains."})}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-go",children:"double_sign.slash_fraction: 0 // no slashing\ndouble_sign.jail_duration: 0 // no jailing\ndowntime.slash_fraction: 0 // no slashing\ndowntime.jail_duration: 0 // no jailing\n"})}),"\n",(0,t.jsx)(n.p,{children:"This means that validators are not punished for infractions on consumer chains.\nThis setting is ideal for testing consumer chains before going in production, as neither validators nor their delegators incur any risk from the validators opting in on these consumer chains."}),"\n"]}),"\n"]}),"\n",(0,t.jsxs)(n.p,{children:["This means that both PoR and testnet consumer chains need only to cover the operational costs of the validators that opt in.\nFor example, if we take ",(0,t.jsx)(n.code,{children:"$600"})," as the cost of running a validator node, a budget of ",(0,t.jsx)(n.code,{children:"$3000"})," will be sufficient to cover the cost of four validators running such a consumer chain and have ",(0,t.jsx)(n.code,{children:"$150"})," profit per validator as incentive.\nIn practice, this can be implemented via the per-consumer-chain commission rate that allows validators to have different commission rates on different consumer chains."]}),"\n",(0,t.jsx)(n.h3,{id:"implementation",children:"Implementation"}),"\n",(0,t.jsx)(n.p,{children:"The implementation of this feature involves the following steps:"}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsxs)(n.li,{children:["Add the ",(0,t.jsx)(n.code,{children:"InfractionParameters"})," to ",(0,t.jsx)(n.code,{children:"MsgCreateConsumer"}),"."]}),"\n",(0,t.jsxs)(n.li,{children:["On slashing events (for either downtime or double signing infractions), use the corresponding ",(0,t.jsx)(n.code,{children:"slash_fraction"})," set by the consumer chain."]}),"\n",(0,t.jsxs)(n.li,{children:["On jailing events (for either downtime or double signing infractions), use the corresponding ",(0,t.jsx)(n.code,{children:"jail_duration"})," set by the consumer chain."]}),"\n",(0,t.jsx)(n.li,{children:"Cryptographic equivocation evidence received for PoR chains results in the misbehaving validators only being tombstoned and not slashed."}),"\n",(0,t.jsxs)(n.li,{children:["(Optional) Add the ",(0,t.jsx)(n.code,{children:"InfractionParameters"})," to ",(0,t.jsx)(n.code,{children:"MsgUpdateConsumer"}),", i.e., allow consumer chains to update the slashing and jailing parameters, but the changes will come into effect after a period equal to the staking module's unbonding period elapses to allow for validators to opt out."]}),"\n"]}),"\n",(0,t.jsx)(n.h2,{id:"consequences",children:"Consequences"}),"\n",(0,t.jsx)(n.h3,{id:"positive",children:"Positive"}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsx)(n.li,{children:"Reduce the cost of ICS by removing the risk of slashing delegators."}),"\n"]}),"\n",(0,t.jsx)(n.h3,{id:"negative",children:"Negative"}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsx)(n.li,{children:"Reduce the economical security of consumer chains with weaker slashing conditions."}),"\n"]}),"\n",(0,t.jsx)(n.h4,{id:"economic-security-model-without-slashing",children:"Economic Security Model without Slashing"}),"\n",(0,t.jsx)(n.p,{children:"The economic security model of most Cosmos chains relies on the following properties:"}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsx)(n.li,{children:"validators are not anonymous, which means that they could be legally liable if they are malicious;"}),"\n",(0,t.jsx)(n.li,{children:"the delegated PoS mechanism creates a reputation-based network of validators;"}),"\n",(0,t.jsx)(n.li,{children:"most validators have most of their stake coming from delegations (i.e., nothing at stake, besides reputation);"}),"\n",(0,t.jsx)(n.li,{children:"it is relatively difficult to enter the active validator set and even more so to climb the voting power ladder."}),"\n"]}),"\n",(0,t.jsx)(n.p,{children:"These properties enable us to make the following assumption:"}),"\n",(0,t.jsxs)(n.ul,{children:["\n",(0,t.jsx)(n.li,{children:"Being permanently removed from the provider validator set is strong enough of a deterrent to misbehaving on consumer chains."}),"\n"]}),"\n",(0,t.jsx)(n.p,{children:"The additional economical security a consumer gets from slashing is limited:\nSince most of the stake is delegated, slashing punishes delegators more than validators."}),"\n",(0,t.jsxs)(n.p,{children:["One benefit of slashing is that it acts as a deterrent for someone buying a large amount of staking tokens in order to attack a consumer chain.\nFor example, an attacker could get ",(0,t.jsx)(n.code,{children:"$15,000,000"})," worth of ATOM, which would give them around ",(0,t.jsx)(n.code,{children:"1%"})," voting power on the Cosmos Hub (at the time of this writing).\nOn a consumer chain, this voting power could be amplified depending on the other validators that opt in.\nHowever, by having the right ",(0,t.jsx)(n.a,{href:"https://cosmos.github.io/interchain-security/features/power-shaping",children:"power shaping"})," settings, the voting power of validators can be capped.\nThis means that even if the attacker manages to double sign without getting slashed, as long as they don't have 1/3+ of the voting power, they cannot benefit from the attack.\nMoreover, the attacker might lose due to other factors, such as ",(0,t.jsx)(n.a,{href:"https://forum.cosmos.network/t/enabling-opt-in-and-mesh-security-with-fraud-votes/10901",children:"token toxicity"}),"."]}),"\n",(0,t.jsx)(n.h3,{id:"neutral",children:"Neutral"}),"\n",(0,t.jsx)(n.p,{children:"NA"}),"\n",(0,t.jsx)(n.h2,{id:"references",children:"References"})]})}function c(e={}){const{wrapper:n}={...(0,o.a)(),...e.components};return n?(0,t.jsx)(n,{...e,children:(0,t.jsx)(d,{...e})}):d(e)}},1151:(e,n,i)=>{i.d(n,{Z:()=>r,a:()=>a});var t=i(7294);const o={},s=t.createContext(o);function a(e){const n=t.useContext(s);return t.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function r(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(o):e.components||o:a(e.components),t.createElement(s.Provider,{value:n},e.children)}}}]);