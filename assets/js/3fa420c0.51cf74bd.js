"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[144],{1493:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>u,contentTitle:()=>s,default:()=>d,frontMatter:()=>a,metadata:()=>l,toc:()=>i});var r=n(4848),o=n(8453);n(1470),n(9365);const a={},s="Query",l={id:"configuration/query",title:"Query",description:"Queries are defined as Protobuf RPCs annotated with the temporal.v1.query method option. They're mapped to workflows using the query workflow option. See the Queries guide for more usage details.",source:"@site/docs/configuration/query.mdx",sourceDirName:"configuration",slug:"/configuration/query",permalink:"/protoc-gen-go-temporal/docs/configuration/query",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/configuration/query.mdx",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"Activity",permalink:"/protoc-gen-go-temporal/docs/configuration/activity"},next:{title:"Signal",permalink:"/protoc-gen-go-temporal/docs/configuration/signal"}},u={},i=[{value:"Options",id:"options",level:2},{value:"name",id:"name",level:3},{value:"xns",id:"xns",level:3}];function c(e){const t={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",h3:"h3",p:"p",pre:"pre",...(0,o.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(t.h1,{id:"query",children:"Query"}),"\n",(0,r.jsxs)(t.p,{children:[(0,r.jsx)(t.a,{href:"https://docs.temporal.io/workflows#query",children:"Queries"})," are defined as Protobuf RPCs annotated with the ",(0,r.jsx)(t.code,{children:"temporal.v1.query"})," method option. They're mapped to workflows using the ",(0,r.jsx)(t.a,{href:"/docs/configuration/workflow#query",children:"query workflow option"}),". See the ",(0,r.jsx)(t.a,{href:"/docs/guides/queries",children:"Queries guide"})," for more usage details."]}),"\n",(0,r.jsx)(t.admonition,{type:"warning",children:(0,r.jsxs)(t.p,{children:["Query definitions must specify a non-empty output parameter. Query definitions can omit an input parameter by specifying the native ",(0,r.jsx)(t.code,{children:"google.protobuf.Empty"})," message type in its place. This requires an additional ",(0,r.jsx)(t.code,{children:"google/protobuf/empty.proto"})," protobuf import."]})}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  // Hello returns a friendly greeting\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      query: { ref: \'GetHelloStatus\' }\n    };\n  }\n\n  // GetHelloStatus retrieves the status of an existing Hello workflow\n  rpc GetHelloStatus(GetHelloStatusInput) returns (GetHelloStatusOutput) {\n    option (temporal.v1.query) = {};\n  }\n}\n'})}),"\n",(0,r.jsx)(t.h2,{id:"options",children:"Options"}),"\n",(0,r.jsx)(t.h3,{id:"name",children:"name"}),"\n",(0,r.jsx)(t.p,{children:(0,r.jsx)(t.code,{children:"string"})}),"\n",(0,r.jsxs)(t.p,{children:["Fully qualified ",(0,r.jsx)(t.a,{href:"https://docs.temporal.io/workflows#query",children:"Query type name"}),". Defaults to protobuf method full name (e.g. ",(0,r.jsx)(t.code,{children:"example.v1.Example.GetHelloStatus"}),")"]}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-protobuf",children:'service Example {\n  rpc GetHelloStatus(GetHelloStatusInput) returns (GetHelloStatusOutput) {\n    option (temporal.v1.query) = {\n      name: "GetHelloStatus"\n    };\n  }\n}\n'})}),"\n",(0,r.jsx)(t.h3,{id:"xns",children:"xns"}),"\n",(0,r.jsx)(t.p,{children:(0,r.jsx)(t.a,{href:"https://buf.build/cludden/protoc-gen-go-temporal/docs/main:temporal.v1#temporal.v1.XNSActivityOptions",children:"temporal.v1.XNSActivityOptions"})}),"\n",(0,r.jsxs)(t.p,{children:["Used to configure ",(0,r.jsx)(t.a,{href:"/docs/guides/xns",children:"cross-namespace"})," activity options."]}),"\n",(0,r.jsx)(t.admonition,{type:"note",children:(0,r.jsxs)(t.p,{children:["This requires the ",(0,r.jsx)(t.a,{href:"/docs/configuration/plugin",children:"enable-xns"})," plugin option to be enabled."]})}),"\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-protobuf",children:"service Example {\n  rpc GetHelloStatus(GetHelloStatusInput) returns (GetHelloStatusOutput) {\n    option (temporal.v1.query) = {\n      xns: {\n        heartbeat_timeout: { seconds: 30 }\n        heartbeat_interval: { seconds: 10 }\n        start_to_close_timeout: { seconds: 300 }\n      }\n    };\n  }\n}\n"})})]})}function d(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,r.jsx)(t,{...e,children:(0,r.jsx)(c,{...e})}):c(e)}},9365:(e,t,n)=>{n.d(t,{A:()=>s});n(6540);var r=n(4164);const o={tabItem:"tabItem_Ymn6"};var a=n(4848);function s(e){let{children:t,hidden:n,className:s}=e;return(0,a.jsx)("div",{role:"tabpanel",className:(0,r.A)(o.tabItem,s),hidden:n,children:t})}},1470:(e,t,n)=>{n.d(t,{A:()=>w});var r=n(6540),o=n(4164),a=n(3104),s=n(6347),l=n(205),u=n(7485),i=n(1682),c=n(9466);function d(e){return r.Children.toArray(e).filter((e=>"\n"!==e)).map((e=>{if(!e||(0,r.isValidElement)(e)&&function(e){const{props:t}=e;return!!t&&"object"==typeof t&&"value"in t}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))?.filter(Boolean)??[]}function p(e){const{values:t,children:n}=e;return(0,r.useMemo)((()=>{const e=t??function(e){return d(e).map((e=>{let{props:{value:t,label:n,attributes:r,default:o}}=e;return{value:t,label:n,attributes:r,default:o}}))}(n);return function(e){const t=(0,i.X)(e,((e,t)=>e.value===t.value));if(t.length>0)throw new Error(`Docusaurus error: Duplicate values "${t.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[t,n])}function h(e){let{value:t,tabValues:n}=e;return n.some((e=>e.value===t))}function m(e){let{queryString:t=!1,groupId:n}=e;const o=(0,s.W6)(),a=function(e){let{queryString:t=!1,groupId:n}=e;if("string"==typeof t)return t;if(!1===t)return null;if(!0===t&&!n)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return n??null}({queryString:t,groupId:n});return[(0,u.aZ)(a),(0,r.useCallback)((e=>{if(!a)return;const t=new URLSearchParams(o.location.search);t.set(a,e),o.replace({...o.location,search:t.toString()})}),[a,o])]}function f(e){const{defaultValue:t,queryString:n=!1,groupId:o}=e,a=p(e),[s,u]=(0,r.useState)((()=>function(e){let{defaultValue:t,tabValues:n}=e;if(0===n.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(t){if(!h({value:t,tabValues:n}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${t}" but none of its children has the corresponding value. Available values are: ${n.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return t}const r=n.find((e=>e.default))??n[0];if(!r)throw new Error("Unexpected error: 0 tabValues");return r.value}({defaultValue:t,tabValues:a}))),[i,d]=m({queryString:n,groupId:o}),[f,g]=function(e){let{groupId:t}=e;const n=function(e){return e?`docusaurus.tab.${e}`:null}(t),[o,a]=(0,c.Dv)(n);return[o,(0,r.useCallback)((e=>{n&&a.set(e)}),[n,a])]}({groupId:o}),b=(()=>{const e=i??f;return h({value:e,tabValues:a})?e:null})();(0,l.A)((()=>{b&&u(b)}),[b]);return{selectedValue:s,selectValue:(0,r.useCallback)((e=>{if(!h({value:e,tabValues:a}))throw new Error(`Can't select invalid tab value=${e}`);u(e),d(e),g(e)}),[d,g,a]),tabValues:a}}var g=n(2303);const b={tabList:"tabList__CuJ",tabItem:"tabItem_LNqP"};var v=n(4848);function x(e){let{className:t,block:n,selectedValue:r,selectValue:s,tabValues:l}=e;const u=[],{blockElementScrollPositionUntilNextRender:i}=(0,a.a_)(),c=e=>{const t=e.currentTarget,n=u.indexOf(t),o=l[n].value;o!==r&&(i(t),s(o))},d=e=>{let t=null;switch(e.key){case"Enter":c(e);break;case"ArrowRight":{const n=u.indexOf(e.currentTarget)+1;t=u[n]??u[0];break}case"ArrowLeft":{const n=u.indexOf(e.currentTarget)-1;t=u[n]??u[u.length-1];break}}t?.focus()};return(0,v.jsx)("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,o.A)("tabs",{"tabs--block":n},t),children:l.map((e=>{let{value:t,label:n,attributes:a}=e;return(0,v.jsx)("li",{role:"tab",tabIndex:r===t?0:-1,"aria-selected":r===t,ref:e=>u.push(e),onKeyDown:d,onClick:c,...a,className:(0,o.A)("tabs__item",b.tabItem,a?.className,{"tabs__item--active":r===t}),children:n??t},t)}))})}function y(e){let{lazy:t,children:n,selectedValue:o}=e;const a=(Array.isArray(n)?n:[n]).filter(Boolean);if(t){const e=a.find((e=>e.props.value===o));return e?(0,r.cloneElement)(e,{className:"margin-top--md"}):null}return(0,v.jsx)("div",{className:"margin-top--md",children:a.map(((e,t)=>(0,r.cloneElement)(e,{key:t,hidden:e.props.value!==o})))})}function j(e){const t=f(e);return(0,v.jsxs)("div",{className:(0,o.A)("tabs-container",b.tabList),children:[(0,v.jsx)(x,{...e,...t}),(0,v.jsx)(y,{...e,...t})]})}function w(e){const t=(0,g.A)();return(0,v.jsx)(j,{...e,children:d(e.children)},String(t))}},8453:(e,t,n)=>{n.d(t,{R:()=>s,x:()=>l});var r=n(6540);const o={},a=r.createContext(o);function s(e){const t=r.useContext(a);return r.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function l(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(o):e.components||o:s(e.components),r.createElement(a.Provider,{value:t},e.children)}}}]);