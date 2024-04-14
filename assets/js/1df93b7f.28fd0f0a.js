/*! For license information please see 1df93b7f.28fd0f0a.js.LICENSE.txt */
(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[583],{9365:(e,n,o)=>{"use strict";o.d(n,{A:()=>a});o(6540);var r=o(4164);const l={tabItem:"tabItem_Ymn6"};var t=o(4848);function a(e){let{children:n,hidden:o,className:a}=e;return(0,t.jsx)("div",{role:"tabpanel",className:(0,r.A)(l.tabItem,a),hidden:o,children:n})}},1470:(e,n,o)=>{"use strict";o.d(n,{A:()=>y});var r=o(6540),l=o(4164),t=o(3104),a=o(6347),s=o(205),i=o(7485),c=o(1682),u=o(9466);function d(e){return r.Children.toArray(e).filter((e=>"\n"!==e)).map((e=>{if(!e||(0,r.isValidElement)(e)&&function(e){const{props:n}=e;return!!n&&"object"==typeof n&&"value"in n}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))?.filter(Boolean)??[]}function p(e){const{values:n,children:o}=e;return(0,r.useMemo)((()=>{const e=n??function(e){return d(e).map((e=>{let{props:{value:n,label:o,attributes:r,default:l}}=e;return{value:n,label:o,attributes:r,default:l}}))}(o);return function(e){const n=(0,c.X)(e,((e,n)=>e.value===n.value));if(n.length>0)throw new Error(`Docusaurus error: Duplicate values "${n.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[n,o])}function g(e){let{value:n,tabValues:o}=e;return o.some((e=>e.value===n))}function h(e){let{queryString:n=!1,groupId:o}=e;const l=(0,a.W6)(),t=function(e){let{queryString:n=!1,groupId:o}=e;if("string"==typeof n)return n;if(!1===n)return null;if(!0===n&&!o)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return o??null}({queryString:n,groupId:o});return[(0,i.aZ)(t),(0,r.useCallback)((e=>{if(!t)return;const n=new URLSearchParams(l.location.search);n.set(t,e),l.replace({...l.location,search:n.toString()})}),[t,l])]}function f(e){const{defaultValue:n,queryString:o=!1,groupId:l}=e,t=p(e),[a,i]=(0,r.useState)((()=>function(e){let{defaultValue:n,tabValues:o}=e;if(0===o.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(n){if(!g({value:n,tabValues:o}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${n}" but none of its children has the corresponding value. Available values are: ${o.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return n}const r=o.find((e=>e.default))??o[0];if(!r)throw new Error("Unexpected error: 0 tabValues");return r.value}({defaultValue:n,tabValues:t}))),[c,d]=h({queryString:o,groupId:l}),[f,m]=function(e){let{groupId:n}=e;const o=function(e){return e?`docusaurus.tab.${e}`:null}(n),[l,t]=(0,u.Dv)(o);return[l,(0,r.useCallback)((e=>{o&&t.set(e)}),[o,t])]}({groupId:l}),w=(()=>{const e=c??f;return g({value:e,tabValues:t})?e:null})();(0,s.A)((()=>{w&&i(w)}),[w]);return{selectedValue:a,selectValue:(0,r.useCallback)((e=>{if(!g({value:e,tabValues:t}))throw new Error(`Can't select invalid tab value=${e}`);i(e),d(e),m(e)}),[d,m,t]),tabValues:t}}var m=o(2303);const w={tabList:"tabList__CuJ",tabItem:"tabItem_LNqP"};var x=o(4848);function b(e){let{className:n,block:o,selectedValue:r,selectValue:a,tabValues:s}=e;const i=[],{blockElementScrollPositionUntilNextRender:c}=(0,t.a_)(),u=e=>{const n=e.currentTarget,o=i.indexOf(n),l=s[o].value;l!==r&&(c(n),a(l))},d=e=>{let n=null;switch(e.key){case"Enter":u(e);break;case"ArrowRight":{const o=i.indexOf(e.currentTarget)+1;n=i[o]??i[0];break}case"ArrowLeft":{const o=i.indexOf(e.currentTarget)-1;n=i[o]??i[i.length-1];break}}n?.focus()};return(0,x.jsx)("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,l.A)("tabs",{"tabs--block":o},n),children:s.map((e=>{let{value:n,label:o,attributes:t}=e;return(0,x.jsx)("li",{role:"tab",tabIndex:r===n?0:-1,"aria-selected":r===n,ref:e=>i.push(e),onKeyDown:d,onClick:u,...t,className:(0,l.A)("tabs__item",w.tabItem,t?.className,{"tabs__item--active":r===n}),children:o??n},n)}))})}function v(e){let{lazy:n,children:o,selectedValue:l}=e;const t=(Array.isArray(o)?o:[o]).filter(Boolean);if(n){const e=t.find((e=>e.props.value===l));return e?(0,r.cloneElement)(e,{className:"margin-top--md"}):null}return(0,x.jsx)("div",{className:"margin-top--md",children:t.map(((e,n)=>(0,r.cloneElement)(e,{key:n,hidden:e.props.value!==l})))})}function k(e){const n=f(e);return(0,x.jsxs)("div",{className:(0,l.A)("tabs-container",w.tabList),children:[(0,x.jsx)(b,{...e,...n}),(0,x.jsx)(v,{...e,...n})]})}function y(e){const n=(0,m.A)();return(0,x.jsx)(k,{...e,children:d(e.children)},String(n))}},1001:(e,n,o)=>{"use strict";o.r(n),o.d(n,{default:()=>k});var r=o(4164),l=o(6942),t=o.n(l),a=o(8774),s=o(4586),i=o(781);const c={features:"features_t9lD",featureSvg:"featureSvg_GfXr"};var u=o(1470),d=o(9365),p=o(1432),g=o(4848),h=o(8453);function f(e){const n={a:"a",code:"code",h2:"h2",li:"li",p:"p",strong:"strong",ul:"ul",...(0,h.R)(),...e.components};return(0,g.jsxs)(g.Fragment,{children:[(0,g.jsx)(n.h2,{id:"about",children:"About"}),"\n",(0,g.jsx)(n.p,{children:"A protoc plugin for generating typed Temporal clients and workers in Go from protobuf schemas. This plugin allows:"}),"\n",(0,g.jsxs)(n.ul,{children:["\n",(0,g.jsx)(n.li,{children:"workflow authors to configure sensible defaults and guardrails"}),"\n",(0,g.jsx)(n.li,{children:"simplifies the implementation and testing of Temporal workers"}),"\n",(0,g.jsx)(n.li,{children:"and streamlines integration by providing typed client SDKs and a generated CLI application"}),"\n"]}),"\n",(0,g.jsx)("iframe",{width:"560",height:"315",src:"https://www.youtube.com/embed/fqKDWZDj-c0?si=3Wgvj_nP2BnSVcum&start=912",title:"YouTube video player",frameborder:"0",allow:"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share",referrerpolicy:"strict-origin-when-cross-origin",allowFullScreen:!0}),"\n",(0,g.jsx)("br",{}),"\n",(0,g.jsx)("br",{}),"\n",(0,g.jsx)(n.h2,{id:"features",children:"Features"}),"\n",(0,g.jsxs)(n.p,{children:["Generated ",(0,g.jsx)(n.strong,{children:"Client"})," with:"]}),"\n",(0,g.jsxs)(n.ul,{children:["\n",(0,g.jsx)(n.li,{children:"methods for executing workflows, queries, signals, and updates"}),"\n",(0,g.jsx)(n.li,{children:"methods for cancelling or terminating workflows"}),"\n",(0,g.jsxs)(n.li,{children:["default ",(0,g.jsx)(n.code,{children:"client.StartWorkflowOptions"})," and ",(0,g.jsx)(n.code,{children:"client.UpdateWorkflowWithOptionsRequest"})]}),"\n",(0,g.jsxs)(n.li,{children:["dynamic workflow ids, update ids, and search attributes via ",(0,g.jsx)(n.a,{href:"#bloblang-expressions",children:"Bloblang expressions"})]}),"\n",(0,g.jsx)(n.li,{children:"default timeouts, id reuse policies, retry policies, wait policies"}),"\n"]}),"\n",(0,g.jsxs)(n.p,{children:["Generated ",(0,g.jsx)(n.strong,{children:"Worker"})," resources with:"]}),"\n",(0,g.jsxs)(n.ul,{children:["\n",(0,g.jsx)(n.li,{children:"functions for calling activities and local activities from workflows"}),"\n",(0,g.jsx)(n.li,{children:"functions for executing child workflows and signalling external workflows"}),"\n",(0,g.jsxs)(n.li,{children:["default ",(0,g.jsx)(n.code,{children:"workflow.ActivityOptions"}),", ",(0,g.jsx)(n.code,{children:"workflow.ChildWorkflowOptions"})]}),"\n",(0,g.jsx)(n.li,{children:"default timeouts, parent cose policies, retry policies"}),"\n"]}),"\n",(0,g.jsxs)(n.p,{children:["Optional ",(0,g.jsx)(n.strong,{children:"CLI"})," with:"]}),"\n",(0,g.jsxs)(n.ul,{children:["\n",(0,g.jsx)(n.li,{children:"commands for executing workflows, synchronously or asynchronously"}),"\n",(0,g.jsx)(n.li,{children:"commands for starting workflows with signals, synchronously or asynchronously"}),"\n",(0,g.jsx)(n.li,{children:"commands for querying existing workflows"}),"\n",(0,g.jsx)(n.li,{children:"commands for sending signals to existing workflows"}),"\n",(0,g.jsx)(n.li,{children:"typed flags for conventiently specifying workflow, query, and signal inputs"}),"\n"]}),"\n",(0,g.jsxs)(n.p,{children:["Generated ",(0,g.jsx)(n.a,{href:"/docs/guides/xns",children:"Cross-Namespace (XNS)"})," helpers: ",(0,g.jsx)(n.strong,{children:"[Experimental]"})]}),"\n",(0,g.jsxs)(n.ul,{children:["\n",(0,g.jsx)(n.li,{children:"with support for invoking a service's workflows, queries, signals, and updates from workflows in a different temporal namespace"}),"\n"]}),"\n",(0,g.jsxs)(n.p,{children:["Generated ",(0,g.jsx)(n.a,{href:"/docs/guides/codec-server",children:"Remote Codec Server"})," helpers"]}),"\n",(0,g.jsxs)(n.p,{children:["Generated ",(0,g.jsx)(n.a,{href:"/docs/guides/documentation",children:"Markdown Documentation"})]}),"\n",(0,g.jsx)(n.h2,{id:"inspiration",children:"Inspiration"}),"\n",(0,g.jsxs)(n.p,{children:["This project was inspired by ",(0,g.jsx)(n.a,{href:"https://github.com/cretz/",children:"Chad Retz's"})," awesome ",(0,g.jsx)(n.a,{href:"https://github.com/cretz/temporal-sdk-go-advanced",children:"github.com/cretz/temporal-sdk-go-advanced"})," and ",(0,g.jsx)(n.a,{href:"https://github.com/jlegrone/",children:"Jacob LeGrone's"})," excellent Replay talk on ",(0,g.jsx)(n.a,{href:"https://youtu.be/LxgkAoTSI8Q?si=ZGwwbfbMz48MkPhj&t=681",children:"Temporal @ Datadog"})]}),"\n",(0,g.jsx)("iframe",{width:"560",height:"315",src:"https://www.youtube.com/embed/LxgkAoTSI8Q?si=L3O5it48sy38dsx7&start=681",title:"YouTube video player",frameborder:"0",allow:"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share",referrerpolicy:"strict-origin-when-cross-origin",allowFullScreen:!0})]})}function m(e={}){const{wrapper:n}={...(0,h.R)(),...e.components};return n?(0,g.jsx)(n,{...e,children:(0,g.jsx)(f,{...e})}):f(e)}const w=[{label:"Annotate",content:"Annotate your protobuf services and methods with Temporal options.",fragments:[{language:"protobuf",title:"proto/helloworld/v1/example.proto",content:'syntax="proto3";\n\npackage helloworld.v1;\n\nimport "google/protobuf/empty.proto";\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "hello-world"\n  };\n\n  // Hello prints a friendly greeting and waits for goodbye\n  rpc Hello(HelloRequest) returns (HelloResponse) {\n    option (temporal.v1.workflow) = {\n      name: "helloworld.v1.Hello"\n      id: \'hello/${! name.or("World") }\'\n      signal: { ref: "Goodbye" }\n    };\n  }\n\n  // Goodbye signals a running workflow to exit\n  rpc Goodbye(GoodbyeRequest) returns (google.protobuf.Empty) {\n    option (temporal.v1.signal) = {};\n  }\n}\n\n// HelloRequest describes the input to a Hello workflow\nmessage HelloRequest {\n  string name = 1;\n}\n\n// HelloResponse describes the output from a Hello workflow\nmessage HelloResponse {\n  string result = 1;\n}\n\nmessage GoodbyeRequest {\n  string message = 1;\n}\n        '}]},{label:"Generate",content:"Generate Go code for implementing Temporal Clients, Workers, and CLI applications.",fragments:[{language:"yaml",title:"buf.gen.yaml",content:"version: v1\nmanaged:\n  enabled: true\n  go_package_prefix:\n    default: example/gen\n    except:\n      - buf.build/cludden/protoc-gen-go-temporal\nplugins:\n  - plugin: go\n    out: gen\n    opt: paths=source_relative\n  - plugin: go_temporal\n    out: gen\n    opt: paths=source_relative,cli-enabled=true,cli-categories=true,workflow-update-enabled=true,docs-out=./proto/README.md\n    strategy: all\n    "},{language:"sh",content:"buf generate"}]},{label:"Implement",content:"Implement the required Workflow and Activity interfaces",fragments:[{language:"go",title:"internal/example/example.go",content:'package example\n\nimport (\n  helloworldv1 "path/to/gen/helloworld/v1"\n  "go.temporal.io/sdk/log"\n  "go.temporal.io/sdk/workflow"\n)\n\ntype (\n  Workflows struct{}\n\n  // HelloWorkflow provides a helloworldv1.HelloWorkflow implementation\n  HelloWorkflow struct {\n    *helloworldv1.HelloWorkflowInput\n    log log.Logger\n  }\n)\n\n// NewHelloWorkflow initializes a new helloworldv1.HelloWorkflow value\nfunc (w *Workflows) Hello(ctx workflow.Context, input *helloworldv1.HelloWorkflowInput) (helloworldv1.HelloWorkflow, error) {\n  return &HelloWorkflow{input, workflow.GetLogger(ctx)}, nil\n}\n\n// Execute defines the entrypoint to a Hello workflow\nfunc (w *HelloWorkflow) Execute(ctx workflow.Context) (*helloworldv1.HelloResponse, error) {\n  w.log.Info("Hello workflow started", "request", w.Req)\n\n  goodbye, _ := w.Goodbye.Receive(ctx)\n  w.log.Info("Goodbye received", "signal", goodbye)\n\n  return &helloworldv1.HelloResponse{}, nil\n}\n    '}]},{label:"Run",content:"Run your Temporal Worker using the generated helpers.",fragments:[{language:"go",title:"main.go",content:'package main\n\nimport (\n  "log"\n  "os"\n\n  example "internal"\n  helloworldv1 "path/to/gen/helloworld/v1"\n  "github.com/urfave/cli/v2"\n  "go.temporal.io/sdk/client"\n  "go.temporal.io/sdk/worker"\n)\n\nfunc main() {\n  app, err := helloworldv1.NewExampleCli(\n    helloworldv1.NewExampleCliOptions().\n      WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {\n        w := worker.New(c, helloworldv1.ExampleTaskQueue, worker.Options{})\n        helloworldv1.RegisterExampleWorkflows(w, &example.Workflows{})\n        return w, nil\n      }),\n  )\n  if err != nil {\n    log.Fatal(err)\n  }\n\n  if err := app.Run(os.Args); err != nil {\n    log.Fatal(err)\n  }\n}\n    '}]},{label:"Client",content:"Interact with your workers from any Go application using the generated Client.",fragments:[{language:"go",title:"cmd/client/main.go",content:'package main\n\nimport (\n  "context"\n  "log"\n  "log/slog"\n  "os"\n  "os/signal"\n  "syscall"\n\n  helloworldv1 "path/to/gen/helloworld/v1"\n  "go.temporal.io/sdk/client"\n  sdklog "go.temporal.io/sdk/log"\n)\n\nfunc main() {\n  ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)\n  defer cancel()\n  logger := slog.New(slog.NewTextHandler(os.Stdout, nil))\n\n  client, err := client.Dial(client.Options{\n    Logger: sdklog.NewStructuredLogger(logger),\n  })\n  if err != nil {\n    log.Fatal(err)\n  }\n  defer client.Close()\n\n  example := helloworldv1.NewExampleClient(client)\n  run, err := example.HelloAsync(ctx, &helloworldv1.HelloRequest{Name: "World"})\n  if err != nil {\n    log.Fatal(err)\n  }\n  logger = logger.With("workflow_id", run.ID())\n  logger.Info("workflow started")\n\n  _, ctx = <-ctx.Done(), context.Background()\n  logger.Info("received shutdown signal, sending Goodbye signal to workflow")\n  if err := run.Goodbye(ctx, &helloworldv1.GoodbyeRequest{}); err != nil {\n    log.Fatal(err)\n  }\n\n  out, err := run.Get(ctx)\n  if err != nil {\n    log.Fatal(err)\n  }\n  logger.Info("workflow completed", "result", out.String())\n}    \n    '}]},{label:"CLI",content:"Or from your local machine using the generated Command Line Interface.",fragments:[{language:"sh",title:"print cli usage",content:"go run main.go -h\nNAME:\n    example - A new cli application\n\nUSAGE:\n    example [global options] command [command options] [arguments...]\n\nCOMMANDS:\n    worker   runs a example.helloworld.v1.Example worker process\n    help, h  Shows a list of commands or help for one command\n    SIGNALS:\n      goodbye  Goodbye signals a running workflow to exit\n    WORKFLOWS:\n      hello  Hello prints a friendly greeting and waits for goodbye\n\nGLOBAL OPTIONS:\n    --help, -h  show help\n    "},{language:"sh",title:"start a workflow",content:"go run main.go hello --name Temporal -d\nsuccess\nworkflow id: hello/Temporal\nrun id: e55c6b09-7d05-418e-ad7e-8b40b9b3b867\n    "},{language:"sh",title:"send a signal",content:"go run main.go goodbye --message \ud83d\udc4b -d\nsuccess\n    "}]},{label:"XNS",content:"Or from other Temporal workflows in a different Namespace or Cluster.",fragments:[{language:"go",title:"main.go",content:'package main\n\nimport (\n  "time"\n\n  helloworldv1 "path/to/gen/helloworld/v1"\n  "path/to/gen/helloworld/v1/helloworldv1xns"\n  "go.temporal.io/sdk/client"\n  "go.temporal.io/sdk/worker"\n  "go.temporal.io/sdk/workflow"\n)\n\nfunc SomeOtherWorkflow(ctx workflow.Context) error {\n  run, err := helloworldv1xns.HelloAsync(ctx, &helloworldv1.HelloRequest{\n    Name: workflow.GetInfo(ctx).WorkflowExecution.ID,\n  })\n  if err != nil {\n    return err\n  }\n\n  workflow.Sleep(ctx, time.Second*30)\n  if err := run.Goodbye(ctx, &helloworldv1.GoodbyeRequest{}); err != nil {\n    return err\n  }\n\n  _, err = run.Get(ctx)\n  return err\n}\n\nfunc main() {\n  c, _ := client.Dial(client.Options{})\n  defer c.Close()\n\n  // initialize client for a different namespace/cluster\n  xnsc, _ := client.NewClientFromExisting(c, client.Options{Namespace: "helloworld"})\n\n  w := worker.New(c, "my-task-queue", worker.Options{})\n  w.RegisterWorkflow(SomeOtherWorkflow)\n  helloworldv1xns.RegisterExampleActivities(w, helloworldv1.NewExampleClient(xnsc))\n  w.Run(w.InterruptCh())\n}\n        '}]}];function x(){return(0,g.jsx)("section",{className:c.features,children:(0,g.jsx)("div",{className:"container",children:(0,g.jsxs)("div",{className:"row",children:[(0,g.jsx)("div",{className:"col col--6",children:(0,g.jsx)(m,{})}),(0,g.jsx)("div",{className:"col col--6",children:(0,g.jsx)(u.A,{children:w.map(((e,n)=>(0,g.jsxs)(d.A,{value:n.toString(),label:e.label,children:[(0,g.jsx)("p",{children:e.content}),e.fragments.map(((e,n)=>(0,g.jsx)(p.A,{language:e.language,title:e.title,children:e.content})))]})))})})]})})})}const b={heroBanner:"heroBanner_qdFl",buttons:"buttons_AeoN"};function v(){const{siteConfig:e}=(0,s.A)();return(0,g.jsx)("header",{className:(0,r.A)("hero hero--primary",b.heroBanner),children:(0,g.jsx)("div",{className:"container",children:(0,g.jsxs)("div",{className:"row",children:[(0,g.jsxs)("div",{className:t()("col col--6"),children:[(0,g.jsx)("h1",{className:"hero__title",children:e.title}),(0,g.jsx)("p",{className:"hero__subtitle",children:e.tagline}),(0,g.jsx)("div",{className:b.buttons,children:(0,g.jsx)(a.A,{className:"button button--secondary button--lg",to:"/docs/about",children:"Get Started"})})]}),(0,g.jsx)("div",{className:t()("col col--6"),children:(0,g.jsx)("img",{width:"300",className:b.heroImg,src:"img/logo.png"})})]})})})}function k(){const{siteConfig:e}=(0,s.A)();return(0,g.jsxs)(i.A,{title:`Hello from ${e.title}`,description:"Description will go into a meta tag in <head />",children:[(0,g.jsx)(v,{}),(0,g.jsx)("main",{children:(0,g.jsx)(x,{})})]})}},6942:(e,n)=>{var o;!function(){"use strict";var r={}.hasOwnProperty;function l(){for(var e="",n=0;n<arguments.length;n++){var o=arguments[n];o&&(e=a(e,t(o)))}return e}function t(e){if("string"==typeof e||"number"==typeof e)return e;if("object"!=typeof e)return"";if(Array.isArray(e))return l.apply(null,e);if(e.toString!==Object.prototype.toString&&!e.toString.toString().includes("[native code]"))return e.toString();var n="";for(var o in e)r.call(e,o)&&e[o]&&(n=a(n,o));return n}function a(e,n){return n?e?e+" "+n:e+n:e}e.exports?(l.default=l,e.exports=l):void 0===(o=function(){return l}.apply(n,[]))||(e.exports=o)}()}}]);