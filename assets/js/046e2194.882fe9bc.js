"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[3568],{7607:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>i,default:()=>p,frontMatter:()=>o,metadata:()=>u,toc:()=>d});var r=n(4848),a=n(8453),l=n(1470),s=n(9365);const o={},i="Install",u={id:"install",title:"Install",description:"This installation method omits detailed version metadata in generated file headers.",source:"@site/docs/install.mdx",sourceDirName:".",slug:"/install",permalink:"/docs/install",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/install.mdx",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"About",permalink:"/docs/about"},next:{title:"Getting Started",permalink:"/docs/getting-started"}},c={},d=[{value:"Dependencies",id:"dependencies",level:2}];function h(e){const t={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",strong:"strong",...(0,a.R)(),...e.components},{Details:n}=t;return n||function(e,t){throw new Error("Expected "+(t?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(t.h1,{id:"install",children:"Install"}),"\n",(0,r.jsxs)(l.A,{children:[(0,r.jsx)(s.A,{value:"brew",label:"brew",children:(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-sh",children:"brew install cludden/formula/protoc-gen-go_temporal\n"})})}),(0,r.jsx)(s.A,{value:"curl",label:"curl",children:(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-sh",children:"curl -L https://raw.githubusercontent.com/cludden/protoc-gen-go-temporal/main/scripts/install.sh | bash\n"})})}),(0,r.jsxs)(s.A,{value:"go",label:"go",children:[(0,r.jsx)(t.admonition,{type:"warning",children:(0,r.jsx)(t.p,{children:"This installation method omits detailed version metadata in generated file headers."})}),(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-sh",children:"go install github.com/cludden/protoc-gen-go-temporal/cmd/protoc-gen-go_temporal@<version>\n"})})]})]}),"\n",(0,r.jsx)(t.h2,{id:"dependencies",children:"Dependencies"}),"\n",(0,r.jsxs)(n,{children:[(0,r.jsxs)("summary",{children:["1. Install ",(0,r.jsx)(t.strong,{children:(0,r.jsx)(t.a,{href:"https://docs.buf.build/installation",children:"buf"})})]}),(0,r.jsxs)(l.A,{children:[(0,r.jsx)(s.A,{value:"brew",label:"brew",children:(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-sh",children:"brew install bufbuild/buf/buf\n"})})}),(0,r.jsx)(s.A,{value:"curl",label:"curl",children:(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-sh",children:'# Substitute BIN for your bin directory.\n# Substitute VERSION for the current released version.\nBIN="/usr/local/bin" && \\\nVERSION="1.30.0" && \\\ncurl -sSL \\\n"https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \\\n-o "${BIN}/buf" && \\\nchmod +x "${BIN}/buf"\n'})})})]})]}),"\n",(0,r.jsxs)(n,{children:[(0,r.jsxs)("summary",{children:["2. Install ",(0,r.jsx)(t.strong,{children:(0,r.jsx)(t.a,{href:"https://github.com/golang/protobuf",children:"protoc-gen-go"})})]}),(0,r.jsxs)(l.A,{children:[(0,r.jsx)(s.A,{value:"brew",label:"brew",children:(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-sh",children:"brew install protoc-gen-go\n"})})}),(0,r.jsx)(s.A,{value:"go",label:"go",children:(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-sh",children:"go install google.golang.org/protobuf/cmd/protoc-gen-go@latest\n"})})})]})]}),"\n",(0,r.jsxs)(n,{children:[(0,r.jsxs)("summary",{children:["3. Install ",(0,r.jsx)(t.strong,{children:(0,r.jsx)(t.a,{href:"https://docs.temporal.io/cli#install",children:"temporal"})})]}),(0,r.jsxs)(l.A,{children:[(0,r.jsx)(s.A,{value:"brew",label:"brew",children:(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-sh",children:"brew install temporal\n"})})}),(0,r.jsx)(s.A,{value:"curl",label:"curl",children:(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-sh",children:"curl -sSf https://temporal.download/cli.sh | sh\n"})})})]})]})]})}function p(e={}){const{wrapper:t}={...(0,a.R)(),...e.components};return t?(0,r.jsx)(t,{...e,children:(0,r.jsx)(h,{...e})}):h(e)}},9365:(e,t,n)=>{n.d(t,{A:()=>s});n(6540);var r=n(4164);const a={tabItem:"tabItem_Ymn6"};var l=n(4848);function s(e){let{children:t,hidden:n,className:s}=e;return(0,l.jsx)("div",{role:"tabpanel",className:(0,r.A)(a.tabItem,s),hidden:n,children:t})}},1470:(e,t,n)=>{n.d(t,{A:()=>y});var r=n(6540),a=n(4164),l=n(3104),s=n(6347),o=n(205),i=n(7485),u=n(1682),c=n(9466);function d(e){return r.Children.toArray(e).filter((e=>"\n"!==e)).map((e=>{if(!e||(0,r.isValidElement)(e)&&function(e){const{props:t}=e;return!!t&&"object"==typeof t&&"value"in t}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))?.filter(Boolean)??[]}function h(e){const{values:t,children:n}=e;return(0,r.useMemo)((()=>{const e=t??function(e){return d(e).map((e=>{let{props:{value:t,label:n,attributes:r,default:a}}=e;return{value:t,label:n,attributes:r,default:a}}))}(n);return function(e){const t=(0,u.X)(e,((e,t)=>e.value===t.value));if(t.length>0)throw new Error(`Docusaurus error: Duplicate values "${t.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[t,n])}function p(e){let{value:t,tabValues:n}=e;return n.some((e=>e.value===t))}function b(e){let{queryString:t=!1,groupId:n}=e;const a=(0,s.W6)(),l=function(e){let{queryString:t=!1,groupId:n}=e;if("string"==typeof t)return t;if(!1===t)return null;if(!0===t&&!n)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return n??null}({queryString:t,groupId:n});return[(0,i.aZ)(l),(0,r.useCallback)((e=>{if(!l)return;const t=new URLSearchParams(a.location.search);t.set(l,e),a.replace({...a.location,search:t.toString()})}),[l,a])]}function m(e){const{defaultValue:t,queryString:n=!1,groupId:a}=e,l=h(e),[s,i]=(0,r.useState)((()=>function(e){let{defaultValue:t,tabValues:n}=e;if(0===n.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(t){if(!p({value:t,tabValues:n}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${t}" but none of its children has the corresponding value. Available values are: ${n.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return t}const r=n.find((e=>e.default))??n[0];if(!r)throw new Error("Unexpected error: 0 tabValues");return r.value}({defaultValue:t,tabValues:l}))),[u,d]=b({queryString:n,groupId:a}),[m,f]=function(e){let{groupId:t}=e;const n=function(e){return e?`docusaurus.tab.${e}`:null}(t),[a,l]=(0,c.Dv)(n);return[a,(0,r.useCallback)((e=>{n&&l.set(e)}),[n,l])]}({groupId:a}),g=(()=>{const e=u??m;return p({value:e,tabValues:l})?e:null})();(0,o.A)((()=>{g&&i(g)}),[g]);return{selectedValue:s,selectValue:(0,r.useCallback)((e=>{if(!p({value:e,tabValues:l}))throw new Error(`Can't select invalid tab value=${e}`);i(e),d(e),f(e)}),[d,f,l]),tabValues:l}}var f=n(2303);const g={tabList:"tabList__CuJ",tabItem:"tabItem_LNqP"};var x=n(4848);function v(e){let{className:t,block:n,selectedValue:r,selectValue:s,tabValues:o}=e;const i=[],{blockElementScrollPositionUntilNextRender:u}=(0,l.a_)(),c=e=>{const t=e.currentTarget,n=i.indexOf(t),a=o[n].value;a!==r&&(u(t),s(a))},d=e=>{let t=null;switch(e.key){case"Enter":c(e);break;case"ArrowRight":{const n=i.indexOf(e.currentTarget)+1;t=i[n]??i[0];break}case"ArrowLeft":{const n=i.indexOf(e.currentTarget)-1;t=i[n]??i[i.length-1];break}}t?.focus()};return(0,x.jsx)("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,a.A)("tabs",{"tabs--block":n},t),children:o.map((e=>{let{value:t,label:n,attributes:l}=e;return(0,x.jsx)("li",{role:"tab",tabIndex:r===t?0:-1,"aria-selected":r===t,ref:e=>i.push(e),onKeyDown:d,onClick:c,...l,className:(0,a.A)("tabs__item",g.tabItem,l?.className,{"tabs__item--active":r===t}),children:n??t},t)}))})}function j(e){let{lazy:t,children:n,selectedValue:a}=e;const l=(Array.isArray(n)?n:[n]).filter(Boolean);if(t){const e=l.find((e=>e.props.value===a));return e?(0,r.cloneElement)(e,{className:"margin-top--md"}):null}return(0,x.jsx)("div",{className:"margin-top--md",children:l.map(((e,t)=>(0,r.cloneElement)(e,{key:t,hidden:e.props.value!==a})))})}function w(e){const t=m(e);return(0,x.jsxs)("div",{className:(0,a.A)("tabs-container",g.tabList),children:[(0,x.jsx)(v,{...e,...t}),(0,x.jsx)(j,{...e,...t})]})}function y(e){const t=(0,f.A)();return(0,x.jsx)(w,{...e,children:d(e.children)},String(t))}},8453:(e,t,n)=>{n.d(t,{R:()=>s,x:()=>o});var r=n(6540);const a={},l=r.createContext(a);function s(e){const t=r.useContext(l);return r.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function o(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(a):e.components||a:s(e.components),r.createElement(l.Provider,{value:t},e.children)}}}]);