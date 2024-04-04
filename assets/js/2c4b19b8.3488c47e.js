"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[940],{2922:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>i,contentTitle:()=>l,default:()=>d,frontMatter:()=>a,metadata:()=>s,toc:()=>c});var n=r(4848),o=r(8453);r(1470),r(9365);const a={},l="Codec Server",s={id:"guides/codec-server",title:"Codec Server",description:"Data Converter",source:"@site/docs/guides/codec-server.mdx",sourceDirName:"guides",slug:"/guides/codec-server",permalink:"/protoc-gen-go-temporal/docs/guides/codec-server",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/guides/codec-server.mdx",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"Bloblang",permalink:"/protoc-gen-go-temporal/docs/guides/bloblang"},next:{title:"Documentation",permalink:"/protoc-gen-go-temporal/docs/guides/documentation"}},i={},c=[{value:"Data Converter",id:"data-converter",level:2},{value:"Codec Server",id:"codec-server-1",level:2}];function u(e){const t={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",p:"p",pre:"pre",...(0,o.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(t.h1,{id:"codec-server",children:"Codec Server"}),"\n",(0,n.jsx)(t.h2,{id:"data-converter",children:"Data Converter"}),"\n",(0,n.jsxs)(t.p,{children:["Temporal's ",(0,n.jsx)(t.a,{href:"https://pkg.go.dev/go.temporal.io/sdk/converter#GetDefaultDataConverter",children:"default data converter"})," will serialize protobuf types using the ",(0,n.jsx)(t.code,{children:"json/protobuf"})," encoding provided by the ",(0,n.jsx)(t.a,{href:"https://pkg.go.dev/go.temporal.io/sdk/converter#ProtoJSONPayloadConverter",children:"ProtoJSONPayloadConverter"}),", which allows the Temporal UI to automatically decode the underlying payload and render it as JSON. If you'd prefer to take advantage of protobuf's binary format for smaller payloads, you can provide an alternative data converter to the Temporal client at initialization that prioritizes the ",(0,n.jsx)(t.a,{href:"https://pkg.go.dev/go.temporal.io/sdk/converter#ProtoPayloadConverter",children:"ProtoPayloadConverter"})," ahead of the ",(0,n.jsx)(t.code,{children:"ProtoJSONPayloadConverter"}),". See below for an example."]}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-go",metastring:'title="worker/main.go"',children:'package main\n\nimport (\n\t"log"\n\t"log/slog"\n\t"os"\n\n\t"path/to/interal/example"\n\texamplev1 "path/to/gen/example/v1"\n\t"github.com/urfave/cli/v2"\n\t"go.temporal.io/sdk/client"\n\t"go.temporal.io/sdk/converter"\n\tsdklog "go.temporal.io/sdk/log"\n\t"go.temporal.io/sdk/worker"\n)\n\nfunc main() {\n\tapp, err := examplev1.NewExampleCli(\n\t\texamplev1.NewExampleCliOptions().\n\t\t\tWithClient(func(cmd *cli.Context) (client.Client, error) {\n\t\t\t\treturn client.Dial(client.Options{\n\t\t\t\t\tDataConverter: converter.NewCompositeDataConverter(\n\t\t\t\t\t\tconverter.NewNilPayloadConverter(),\n\t\t\t\t\t\tconverter.NewByteSlicePayloadConverter(),\n\t\t\t\t\t\tconverter.NewProtoPayloadConverter(),\n\t\t\t\t\t),\n\t\t\t\t\tLogger: sdklog.NewStructuredLogger(slog.Default()),\n\t\t\t\t})\n\t\t\t}).\n\t\t\tWithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {\n\t\t\t\tw := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})\n\t\t\t\texamplev1.RegisterExampleActivities(w, &example.Activities{})\n\t\t\t\texamplev1.RegisterExampleWorkflows(w, &example.Workflows{})\n\t\t\t\treturn w, nil\n\t\t\t}),\n\t)\n\tif err != nil {\n\t\tlog.Fatalf("error initializing example cli: %v", err)\n\t}\n\n    if err := app.Run(os.Args); err != nil {\n        log.Fatal(err)\n    }\n}\n'})}),"\n",(0,n.jsx)(t.h2,{id:"codec-server-1",children:"Codec Server"}),"\n",(0,n.jsxs)(t.p,{children:["If you choose to use ",(0,n.jsx)(t.code,{children:"binary/protobuf"})," encoding, you'll lose the ability to view decoded payloads in the Temporal UI unless you configure the ",(0,n.jsx)(t.a,{href:"https://docs.temporal.io/dataconversion#codec-server",children:"Remote Codec Server"})," integration. The plugin can generate helpers that simplify the process of implementing a remote codec server for use with the Temporal UI to support conversion between ",(0,n.jsx)(t.code,{children:"binary/protobuf"})," and ",(0,n.jsx)(t.code,{children:"json/protobuf"})," or ",(0,n.jsx)(t.code,{children:"json/plain"})," payload encodings. See below for a simple example. For a more advanced example that supports different codecs per namespace, cors, and authentication, see the ",(0,n.jsx)(t.a,{href:"https://github.com/temporalio/samples-go/blob/main/codec-server/codec-server/main.go",children:"codec-server"})," go sample."]}),"\n",(0,n.jsx)(t.admonition,{type:"info",children:(0,n.jsxs)(t.p,{children:["This requires the ",(0,n.jsx)(t.a,{href:"/docs/configuration/plugin#enable-codec",children:"enable-codec"})," plugin option to be enabled"]})}),"\n",(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-go",metastring:'title="codecserver/main.go"',children:'package main\n\nimport (\n\t"context"\n\t"errors"\n\t"log"\n\t"log/slog"\n\t"net/http"\n\t"os"\n\t"os/signal"\n\t"syscall"\n\n\texamplev1 "path/to/gen/example/v1"\n\t"github.com/cludden/protoc-gen-go-temporal/pkg/codec"\n\t"github.com/cludden/protoc-gen-go-temporal/pkg/scheme"\n\t"github.com/urfave/cli/v2"\n\t"go.temporal.io/sdk/converter"\n)\n\nfunc main() {\n\tapp, err := examplev1.NewExampleCli(/* ... */)\n\tif err != nil {\n\t\tlog.Fatalf("error initializing example cli: %v", err)\n\t}\n\n\tapp.Commands = append(app.Commands, &cli.Command{\n\t\tName:  "codec",\n\t\tUsage: "run remote codec server",\n\t\tAction: func(cmd *cli.Context) error {\n\t\t\thandler := converter.NewPayloadCodecHTTPHandler(\n\t\t\t\tcodec.NewProtoJSONCodec(\n\t\t\t\t\tscheme.New(\n\t\t\t\t\t\texamplev1.WithExampleSchemeTypes(),\n\t\t\t\t\t),\n\t\t\t\t),\n\t\t\t)\n\n\t\t\tsrv := &http.Server{\n\t\t\t\tAddr:    "0.0.0.0:8080",\n\t\t\t\tHandler: handler,\n\t\t\t}\n\n\t\t\tgo func() {\n\t\t\t\tsigChan := make(chan os.Signal, 1)\n\t\t\t\tsignal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)\n\t\t\t\t<-sigChan\n\n\t\t\t\tif err := srv.Shutdown(context.Background()); err != nil {\n\t\t\t\t\tlog.Fatalf("error shutting down server: %v", err)\n\t\t\t\t}\n\t\t\t}()\n\n\t\t\tif err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {\n\t\t\t\tlog.Fatalf("server error: %v", err)\n\t\t\t}\n\t\t\treturn nil\n\t\t},\n\t})\n\n\t// run cli\n\tif err := app.Run(os.Args); err != nil {\n\t\tlog.Fatal(err)\n\t}\n}\n'})}),"\n",(0,n.jsxs)(t.p,{children:["See the ",(0,n.jsx)(t.a,{href:"/docs/examples/codecserver/",children:"codecserver"})," example for more details."]})]})}function d(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(u,{...e})}):u(e)}},9365:(e,t,r)=>{r.d(t,{A:()=>l});r(6540);var n=r(4164);const o={tabItem:"tabItem_Ymn6"};var a=r(4848);function l(e){let{children:t,hidden:r,className:l}=e;return(0,a.jsx)("div",{role:"tabpanel",className:(0,n.A)(o.tabItem,l),hidden:r,children:t})}},1470:(e,t,r)=>{r.d(t,{A:()=>k});var n=r(6540),o=r(4164),a=r(3104),l=r(6347),s=r(205),i=r(7485),c=r(1682),u=r(9466);function d(e){return n.Children.toArray(e).filter((e=>"\n"!==e)).map((e=>{if(!e||(0,n.isValidElement)(e)&&function(e){const{props:t}=e;return!!t&&"object"==typeof t&&"value"in t}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))?.filter(Boolean)??[]}function p(e){const{values:t,children:r}=e;return(0,n.useMemo)((()=>{const e=t??function(e){return d(e).map((e=>{let{props:{value:t,label:r,attributes:n,default:o}}=e;return{value:t,label:r,attributes:n,default:o}}))}(r);return function(e){const t=(0,c.X)(e,((e,t)=>e.value===t.value));if(t.length>0)throw new Error(`Docusaurus error: Duplicate values "${t.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[t,r])}function m(e){let{value:t,tabValues:r}=e;return r.some((e=>e.value===t))}function h(e){let{queryString:t=!1,groupId:r}=e;const o=(0,l.W6)(),a=function(e){let{queryString:t=!1,groupId:r}=e;if("string"==typeof t)return t;if(!1===t)return null;if(!0===t&&!r)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return r??null}({queryString:t,groupId:r});return[(0,i.aZ)(a),(0,n.useCallback)((e=>{if(!a)return;const t=new URLSearchParams(o.location.search);t.set(a,e),o.replace({...o.location,search:t.toString()})}),[a,o])]}function v(e){const{defaultValue:t,queryString:r=!1,groupId:o}=e,a=p(e),[l,i]=(0,n.useState)((()=>function(e){let{defaultValue:t,tabValues:r}=e;if(0===r.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(t){if(!m({value:t,tabValues:r}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${t}" but none of its children has the corresponding value. Available values are: ${r.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return t}const n=r.find((e=>e.default))??r[0];if(!n)throw new Error("Unexpected error: 0 tabValues");return n.value}({defaultValue:t,tabValues:a}))),[c,d]=h({queryString:r,groupId:o}),[v,g]=function(e){let{groupId:t}=e;const r=function(e){return e?`docusaurus.tab.${e}`:null}(t),[o,a]=(0,u.Dv)(r);return[o,(0,n.useCallback)((e=>{r&&a.set(e)}),[r,a])]}({groupId:o}),f=(()=>{const e=c??v;return m({value:e,tabValues:a})?e:null})();(0,s.A)((()=>{f&&i(f)}),[f]);return{selectedValue:l,selectValue:(0,n.useCallback)((e=>{if(!m({value:e,tabValues:a}))throw new Error(`Can't select invalid tab value=${e}`);i(e),d(e),g(e)}),[d,g,a]),tabValues:a}}var g=r(2303);const f={tabList:"tabList__CuJ",tabItem:"tabItem_LNqP"};var b=r(4848);function x(e){let{className:t,block:r,selectedValue:n,selectValue:l,tabValues:s}=e;const i=[],{blockElementScrollPositionUntilNextRender:c}=(0,a.a_)(),u=e=>{const t=e.currentTarget,r=i.indexOf(t),o=s[r].value;o!==n&&(c(t),l(o))},d=e=>{let t=null;switch(e.key){case"Enter":u(e);break;case"ArrowRight":{const r=i.indexOf(e.currentTarget)+1;t=i[r]??i[0];break}case"ArrowLeft":{const r=i.indexOf(e.currentTarget)-1;t=i[r]??i[i.length-1];break}}t?.focus()};return(0,b.jsx)("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,o.A)("tabs",{"tabs--block":r},t),children:s.map((e=>{let{value:t,label:r,attributes:a}=e;return(0,b.jsx)("li",{role:"tab",tabIndex:n===t?0:-1,"aria-selected":n===t,ref:e=>i.push(e),onKeyDown:d,onClick:u,...a,className:(0,o.A)("tabs__item",f.tabItem,a?.className,{"tabs__item--active":n===t}),children:r??t},t)}))})}function y(e){let{lazy:t,children:r,selectedValue:o}=e;const a=(Array.isArray(r)?r:[r]).filter(Boolean);if(t){const e=a.find((e=>e.props.value===o));return e?(0,n.cloneElement)(e,{className:"margin-top--md"}):null}return(0,b.jsx)("div",{className:"margin-top--md",children:a.map(((e,t)=>(0,n.cloneElement)(e,{key:t,hidden:e.props.value!==o})))})}function w(e){const t=v(e);return(0,b.jsxs)("div",{className:(0,o.A)("tabs-container",f.tabList),children:[(0,b.jsx)(x,{...e,...t}),(0,b.jsx)(y,{...e,...t})]})}function k(e){const t=(0,g.A)();return(0,b.jsx)(w,{...e,children:d(e.children)},String(t))}},8453:(e,t,r)=>{r.d(t,{R:()=>l,x:()=>s});var n=r(6540);const o={},a=n.createContext(o);function l(e){const t=n.useContext(a);return n.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function s(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(o):e.components||o:l(e.components),n.createElement(a.Provider,{value:t},e.children)}}}]);