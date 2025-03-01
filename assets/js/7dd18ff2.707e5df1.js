"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[8106],{37423:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>c,contentTitle:()=>o,default:()=>l,frontMatter:()=>r,metadata:()=>s,toc:()=>h});var i=a(85893),n=a(11151);const r={sidebar_position:5},o="Partial Set Security",s={id:"features/partial-set-security",title:"Partial Set Security",description:"Partial Set Security (PSS) is a complete revamp of the Hub's security offering. It allows consumer chains to leverage only a subset of validators from the provider chain, which offers more flexibility than the previous Replicated Security model. By introducing the top_N parameter, each consumer chain can choose the extent of security needed:",source:"@site/versioned_docs/version-v6.4.0/features/partial-set-security.md",sourceDirName:"features",slug:"/features/partial-set-security",permalink:"/interchain-security/v6.4.0/features/partial-set-security",draft:!1,unlisted:!1,tags:[],version:"v6.4.0",sidebarPosition:5,frontMatter:{sidebar_position:5},sidebar:"tutorialSidebar",previous:{title:"Consumer Initiated Slashing",permalink:"/interchain-security/v6.4.0/features/slashing"},next:{title:"Power Shaping",permalink:"/interchain-security/v6.4.0/features/power-shaping"}},c={},h=[];function d(e){const t={a:"a",code:"code",em:"em",h1:"h1",li:"li",p:"p",strong:"strong",ul:"ul",...(0,n.a)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(t.h1,{id:"partial-set-security",children:"Partial Set Security"}),"\n",(0,i.jsxs)(t.p,{children:["Partial Set Security (PSS) is a complete revamp of the Hub's security offering. It allows consumer chains to leverage only a subset of validators from the provider chain, which offers more flexibility than the previous Replicated Security model. By introducing the ",(0,i.jsx)(t.code,{children:"top_N"})," parameter, each consumer chain can choose the extent of security needed:"]}),"\n",(0,i.jsxs)(t.ul,{children:["\n",(0,i.jsxs)(t.li,{children:["\n",(0,i.jsx)(t.p,{children:"Top N: Requires the top N% validators from the provider chain to secure the consumer chain. This guarantees that the validators with the most power on the provider will validate the consumer chain, while others can voluntarily opt in."}),"\n"]}),"\n",(0,i.jsxs)(t.li,{children:["\n",(0,i.jsxs)(t.p,{children:["Opt-In: If the ",(0,i.jsx)(t.code,{children:"top_N"})," parameter is set to zero, no validator is mandated to secure the consumer chain. Instead, any validator from the provider chain can opt in using a dedicated transaction."]}),"\n"]}),"\n"]}),"\n",(0,i.jsx)(t.p,{children:"An advantage of Top N chains is that they are guaranteed to receive at least a certain fraction of the market cap of the provider chain in security.\nIn turn, Top N chains need to be approved by governance, since some validators will be forced to validate on them.\nThus, Top N chains should typically expect to need to provide a strong case for why they should be added to the provider chain, and they should make sure they offer enough rewards to incentivize validators and delegators to vote for their proposal."}),"\n",(0,i.jsxs)(t.p,{children:["Opt-In chains, on the other hand, are more flexible.\nValidators are never forced to validate these chains and simply opt in if they want to.\nBecause of this, Opt-In chains can be ",(0,i.jsx)(t.strong,{children:(0,i.jsx)(t.em,{children:"launch completely permissionlessly"})})," by sending a transaction to the provider chain.\nAs a trade-off though, Opt-In chains do not get a fixed amount of security as a relation of the market cap of the provider as top N chains do, so Opt-In chains might want to keep an eye on how many validators have opted in to validate their chain and adjust their reward emissions accordingly to incentivize validators."]}),"\n",(0,i.jsxs)(t.p,{children:["Note that Top N consumer chains can become Opt-In chains or vice versa via a ",(0,i.jsx)(t.a,{href:"/interchain-security/v6.4.0/features/permissionless",children:(0,i.jsx)(t.code,{children:"MsgUpdateConsumer"})})," message."]}),"\n",(0,i.jsx)(t.p,{children:"Partial Set Security is handled only by the provider chain - the consumer chains are simply sent validator sets, and they are not aware that this represents only a subset of the provider chain's validator set."})]})}function l(e={}){const{wrapper:t}={...(0,n.a)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},11151:(e,t,a)=>{a.d(t,{Z:()=>s,a:()=>o});var i=a(67294);const n={},r=i.createContext(n);function o(e){const t=i.useContext(r);return i.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function s(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(n):e.components||n:o(e.components),i.createElement(r.Provider,{value:t},e.children)}}}]);