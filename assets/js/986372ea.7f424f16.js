"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[2896],{12869:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>s,default:()=>d,frontMatter:()=>a,metadata:()=>i,toc:()=>m});var o=t(85893),r=t(11151);const a={sidebar_position:3},s="Consumer Chain Governance",i={id:"consumer-development/consumer-chain-governance",title:"Consumer Chain Governance",description:"Different consumer chains can do governance in different ways.",source:"@site/versioned_docs/version-v6.4.0/consumer-development/consumer-chain-governance.md",sourceDirName:"consumer-development",slug:"/consumer-development/consumer-chain-governance",permalink:"/interchain-security/v6.4.0/consumer-development/consumer-chain-governance",draft:!1,unlisted:!1,tags:[],version:"v6.4.0",sidebarPosition:3,frontMatter:{sidebar_position:3},sidebar:"tutorialSidebar",previous:{title:"Developing an ICS consumer chain",permalink:"/interchain-security/v6.4.0/consumer-development/app-integration"},next:{title:"Onboarding Checklist",permalink:"/interchain-security/v6.4.0/consumer-development/onboarding"}},c={},m=[{value:"Democracy module",id:"democracy-module",level:2},{value:"CosmWasm",id:"cosmwasm",level:2}];function h(e){const n={a:"a",h1:"h1",h2:"h2",p:"p",...(0,r.a)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(n.h1,{id:"consumer-chain-governance",children:"Consumer Chain Governance"}),"\n",(0,o.jsx)(n.p,{children:"Different consumer chains can do governance in different ways."}),"\n",(0,o.jsx)(n.h2,{id:"democracy-module",children:"Democracy module"}),"\n",(0,o.jsx)(n.p,{children:"The democracy module provides a governance experience identical to what exists on a standalone Cosmos chain, with one small but important difference. On a standalone Cosmos chain validators can act as representatives for their delegators by voting with their stake, but only if the delegator themselves does not vote. This is a lightweight form of liquid democracy."}),"\n",(0,o.jsx)(n.p,{children:"Using the democracy module on a consumer chain is the exact same experience, except for the fact that it is not the actual validator set of the chain (since it is a consumer chain, these are the Cosmos Hub validators) acting as representatives. Instead, there is a separate representative role who token holders can delegate to and who can perform the functions that validators do in Cosmos governance, without participating in proof of stake consensus."}),"\n",(0,o.jsxs)(n.p,{children:["For an example, see the ",(0,o.jsx)(n.a,{href:"https://github.com/cosmos/interchain-security/tree/main/app/consumer-democracy",children:"Democracy Consumer"})]}),"\n",(0,o.jsx)(n.h2,{id:"cosmwasm",children:"CosmWasm"}),"\n",(0,o.jsx)(n.p,{children:"There are several great DAO and governance frameworks written as CosmWasm contracts. These can be used as the main governance system for a consumer chain. Actions triggered by the CosmWasm governance contracts are able to affect parameters and trigger actions on the consumer chain."}),"\n",(0,o.jsxs)(n.p,{children:["For an example, see ",(0,o.jsx)(n.a,{href:"https://github.com/neutron-org/neutron/",children:"Neutron"}),"."]})]})}function d(e={}){const{wrapper:n}={...(0,r.a)(),...e.components};return n?(0,o.jsx)(n,{...e,children:(0,o.jsx)(h,{...e})}):h(e)}},11151:(e,n,t)=>{t.d(n,{Z:()=>i,a:()=>s});var o=t(67294);const r={},a=o.createContext(r);function s(e){const n=o.useContext(a);return o.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function i(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:s(e.components),o.createElement(a.Provider,{value:n},e.children)}}}]);