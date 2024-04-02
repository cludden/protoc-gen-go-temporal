"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[5995],{528:(e,n,o)=>{o.r(n),o.d(n,{assets:()=>a,contentTitle:()=>s,default:()=>u,frontMatter:()=>i,metadata:()=>c,toc:()=>l});var t=o(4848),r=o(8453);const i={},s="Service",c={id:"configuration/service",title:"Service",description:"Service options apply to all Temporal resources defined within an individual protobuf service.",source:"@site/docs/configuration/service.mdx",sourceDirName:"configuration",slug:"/configuration/service",permalink:"/protoc-gen-go-temporal/docs/configuration/service",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/configuration/service.mdx",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"Plugin",permalink:"/protoc-gen-go-temporal/docs/configuration/plugin"},next:{title:"Workflow",permalink:"/protoc-gen-go-temporal/docs/configuration/workflow"}},a={},l=[{value:"Options",id:"options",level:2},{value:"task_queue",id:"task_queue",level:3}];function p(e){const n={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",p:"p",pre:"pre",...(0,r.R)(),...e.components};return(0,t.jsxs)(t.Fragment,{children:[(0,t.jsx)(n.h1,{id:"service",children:"Service"}),"\n",(0,t.jsx)(n.p,{children:"Service options apply to all Temporal resources defined within an individual protobuf service."}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  };\n}\n'})}),"\n",(0,t.jsx)(n.h2,{id:"options",children:"Options"}),"\n",(0,t.jsx)(n.h3,{id:"task_queue",children:"task_queue"}),"\n",(0,t.jsx)(n.p,{children:(0,t.jsx)(n.code,{children:"string"})}),"\n",(0,t.jsxs)(n.p,{children:["Specifies the default ",(0,t.jsx)(n.a,{href:"https://docs.temporal.io/workers#task-queue",children:"Task Queue"})," name."]}),"\n",(0,t.jsx)(n.pre,{children:(0,t.jsx)(n.code,{className:"language-protobuf",children:"service Example {\n  option (temporal.v1.service) = {\n    task_queue: 'example-v1'\n  };\n}\n"})})]})}function u(e={}){const{wrapper:n}={...(0,r.R)(),...e.components};return n?(0,t.jsx)(n,{...e,children:(0,t.jsx)(p,{...e})}):p(e)}},8453:(e,n,o)=>{o.d(n,{R:()=>s,x:()=>c});var t=o(6540);const r={},i=t.createContext(r);function s(e){const n=t.useContext(i);return t.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function c(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:s(e.components),t.createElement(i.Provider,{value:n},e.children)}}}]);