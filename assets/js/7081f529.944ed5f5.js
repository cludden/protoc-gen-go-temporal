"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[5995],{528:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>a,contentTitle:()=>s,default:()=>p,frontMatter:()=>i,metadata:()=>c,toc:()=>u});var o=t(4848),r=t(8453);const i={},s="Service",c={id:"configuration/service",title:"Service",description:"Service options apply to all Temporal resources defined within an individual protobuf service.",source:"@site/docs/configuration/service.mdx",sourceDirName:"configuration",slug:"/configuration/service",permalink:"/docs/configuration/service",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/configuration/service.mdx",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"Plugin",permalink:"/docs/configuration/plugin"},next:{title:"Workflow",permalink:"/docs/configuration/workflow"}},a={},u=[{value:"Options",id:"options",level:2},{value:"task_queue",id:"task_queue",level:3}];function l(e){const n={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",p:"p",pre:"pre",...(0,r.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(n.h1,{id:"service",children:"Service"}),"\n",(0,o.jsx)(n.p,{children:"Service options apply to all Temporal resources defined within an individual protobuf service."}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  };\n}\n'})}),"\n",(0,o.jsx)(n.h2,{id:"options",children:"Options"}),"\n",(0,o.jsx)(n.h3,{id:"task_queue",children:"task_queue"}),"\n",(0,o.jsx)(n.p,{children:(0,o.jsx)(n.code,{children:"string"})}),"\n",(0,o.jsxs)(n.p,{children:["Specifies the default ",(0,o.jsx)(n.a,{href:"https://docs.temporal.io/workers#task-queue",children:"Task Queue"})," name."]}),"\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-protobuf",children:"service Example {\n  option (temporal.v1.service) = {\n    task_queue: 'example-v1'\n  };\n}\n"})})]})}function p(e={}){const{wrapper:n}={...(0,r.R)(),...e.components};return n?(0,o.jsx)(n,{...e,children:(0,o.jsx)(l,{...e})}):l(e)}},8453:(e,n,t)=>{t.d(n,{R:()=>s,x:()=>c});var o=t(6540);const r={},i=o.createContext(r);function s(e){const n=o.useContext(i);return o.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function c(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:s(e.components),o.createElement(i.Provider,{value:n},e.children)}}}]);