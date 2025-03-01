"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[2761],{46563:(e,i,n)=>{n.r(i),n.d(i,{assets:()=>d,contentTitle:()=>a,default:()=>h,frontMatter:()=>s,metadata:()=>c,toc:()=>o});var r=n(85893),t=n(11151);const s={sidebar_position:1,title:"Overview"},a="Overview",c={id:"adrs/intro",title:"Overview",description:"This is a location to record all high-level architecture decisions in the Interchain Security project.",source:"@site/versioned_docs/version-v6.4.0/adrs/intro.md",sourceDirName:"adrs",slug:"/adrs/intro",permalink:"/interchain-security/v6.4.0/adrs/intro",draft:!1,unlisted:!1,tags:[],version:"v6.4.0",sidebarPosition:1,frontMatter:{sidebar_position:1,title:"Overview"},sidebar:"tutorialSidebar",previous:{title:"Inactive Validators Integration Guide",permalink:"/interchain-security/v6.4.0/integrators/integrating_inactive_validators"},next:{title:"Denom DOS fixes",permalink:"/interchain-security/v6.4.0/adrs/adr-004-denom-dos-fixes"}},d={},o=[{value:"Table of Contents",id:"table-of-contents",level:2},{value:"Accepted",id:"accepted",level:3},{value:"Proposed",id:"proposed",level:3},{value:"Rejected",id:"rejected",level:3},{value:"Deprecated",id:"deprecated",level:3}];function l(e){const i={a:"a",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",ul:"ul",...(0,t.a)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(i.h1,{id:"overview",children:"Overview"}),"\n",(0,r.jsx)(i.p,{children:"This is a location to record all high-level architecture decisions in the Interchain Security project."}),"\n",(0,r.jsxs)(i.p,{children:["You can read more about the Architecture Decision Record (ADR) concept in this ",(0,r.jsx)(i.a,{href:"https://product.reverb.com/documenting-architecture-decisions-the-reverb-way-a3563bb24bd0#.78xhdix6t",children:"blog post"}),"."]}),"\n",(0,r.jsx)(i.p,{children:"An ADR should provide:"}),"\n",(0,r.jsxs)(i.ul,{children:["\n",(0,r.jsx)(i.li,{children:"Context on the relevant goals and the current state"}),"\n",(0,r.jsx)(i.li,{children:"Proposed changes to achieve the goals"}),"\n",(0,r.jsx)(i.li,{children:"Summary of pros and cons"}),"\n",(0,r.jsx)(i.li,{children:"References"}),"\n",(0,r.jsx)(i.li,{children:"Changelog"}),"\n"]}),"\n",(0,r.jsx)(i.p,{children:"Note the distinction between an ADR and a spec. The ADR provides the context, intuition, reasoning, and\njustification for a change in architecture, or for the architecture of something\nnew. The spec is much more compressed and streamlined summary of everything as\nit is or should be."}),"\n",(0,r.jsx)(i.p,{children:"If recorded decisions turned out to be lacking, convene a discussion, record the new decisions here, and then modify the code to match."}),"\n",(0,r.jsx)(i.p,{children:"Note the context/background should be written in the present tense."}),"\n",(0,r.jsxs)(i.p,{children:["To suggest an ADR, please make use of the ",(0,r.jsx)(i.a,{href:"https://github.com/cosmos/interchain-security/blob/main/docs/docs/adrs/templates/adr-template.md",children:"ADR template"})," provided."]}),"\n",(0,r.jsx)(i.h2,{id:"table-of-contents",children:"Table of Contents"}),"\n",(0,r.jsx)(i.h3,{id:"accepted",children:"Accepted"}),"\n",(0,r.jsxs)(i.ul,{children:["\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-001-key-assignment",children:"ADR 001: Key Assignment"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-002-throttle",children:"ADR 002: Jail Throttling"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-004-denom-dos-fixes",children:"ADR 004: Denom DOS fixes"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-005-cryptographic-equivocation-verification",children:"ADR 005: Cryptographic verification of equivocation evidence"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-008-throttle-retries",children:"ADR 008: Throttle with retries"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-010-standalone-changeover",children:"ADR 010: Standalone to Consumer Changeover"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-013-equivocation-slashing",children:"ADR 013: Slashing on the provider for consumer equivocation"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-014-epochs",children:"ADR 014: Epochs"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-015-partial-set-security",children:"ADR 015: Partial Set Security"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-017-allowing-inactive-validators",children:"ADR 017: ICS with Inactive Provider Validators"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-018-remove-vscmatured",children:"ADR 018: Remove VSCMatured Packets"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-019-permissionless-ics",children:"ADR 019: Permissionless Interchain Security"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-020-cutomizable_slashing_and_jailing",children:"ADR 020: Customizable Slashing and Jailing"})}),"\n"]}),"\n",(0,r.jsx)(i.h3,{id:"proposed",children:"Proposed"}),"\n",(0,r.jsxs)(i.ul,{children:["\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-011-improving-test-confidence",children:"ADR 011: Improving testing and increasing confidence"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-016-securityaggregation",children:"ADR 016: Security aggregation"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-021-consumer-chain-clients",children:"ADR 021: Consumer Chain Clients"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-022-fault-resolutions",children:"ADR 022: Fault Resolutions"})}),"\n"]}),"\n",(0,r.jsx)(i.h3,{id:"rejected",children:"Rejected"}),"\n",(0,r.jsxs)(i.ul,{children:["\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-007-pause-unbonding-on-eqv-prop",children:"ADR 007: Pause validator unbonding during equivocation proposal"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-012-separate-releasing",children:"ADR 012: Separate Releasing"})}),"\n"]}),"\n",(0,r.jsx)(i.h3,{id:"deprecated",children:"Deprecated"}),"\n",(0,r.jsxs)(i.ul,{children:["\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-003-equivocation-gov-proposal",children:"ADR 003: Equivocation governance proposal"})}),"\n",(0,r.jsx)(i.li,{children:(0,r.jsx)(i.a,{href:"/interchain-security/v6.4.0/adrs/adr-009-soft-opt-out",children:"ADR 009: Soft Opt-Out"})}),"\n"]})]})}function h(e={}){const{wrapper:i}={...(0,t.a)(),...e.components};return i?(0,r.jsx)(i,{...e,children:(0,r.jsx)(l,{...e})}):l(e)}},11151:(e,i,n)=>{n.d(i,{Z:()=>c,a:()=>a});var r=n(67294);const t={},s=r.createContext(t);function a(e){const i=r.useContext(s);return r.useMemo((function(){return"function"==typeof e?e(i):{...i,...e}}),[i,e])}function c(e){let i;return i=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:a(e.components),r.createElement(s.Provider,{value:i},e.children)}}}]);