"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[80],{6377:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>i,contentTitle:()=>s,default:()=>d,frontMatter:()=>a,metadata:()=>l,toc:()=>u});var r=t(4848),o=t(8453);t(1470),t(9365);const a={},s="Signal",l={id:"configuration/signal",title:"Signal",description:"Signals are defined as Protobuf RPCs annotated with the temporal.v1.signal method option. They're mapped to workflows using the signal workflow option. See the Signals guide for more usage details.",source:"@site/docs/configuration/signal.mdx",sourceDirName:"configuration",slug:"/configuration/signal",permalink:"/protoc-gen-go-temporal/docs/configuration/signal",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/configuration/signal.mdx",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"Query",permalink:"/protoc-gen-go-temporal/docs/configuration/query"},next:{title:"Update",permalink:"/protoc-gen-go-temporal/docs/configuration/update"}},i={},u=[{value:"Options",id:"options",level:2},{value:"name",id:"name",level:3},{value:"xns",id:"xns",level:3}];function c(e){const n={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",h3:"h3",p:"p",pre:"pre",...(0,o.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(n.h1,{id:"signal",children:"Signal"}),"\n",(0,r.jsxs)(n.p,{children:[(0,r.jsx)(n.a,{href:"https://docs.temporal.io/workflows#signal",children:"Signals"})," are defined as Protobuf RPCs annotated with the ",(0,r.jsx)(n.code,{children:"temporal.v1.signal"})," method option. They're mapped to workflows using the ",(0,r.jsx)(n.a,{href:"/docs/configuration/workflow#signal",children:"signal workflow option"}),". See the ",(0,r.jsx)(n.a,{href:"/docs/guides/signals",children:"Signals guide"})," for more usage details."]}),"\n",(0,r.jsx)(n.admonition,{type:"warning",children:(0,r.jsxs)(n.p,{children:["Signals definitions must use ",(0,r.jsx)(n.code,{children:"google.protobuf.Empty"})," as their return value. This requires an additional ",(0,r.jsx)(n.code,{children:"google/protobuf/empty.proto"})," protobuf import."]})}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nimport "google/protobuf/empty.proto";\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  // Hello returns a friendly greeting\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      signal: { ref: \'Ping\' }\n    };\n  }\n\n  // Ping sends a signal to an existing workflow\n  rpc Ping(PingInput) returns (google.protobuf.Empty) {\n    option (temporal.v1.signal) = {};\n  }\n}\n'})}),"\n",(0,r.jsx)(n.h2,{id:"options",children:"Options"}),"\n",(0,r.jsx)(n.h3,{id:"name",children:"name"}),"\n",(0,r.jsx)(n.p,{children:(0,r.jsx)(n.code,{children:"string"})}),"\n",(0,r.jsxs)(n.p,{children:["Fully qualified ",(0,r.jsx)(n.a,{href:"https://docs.temporal.io/workflows#signal",children:"Signal type name"}),". Defaults to protobuf method full name (e.g. ",(0,r.jsx)(n.code,{children:"example.v1.Example.Ping"}),")"]}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",children:'service Example {\n  rpc Ping(PingInput) returns (google.protobuf.Empty) {\n    option (temporal.v1.signal) = {\n      name: "Ping"\n    };\n  }\n}\n'})}),"\n",(0,r.jsx)(n.h3,{id:"xns",children:"xns"}),"\n",(0,r.jsx)(n.p,{children:(0,r.jsx)(n.a,{href:"https://buf.build/cludden/protoc-gen-go-temporal/docs/main:temporal.v1#temporal.v1.XNSActivityOptions",children:"temporal.v1.XNSActivityOptions"})}),"\n",(0,r.jsxs)(n.p,{children:["Used to configure ",(0,r.jsx)(n.a,{href:"/docs/guides/xns",children:"cross-namespace"})," activity options."]}),"\n",(0,r.jsx)(n.admonition,{type:"note",children:(0,r.jsxs)(n.p,{children:["This requires the ",(0,r.jsx)(n.a,{href:"/docs/configuration/plugin",children:"enable-xns"})," plugin option to be enabled."]})}),"\n",(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",children:"service Example {\n  rpc Ping(PingInput) returns (google.protobuf.Empty) {\n    option (temporal.v1.signal) = {\n      xns: {\n        heartbeat_timeout: { seconds: 30 }\n        heartbeat_interval: { seconds: 10 }\n        start_to_close_timeout: { seconds: 300 }\n      }\n    };\n  }\n}\n"})})]})}function d(e={}){const{wrapper:n}={...(0,o.R)(),...e.components};return n?(0,r.jsx)(n,{...e,children:(0,r.jsx)(c,{...e})}):c(e)}},9365:(e,n,t)=>{t.d(n,{A:()=>s});t(6540);var r=t(4164);const o={tabItem:"tabItem_Ymn6"};var a=t(4848);function s(e){let{children:n,hidden:t,className:s}=e;return(0,a.jsx)("div",{role:"tabpanel",className:(0,r.A)(o.tabItem,s),hidden:t,children:n})}},1470:(e,n,t)=>{t.d(n,{A:()=>w});var r=t(6540),o=t(4164),a=t(3104),s=t(6347),l=t(205),i=t(7485),u=t(1682),c=t(9466);function d(e){return r.Children.toArray(e).filter((e=>"\n"!==e)).map((e=>{if(!e||(0,r.isValidElement)(e)&&function(e){const{props:n}=e;return!!n&&"object"==typeof n&&"value"in n}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))?.filter(Boolean)??[]}function p(e){const{values:n,children:t}=e;return(0,r.useMemo)((()=>{const e=n??function(e){return d(e).map((e=>{let{props:{value:n,label:t,attributes:r,default:o}}=e;return{value:n,label:t,attributes:r,default:o}}))}(t);return function(e){const n=(0,u.X)(e,((e,n)=>e.value===n.value));if(n.length>0)throw new Error(`Docusaurus error: Duplicate values "${n.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[n,t])}function h(e){let{value:n,tabValues:t}=e;return t.some((e=>e.value===n))}function g(e){let{queryString:n=!1,groupId:t}=e;const o=(0,s.W6)(),a=function(e){let{queryString:n=!1,groupId:t}=e;if("string"==typeof n)return n;if(!1===n)return null;if(!0===n&&!t)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return t??null}({queryString:n,groupId:t});return[(0,i.aZ)(a),(0,r.useCallback)((e=>{if(!a)return;const n=new URLSearchParams(o.location.search);n.set(a,e),o.replace({...o.location,search:n.toString()})}),[a,o])]}function m(e){const{defaultValue:n,queryString:t=!1,groupId:o}=e,a=p(e),[s,i]=(0,r.useState)((()=>function(e){let{defaultValue:n,tabValues:t}=e;if(0===t.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(n){if(!h({value:n,tabValues:t}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${n}" but none of its children has the corresponding value. Available values are: ${t.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return n}const r=t.find((e=>e.default))??t[0];if(!r)throw new Error("Unexpected error: 0 tabValues");return r.value}({defaultValue:n,tabValues:a}))),[u,d]=g({queryString:t,groupId:o}),[m,f]=function(e){let{groupId:n}=e;const t=function(e){return e?`docusaurus.tab.${e}`:null}(n),[o,a]=(0,c.Dv)(t);return[o,(0,r.useCallback)((e=>{t&&a.set(e)}),[t,a])]}({groupId:o}),b=(()=>{const e=u??m;return h({value:e,tabValues:a})?e:null})();(0,l.A)((()=>{b&&i(b)}),[b]);return{selectedValue:s,selectValue:(0,r.useCallback)((e=>{if(!h({value:e,tabValues:a}))throw new Error(`Can't select invalid tab value=${e}`);i(e),d(e),f(e)}),[d,f,a]),tabValues:a}}var f=t(2303);const b={tabList:"tabList__CuJ",tabItem:"tabItem_LNqP"};var v=t(4848);function x(e){let{className:n,block:t,selectedValue:r,selectValue:s,tabValues:l}=e;const i=[],{blockElementScrollPositionUntilNextRender:u}=(0,a.a_)(),c=e=>{const n=e.currentTarget,t=i.indexOf(n),o=l[t].value;o!==r&&(u(n),s(o))},d=e=>{let n=null;switch(e.key){case"Enter":c(e);break;case"ArrowRight":{const t=i.indexOf(e.currentTarget)+1;n=i[t]??i[0];break}case"ArrowLeft":{const t=i.indexOf(e.currentTarget)-1;n=i[t]??i[i.length-1];break}}n?.focus()};return(0,v.jsx)("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,o.A)("tabs",{"tabs--block":t},n),children:l.map((e=>{let{value:n,label:t,attributes:a}=e;return(0,v.jsx)("li",{role:"tab",tabIndex:r===n?0:-1,"aria-selected":r===n,ref:e=>i.push(e),onKeyDown:d,onClick:c,...a,className:(0,o.A)("tabs__item",b.tabItem,a?.className,{"tabs__item--active":r===n}),children:t??n},n)}))})}function j(e){let{lazy:n,children:t,selectedValue:o}=e;const a=(Array.isArray(t)?t:[t]).filter(Boolean);if(n){const e=a.find((e=>e.props.value===o));return e?(0,r.cloneElement)(e,{className:"margin-top--md"}):null}return(0,v.jsx)("div",{className:"margin-top--md",children:a.map(((e,n)=>(0,r.cloneElement)(e,{key:n,hidden:e.props.value!==o})))})}function y(e){const n=m(e);return(0,v.jsxs)("div",{className:(0,o.A)("tabs-container",b.tabList),children:[(0,v.jsx)(x,{...e,...n}),(0,v.jsx)(j,{...e,...n})]})}function w(e){const n=(0,f.A)();return(0,v.jsx)(y,{...e,children:d(e.children)},String(n))}},8453:(e,n,t)=>{t.d(n,{R:()=>s,x:()=>l});var r=t(6540);const o={},a=r.createContext(o);function s(e){const n=r.useContext(a);return r.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function l(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(o):e.components||o:s(e.components),r.createElement(a.Provider,{value:n},e.children)}}}]);