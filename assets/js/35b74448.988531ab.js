"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[9614],{5765:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>i,contentTitle:()=>s,default:()=>d,frontMatter:()=>o,metadata:()=>l,toc:()=>u});var a=n(4848),r=n(8453);n(1470),n(9365);const o={},s="Bloblang",l={id:"guides/bloblang",title:"Bloblang",description:"Default workflow IDs, update IDs, and search attributes can be defined using Bloblang expressions via the ${!} interpolation syntax. The expression is evaluated against the protojson serialized input, allowing it to leverage fields from the input parameter, as well as Bloblang's native functions and methods.",source:"@site/docs/guides/bloblang.mdx",sourceDirName:"guides",slug:"/guides/bloblang",permalink:"/protoc-gen-go-temporal/docs/guides/bloblang",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/guides/bloblang.mdx",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"Activities",permalink:"/protoc-gen-go-temporal/docs/guides/activities"},next:{title:"CLI",permalink:"/protoc-gen-go-temporal/docs/guides/cli"}},i={},u=[{value:"Example",id:"example",level:2}];function c(e){const t={a:"a",code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",...(0,r.R)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(t.h1,{id:"bloblang",children:"Bloblang"}),"\n",(0,a.jsxs)(t.p,{children:["Default workflow IDs, update IDs, and search attributes can be defined using ",(0,a.jsx)(t.a,{href:"https://www.benthos.dev/docs/guides/bloblang/about",children:"Bloblang"})," expressions via the ",(0,a.jsx)(t.code,{children:"${!<expression>}"})," interpolation syntax. The expression is evaluated against the protojson serialized input, allowing it to leverage fields from the input parameter, as well as Bloblang's native ",(0,a.jsx)(t.a,{href:"https://www.benthos.dev/docs/guides/bloblang/functions",children:"functions"})," and ",(0,a.jsx)(t.a,{href:"https://www.benthos.dev/docs/guides/bloblang/methods",children:"methods"}),"."]}),"\n",(0,a.jsx)(t.h2,{id:"example",children:"Example"}),"\n",(0,a.jsx)(t.pre,{children:(0,a.jsx)(t.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3"\n\npackage example.v1;\n\nimport "google/protobuf/empty.proto";\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  rpc SayGreeting(SayGreetingRequest) returns (google.protobuf.Empty) {\n    option (temporal.v1.workflow) = {\n      id: \'say-greeting/${! greeting.or("hello").capitalize() }/${! subject.or("world").capitalize() }/${! uuid_v4() }\'\n    };\n  }\n}\n\nmessage SayGreetingRequest {\n  string greeting = 1;\n  string subject = 2;\n}\n'})}),"\n",(0,a.jsx)(t.pre,{children:(0,a.jsx)(t.code,{className:"language-go",metastring:'title="main.go"',children:'c, _ := client.Dial(client.Options{})\nexample := examplev1.NewClient(c)\n\nrun, _ := example.ExecuteSayGreeting(context.Background(), &examplev1.SayGreetingRequest{})\nrequire.Regexp(`^say-greeting/Hello/World/[a-f0-9-]{32}$`, run.ID())\n\nrun, _ := example.ExecuteSayGreeting(context.Background(), &examplev1.SayGreetingRequest{\n    Greeting: "howdy",\n    Subject: "stranger",\n})\nrequire.Regexp(`^say-greeting/Howdy/Stranger/[a-f0-9-]{32}$`, run.ID())\n'})})]})}function d(e={}){const{wrapper:t}={...(0,r.R)(),...e.components};return t?(0,a.jsx)(t,{...e,children:(0,a.jsx)(c,{...e})}):c(e)}},9365:(e,t,n)=>{n.d(t,{A:()=>s});n(6540);var a=n(4164);const r={tabItem:"tabItem_Ymn6"};var o=n(4848);function s(e){let{children:t,hidden:n,className:s}=e;return(0,o.jsx)("div",{role:"tabpanel",className:(0,a.A)(r.tabItem,s),hidden:n,children:t})}},1470:(e,t,n)=>{n.d(t,{A:()=>j});var a=n(6540),r=n(4164),o=n(3104),s=n(6347),l=n(205),i=n(7485),u=n(1682),c=n(9466);function d(e){return a.Children.toArray(e).filter((e=>"\n"!==e)).map((e=>{if(!e||(0,a.isValidElement)(e)&&function(e){const{props:t}=e;return!!t&&"object"==typeof t&&"value"in t}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))?.filter(Boolean)??[]}function p(e){const{values:t,children:n}=e;return(0,a.useMemo)((()=>{const e=t??function(e){return d(e).map((e=>{let{props:{value:t,label:n,attributes:a,default:r}}=e;return{value:t,label:n,attributes:a,default:r}}))}(n);return function(e){const t=(0,u.X)(e,((e,t)=>e.value===t.value));if(t.length>0)throw new Error(`Docusaurus error: Duplicate values "${t.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[t,n])}function g(e){let{value:t,tabValues:n}=e;return n.some((e=>e.value===t))}function b(e){let{queryString:t=!1,groupId:n}=e;const r=(0,s.W6)(),o=function(e){let{queryString:t=!1,groupId:n}=e;if("string"==typeof t)return t;if(!1===t)return null;if(!0===t&&!n)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return n??null}({queryString:t,groupId:n});return[(0,i.aZ)(o),(0,a.useCallback)((e=>{if(!o)return;const t=new URLSearchParams(r.location.search);t.set(o,e),r.replace({...r.location,search:t.toString()})}),[o,r])]}function m(e){const{defaultValue:t,queryString:n=!1,groupId:r}=e,o=p(e),[s,i]=(0,a.useState)((()=>function(e){let{defaultValue:t,tabValues:n}=e;if(0===n.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(t){if(!g({value:t,tabValues:n}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${t}" but none of its children has the corresponding value. Available values are: ${n.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return t}const a=n.find((e=>e.default))??n[0];if(!a)throw new Error("Unexpected error: 0 tabValues");return a.value}({defaultValue:t,tabValues:o}))),[u,d]=b({queryString:n,groupId:r}),[m,h]=function(e){let{groupId:t}=e;const n=function(e){return e?`docusaurus.tab.${e}`:null}(t),[r,o]=(0,c.Dv)(n);return[r,(0,a.useCallback)((e=>{n&&o.set(e)}),[n,o])]}({groupId:r}),f=(()=>{const e=u??m;return g({value:e,tabValues:o})?e:null})();(0,l.A)((()=>{f&&i(f)}),[f]);return{selectedValue:s,selectValue:(0,a.useCallback)((e=>{if(!g({value:e,tabValues:o}))throw new Error(`Can't select invalid tab value=${e}`);i(e),d(e),h(e)}),[d,h,o]),tabValues:o}}var h=n(2303);const f={tabList:"tabList__CuJ",tabItem:"tabItem_LNqP"};var v=n(4848);function x(e){let{className:t,block:n,selectedValue:a,selectValue:s,tabValues:l}=e;const i=[],{blockElementScrollPositionUntilNextRender:u}=(0,o.a_)(),c=e=>{const t=e.currentTarget,n=i.indexOf(t),r=l[n].value;r!==a&&(u(t),s(r))},d=e=>{let t=null;switch(e.key){case"Enter":c(e);break;case"ArrowRight":{const n=i.indexOf(e.currentTarget)+1;t=i[n]??i[0];break}case"ArrowLeft":{const n=i.indexOf(e.currentTarget)-1;t=i[n]??i[i.length-1];break}}t?.focus()};return(0,v.jsx)("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,r.A)("tabs",{"tabs--block":n},t),children:l.map((e=>{let{value:t,label:n,attributes:o}=e;return(0,v.jsx)("li",{role:"tab",tabIndex:a===t?0:-1,"aria-selected":a===t,ref:e=>i.push(e),onKeyDown:d,onClick:c,...o,className:(0,r.A)("tabs__item",f.tabItem,o?.className,{"tabs__item--active":a===t}),children:n??t},t)}))})}function w(e){let{lazy:t,children:n,selectedValue:r}=e;const o=(Array.isArray(n)?n:[n]).filter(Boolean);if(t){const e=o.find((e=>e.props.value===r));return e?(0,a.cloneElement)(e,{className:"margin-top--md"}):null}return(0,v.jsx)("div",{className:"margin-top--md",children:o.map(((e,t)=>(0,a.cloneElement)(e,{key:t,hidden:e.props.value!==r})))})}function y(e){const t=m(e);return(0,v.jsxs)("div",{className:(0,r.A)("tabs-container",f.tabList),children:[(0,v.jsx)(x,{...e,...t}),(0,v.jsx)(w,{...e,...t})]})}function j(e){const t=(0,h.A)();return(0,v.jsx)(y,{...e,children:d(e.children)},String(t))}},8453:(e,t,n)=>{n.d(t,{R:()=>s,x:()=>l});var a=n(6540);const r={},o=a.createContext(r);function s(e){const t=a.useContext(o);return a.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function l(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(r):e.components||r:s(e.components),a.createElement(o.Provider,{value:t},e.children)}}}]);