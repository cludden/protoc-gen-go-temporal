"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[7104],{58:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>u,default:()=>d,frontMatter:()=>s,metadata:()=>o,toc:()=>i});var r=n(4848),a=n(8453);n(1470),n(9365);const s={},u="Testing",o={id:"guides/testing",title:"Testing",description:"Coming Soon...",source:"@site/docs/guides/testing.mdx",sourceDirName:"guides",slug:"/guides/testing",permalink:"/docs/guides/testing",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/guides/testing.mdx",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"Signals",permalink:"/docs/guides/signals"},next:{title:"Updates",permalink:"/docs/guides/updates"}},l={},i=[];function c(e){const t={h1:"h1",p:"p",strong:"strong",...(0,a.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(t.h1,{id:"testing",children:"Testing"}),"\n",(0,r.jsx)(t.p,{children:(0,r.jsx)(t.strong,{children:"Coming Soon..."})})]})}function d(e={}){const{wrapper:t}={...(0,a.R)(),...e.components};return t?(0,r.jsx)(t,{...e,children:(0,r.jsx)(c,{...e})}):c(e)}},9365:(e,t,n)=>{n.d(t,{A:()=>u});n(6540);var r=n(4164);const a={tabItem:"tabItem_Ymn6"};var s=n(4848);function u(e){let{children:t,hidden:n,className:u}=e;return(0,s.jsx)("div",{role:"tabpanel",className:(0,r.A)(a.tabItem,u),hidden:n,children:t})}},1470:(e,t,n)=>{n.d(t,{A:()=>k});var r=n(6540),a=n(4164),s=n(3104),u=n(6347),o=n(205),l=n(7485),i=n(1682),c=n(9466);function d(e){return r.Children.toArray(e).filter((e=>"\n"!==e)).map((e=>{if(!e||(0,r.isValidElement)(e)&&function(e){const{props:t}=e;return!!t&&"object"==typeof t&&"value"in t}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))?.filter(Boolean)??[]}function p(e){const{values:t,children:n}=e;return(0,r.useMemo)((()=>{const e=t??function(e){return d(e).map((e=>{let{props:{value:t,label:n,attributes:r,default:a}}=e;return{value:t,label:n,attributes:r,default:a}}))}(n);return function(e){const t=(0,i.X)(e,((e,t)=>e.value===t.value));if(t.length>0)throw new Error(`Docusaurus error: Duplicate values "${t.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[t,n])}function f(e){let{value:t,tabValues:n}=e;return n.some((e=>e.value===t))}function m(e){let{queryString:t=!1,groupId:n}=e;const a=(0,u.W6)(),s=function(e){let{queryString:t=!1,groupId:n}=e;if("string"==typeof t)return t;if(!1===t)return null;if(!0===t&&!n)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return n??null}({queryString:t,groupId:n});return[(0,l.aZ)(s),(0,r.useCallback)((e=>{if(!s)return;const t=new URLSearchParams(a.location.search);t.set(s,e),a.replace({...a.location,search:t.toString()})}),[s,a])]}function b(e){const{defaultValue:t,queryString:n=!1,groupId:a}=e,s=p(e),[u,l]=(0,r.useState)((()=>function(e){let{defaultValue:t,tabValues:n}=e;if(0===n.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(t){if(!f({value:t,tabValues:n}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${t}" but none of its children has the corresponding value. Available values are: ${n.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return t}const r=n.find((e=>e.default))??n[0];if(!r)throw new Error("Unexpected error: 0 tabValues");return r.value}({defaultValue:t,tabValues:s}))),[i,d]=m({queryString:n,groupId:a}),[b,h]=function(e){let{groupId:t}=e;const n=function(e){return e?`docusaurus.tab.${e}`:null}(t),[a,s]=(0,c.Dv)(n);return[a,(0,r.useCallback)((e=>{n&&s.set(e)}),[n,s])]}({groupId:a}),g=(()=>{const e=i??b;return f({value:e,tabValues:s})?e:null})();(0,o.A)((()=>{g&&l(g)}),[g]);return{selectedValue:u,selectValue:(0,r.useCallback)((e=>{if(!f({value:e,tabValues:s}))throw new Error(`Can't select invalid tab value=${e}`);l(e),d(e),h(e)}),[d,h,s]),tabValues:s}}var h=n(2303);const g={tabList:"tabList__CuJ",tabItem:"tabItem_LNqP"};var v=n(4848);function x(e){let{className:t,block:n,selectedValue:r,selectValue:u,tabValues:o}=e;const l=[],{blockElementScrollPositionUntilNextRender:i}=(0,s.a_)(),c=e=>{const t=e.currentTarget,n=l.indexOf(t),a=o[n].value;a!==r&&(i(t),u(a))},d=e=>{let t=null;switch(e.key){case"Enter":c(e);break;case"ArrowRight":{const n=l.indexOf(e.currentTarget)+1;t=l[n]??l[0];break}case"ArrowLeft":{const n=l.indexOf(e.currentTarget)-1;t=l[n]??l[l.length-1];break}}t?.focus()};return(0,v.jsx)("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,a.A)("tabs",{"tabs--block":n},t),children:o.map((e=>{let{value:t,label:n,attributes:s}=e;return(0,v.jsx)("li",{role:"tab",tabIndex:r===t?0:-1,"aria-selected":r===t,ref:e=>l.push(e),onKeyDown:d,onClick:c,...s,className:(0,a.A)("tabs__item",g.tabItem,s?.className,{"tabs__item--active":r===t}),children:n??t},t)}))})}function y(e){let{lazy:t,children:n,selectedValue:a}=e;const s=(Array.isArray(n)?n:[n]).filter(Boolean);if(t){const e=s.find((e=>e.props.value===a));return e?(0,r.cloneElement)(e,{className:"margin-top--md"}):null}return(0,v.jsx)("div",{className:"margin-top--md",children:s.map(((e,t)=>(0,r.cloneElement)(e,{key:t,hidden:e.props.value!==a})))})}function w(e){const t=b(e);return(0,v.jsxs)("div",{className:(0,a.A)("tabs-container",g.tabList),children:[(0,v.jsx)(x,{...e,...t}),(0,v.jsx)(y,{...e,...t})]})}function k(e){const t=(0,h.A)();return(0,v.jsx)(w,{...e,children:d(e.children)},String(t))}},8453:(e,t,n)=>{n.d(t,{R:()=>u,x:()=>o});var r=n(6540);const a={},s=r.createContext(a);function u(e){const t=r.useContext(s);return r.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function o(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(a):e.components||a:u(e.components),r.createElement(s.Provider,{value:t},e.children)}}}]);