"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[383],{5514:(e,t,o)=>{o.r(t),o.d(t,{assets:()=>v,contentTitle:()=>x,default:()=>F,frontMatter:()=>w,metadata:()=>h,toc:()=>k});var r=o(4848),n=o(8453),s=o(1432),a=o(1470),l=o(9365),i=o(3595);const u="version: v1\ndeps:\n  - buf.build/cludden/protoc-gen-go-temporal\nbreaking:\n  use:\n    - FILE\nlint:\n  allow_comment_ignores: true\n  use:\n    - BASIC\n",c="version: v1\nmanaged:\n  enabled: true\n  go_package_prefix:\n    default: example/gen\n    except:\n      - buf.build/cludden/protoc-gen-go-temporal\nplugins:\n  - plugin: go\n    out: gen\n    opt: paths=source_relative\n  - plugin: go_temporal\n    out: gen\n    opt: paths=source_relative,cli-enabled=true,cli-categories=true,workflow-update-enabled=true,docs-out=./proto/README.md\n    strategy: all\n";var p=o(7429),g=o(3238),m=o(2397);let f=[{language:"sh",title:"build cli binary",output:"",content:"go build -o example cmd/example/main.go"},{language:"sh",title:"print cli usage details",output:"img/cli-usage.png",content:"example -h"},{language:"sh",title:"start a workflow",output:"img/cli-start-workflow.png",content:"example create-foo --name test -d"},{language:"sh",title:"send a signal",output:"",content:"example set-foo-progress -w create-foo/test --progress 5.7"},{language:"sh",title:"query workflow",output:"img/cli-query.png",content:"example get-foo-progress -w create-foo/test"},{language:"sh",title:"update workflow",output:"img/cli-update.png",content:"example update-foo-progress -w create-foo/test --progress 100"}],d='package xns\n\nimport (\n  "fmt"\n\n  examplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"\n  "github.com/cludden/protoc-gen-go-temporal/gen/example/v1/examplev1xns"\n  "go.temporal.io/sdk/workflow"\n)\n\nfunc SomeWorkflow(ctx workflow.Context) error {\n  log := workflow.GetLogger(ctx)\n\n  run, err := examplev1xns.CreateFooAsync(ctx, &examplev1.CreateFooRequest{Name: w.Req.GetName()})\n\tif err != nil {\n\t\treturn fmt.Errorf("error initializing CreateFoo workflow: %w", err)\n\t}\n\n\tif err := run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7}); err != nil {\n\t\treturn fmt.Errorf("error signaling SetFooProgress: %w", err)\n\t}\n\tlog.Info("SetFooProgress", "progress", 5.7)\n\n\tprogress, err := run.GetFooProgress(ctx)\n\tif err != nil {\n\t\treturn fmt.Errorf("error querying GetFooProgress: %w", err)\n\t}\n\tlog.Info("GetFooProgress", "status", progress.GetStatus().String(), "progress", progress.GetProgress())\n\n\tupdate, err := run.UpdateFooProgressAsync(ctx, &examplev1.SetFooProgressRequest{Progress: 100})\n\tif err != nil {\n\t\treturn fmt.Errorf("error initializing UpdateFooProgress: %w", err)\n\t}\n\tprogress, err = update.Get(ctx)\n\tif err != nil {\n\t\treturn fmt.Errorf("error updating UpdateFooProgress: %w", err)\n\t}\n\tlog.Info("UpdateFooProgress", "status", progress.GetStatus().String(), "progress", progress.GetProgress())\n\n\tresp, err := run.Get(ctx)\n\tif err != nil {\n\t\treturn err\n\t}\n\treturn nil\n}\n';const w={},x=void 0,h={id:"how-it-works",title:"how-it-works",description:"Annotate your protobuf services and methods with Temporal options.",source:"@site/docs/how-it-works.mdx",sourceDirName:".",slug:"/how-it-works",permalink:"/protoc-gen-go-temporal/docs/how-it-works",draft:!1,unlisted:!1,editUrl:"https://github.com/cludden/protoc-gen-go-temporal/tree/main/docs/docs/how-it-works.mdx",tags:[],version:"current",frontMatter:{}},v={},k=[];function b(e){const t={img:"img",p:"p",section:"section",...(0,n.R)(),...e.components};return(0,r.jsxs)(a.A,{children:[(0,r.jsxs)(l.A,{value:"annotate",label:"Annotate",children:[(0,r.jsx)(t.p,{children:"Annotate your protobuf services and methods with Temporal options."}),(0,r.jsx)(s.A,{language:"protobuf",title:"proto/example/v1/example.proto",children:i.A})]}),(0,r.jsxs)(l.A,{value:"generate",label:"Generate",children:[(0,r.jsx)(t.p,{children:"Generate Go code for implementing Temporal Clients, Workers, and CLI applications."}),(0,r.jsx)(s.A,{language:"yaml",title:"buf.yaml",children:u}),(0,r.jsx)(s.A,{language:"yaml",title:"buf.gen.yaml",children:c}),(0,r.jsx)(s.A,{language:"sh",children:"buf generate"})]}),(0,r.jsxs)(l.A,{value:"implement",label:"Implement",children:[(0,r.jsx)(t.p,{children:"Implement the required Workflow and Activity interfaces."}),(0,r.jsx)(s.A,{language:"go",title:"example.go",children:p.A})]}),(0,r.jsxs)(l.A,{value:"run",label:"Run",children:[(0,r.jsx)(t.p,{children:"Run your Temporal Worker using the generated helpers."}),(0,r.jsx)(s.A,{language:"go",title:"cmd/example/main.go",children:g.A}),(0,r.jsx)(s.A,{language:"sh",children:"go run cmd/example/main.go worker"})]}),(0,r.jsxs)(l.A,{value:"client",label:"Client",children:[(0,r.jsx)(t.p,{children:"Interact with your workers from any Go application using the generated Client."}),(0,r.jsx)(s.A,{language:"go",title:"cmd/client/main.go",children:m.A}),(0,r.jsx)(s.A,{language:"sh",children:"go run cmd/client/main.go worker"})]}),(0,r.jsxs)(l.A,{value:"cli",label:"CLI",children:[(0,r.jsx)(t.p,{children:"Or from your local machine using the generated Command Line Interface."}),f.map(((e,o)=>(0,r.jsxs)(t.section,{children:[(0,r.jsx)(s.A,{language:e.language,title:e.title,children:e.content}),""!=e.output&&(0,r.jsx)(t.img,{style:{"margin-bottom":"20px"},src:e.output})]})))]}),(0,r.jsxs)(l.A,{value:"xns",label:"XNS",children:[(0,r.jsx)(t.p,{children:"Or from other Temporal workflows in a different Namespace or Cluster."}),(0,r.jsx)(s.A,{language:"go",title:"xns.go",children:d})]})]})}function F(e={}){const{wrapper:t}={...(0,n.R)(),...e.components};return t?(0,r.jsx)(t,{...e,children:(0,r.jsx)(b,{...e})}):b(e)}},9365:(e,t,o)=>{o.d(t,{A:()=>a});o(6540);var r=o(4164);const n={tabItem:"tabItem_Ymn6"};var s=o(4848);function a(e){let{children:t,hidden:o,className:a}=e;return(0,s.jsx)("div",{role:"tabpanel",className:(0,r.A)(n.tabItem,a),hidden:o,children:t})}},1470:(e,t,o)=>{o.d(t,{A:()=>F});var r=o(6540),n=o(4164),s=o(3104),a=o(6347),l=o(205),i=o(7485),u=o(1682),c=o(9466);function p(e){return r.Children.toArray(e).filter((e=>"\n"!==e)).map((e=>{if(!e||(0,r.isValidElement)(e)&&function(e){const{props:t}=e;return!!t&&"object"==typeof t&&"value"in t}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))?.filter(Boolean)??[]}function g(e){const{values:t,children:o}=e;return(0,r.useMemo)((()=>{const e=t??function(e){return p(e).map((e=>{let{props:{value:t,label:o,attributes:r,default:n}}=e;return{value:t,label:o,attributes:r,default:n}}))}(o);return function(e){const t=(0,u.X)(e,((e,t)=>e.value===t.value));if(t.length>0)throw new Error(`Docusaurus error: Duplicate values "${t.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[t,o])}function m(e){let{value:t,tabValues:o}=e;return o.some((e=>e.value===t))}function f(e){let{queryString:t=!1,groupId:o}=e;const n=(0,a.W6)(),s=function(e){let{queryString:t=!1,groupId:o}=e;if("string"==typeof t)return t;if(!1===t)return null;if(!0===t&&!o)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return o??null}({queryString:t,groupId:o});return[(0,i.aZ)(s),(0,r.useCallback)((e=>{if(!s)return;const t=new URLSearchParams(n.location.search);t.set(s,e),n.replace({...n.location,search:t.toString()})}),[s,n])]}function d(e){const{defaultValue:t,queryString:o=!1,groupId:n}=e,s=g(e),[a,i]=(0,r.useState)((()=>function(e){let{defaultValue:t,tabValues:o}=e;if(0===o.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(t){if(!m({value:t,tabValues:o}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${t}" but none of its children has the corresponding value. Available values are: ${o.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return t}const r=o.find((e=>e.default))??o[0];if(!r)throw new Error("Unexpected error: 0 tabValues");return r.value}({defaultValue:t,tabValues:s}))),[u,p]=f({queryString:o,groupId:n}),[d,w]=function(e){let{groupId:t}=e;const o=function(e){return e?`docusaurus.tab.${e}`:null}(t),[n,s]=(0,c.Dv)(o);return[n,(0,r.useCallback)((e=>{o&&s.set(e)}),[o,s])]}({groupId:n}),x=(()=>{const e=u??d;return m({value:e,tabValues:s})?e:null})();(0,l.A)((()=>{x&&i(x)}),[x]);return{selectedValue:a,selectValue:(0,r.useCallback)((e=>{if(!m({value:e,tabValues:s}))throw new Error(`Can't select invalid tab value=${e}`);i(e),p(e),w(e)}),[p,w,s]),tabValues:s}}var w=o(2303);const x={tabList:"tabList__CuJ",tabItem:"tabItem_LNqP"};var h=o(4848);function v(e){let{className:t,block:o,selectedValue:r,selectValue:a,tabValues:l}=e;const i=[],{blockElementScrollPositionUntilNextRender:u}=(0,s.a_)(),c=e=>{const t=e.currentTarget,o=i.indexOf(t),n=l[o].value;n!==r&&(u(t),a(n))},p=e=>{let t=null;switch(e.key){case"Enter":c(e);break;case"ArrowRight":{const o=i.indexOf(e.currentTarget)+1;t=i[o]??i[0];break}case"ArrowLeft":{const o=i.indexOf(e.currentTarget)-1;t=i[o]??i[i.length-1];break}}t?.focus()};return(0,h.jsx)("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,n.A)("tabs",{"tabs--block":o},t),children:l.map((e=>{let{value:t,label:o,attributes:s}=e;return(0,h.jsx)("li",{role:"tab",tabIndex:r===t?0:-1,"aria-selected":r===t,ref:e=>i.push(e),onKeyDown:p,onClick:c,...s,className:(0,n.A)("tabs__item",x.tabItem,s?.className,{"tabs__item--active":r===t}),children:o??t},t)}))})}function k(e){let{lazy:t,children:o,selectedValue:n}=e;const s=(Array.isArray(o)?o:[o]).filter(Boolean);if(t){const e=s.find((e=>e.props.value===n));return e?(0,r.cloneElement)(e,{className:"margin-top--md"}):null}return(0,h.jsx)("div",{className:"margin-top--md",children:s.map(((e,t)=>(0,r.cloneElement)(e,{key:t,hidden:e.props.value!==n})))})}function b(e){const t=d(e);return(0,h.jsxs)("div",{className:(0,n.A)("tabs-container",x.tabList),children:[(0,h.jsx)(v,{...e,...t}),(0,h.jsx)(k,{...e,...t})]})}function F(e){const t=(0,w.A)();return(0,h.jsx)(b,{...e,children:p(e.children)},String(t))}},2397:(e,t,o)=>{o.d(t,{A:()=>r});const r='package main\n\nimport (\n\t"context"\n\t"log"\n\n\texamplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"\n\t"go.temporal.io/sdk/client"\n)\n\nfunc main() {\n\t// initialize service client with sdk client\n\tc, _ := client.Dial(client.Options{})\n\tclient, ctx := examplev1.NewExampleClient(c), context.Background()\n\n\t// execute a workflow asynchronously\n\trun, _ := client.CreateFooAsync(ctx, &examplev1.CreateFooRequest{Name: "test"})\n\tlog.Printf("started workflow: workflow_id=%s, run_id=%s\\n", run.ID(), run.RunID())\n\n\t// send a signal to the workflow\n\tlog.Println("signalling progress")\n\t_ = run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7})\n\n\t// query the workflow\n\tprogress, _ := run.GetFooProgress(ctx)\n\tlog.Printf("queried progress: %s\\n", progress.String())\n\n\t// update the workflow\n\tupdate, _ := run.UpdateFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 100})\n\tlog.Printf("updated progress: %s\\n", update.String())\n\n\t// block on workflow completion\n\tresp, _ := run.Get(ctx)\n\tlog.Printf("workflow completed: %s\\n", resp.String())\n}\n'},3238:(e,t,o)=>{o.d(t,{A:()=>r});const r='package main\n\nimport (\n\t"log"\n\t"log/slog"\n\t"os"\n\n\t"github.com/cludden/protoc-gen-go-temporal/examples/example"\n\texamplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"\n\t"github.com/urfave/cli/v2"\n\t"go.temporal.io/sdk/client"\n\tlogsdk "go.temporal.io/sdk/log"\n\t"go.temporal.io/sdk/worker"\n)\n\nfunc main() {\n\t// initialize the generated cli application\n\tapp, err := examplev1.NewExampleCli(\n\t\texamplev1.NewExampleCliOptions().\n\t\t\tWithClient(func(cmd *cli.Context) (client.Client, error) {\n\t\t\t\treturn client.Dial(client.Options{\n\t\t\t\t\tLogger: logsdk.NewStructuredLogger(slog.New(slog.NewTextHandler(os.Stdout, nil))),\n\t\t\t\t})\n\t\t\t}).\n\t\t\tWithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {\n\t\t\t\t// register activities and workflows using generated helpers\n\t\t\t\tw := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})\n\t\t\t\texamplev1.RegisterExampleActivities(w, &example.Activities{})\n\t\t\t\texamplev1.RegisterExampleWorkflows(w, &example.Workflows{})\n\t\t\t\treturn w, nil\n\t\t\t}),\n\t)\n\tif err != nil {\n\t\tlog.Fatalf("error initializing example cli: %v", err)\n\t}\n\n\t// run cli\n\tif err := app.Run(os.Args); err != nil {\n\t\tlog.Fatal(err)\n\t}\n}\n'},7429:(e,t,o)=>{o.d(t,{A:()=>r});const r='package example\n\nimport (\n\t"context"\n\t"fmt"\n\n\texamplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"\n\t"go.temporal.io/sdk/activity"\n\t"go.temporal.io/sdk/log"\n\t"go.temporal.io/sdk/workflow"\n)\n\ntype (\n\t// Workflows manages shared state for workflow constructors and is used to\n\t// register workflows with a worker\n\tWorkflows struct{}\n\n\t// Activities manages shared state for activities and is used to register\n\t// activities with a worker\n\tActivities struct{}\n\n\t// CreateFooWorkflow manages workflow state for a CreateFoo workflow\n\tCreateFooWorkflow struct {\n\t\t// it embeds the generated workflow Input type that contains the workflow\n\t\t// input and signal helpers\n\t\t*examplev1.CreateFooWorkflowInput\n\n\t\tlog      log.Logger\n\t\tprogress float32\n\t\tstatus   examplev1.Foo_Status\n\t}\n)\n\n// CreateFoo initializes a new examplev1.CreateFooWorkflow value\nfunc (w *Workflows) CreateFoo(ctx workflow.Context, input *examplev1.CreateFooWorkflowInput) (examplev1.CreateFooWorkflow, error) {\n\treturn &CreateFooWorkflow{\n\t\tCreateFooWorkflowInput: input,\n\t\tlog:                    workflow.GetLogger(ctx),\n\t\tstatus:                 examplev1.Foo_FOO_STATUS_CREATING,\n\t}, nil\n}\n\n// Execute defines the entrypoint to a example.v1.Example.CreateFoo workflow\nfunc (wf *CreateFooWorkflow) Execute(ctx workflow.Context) (*examplev1.CreateFooResponse, error) {\n\t// listen for signals using generated signal provided by workflow input\n\tworkflow.Go(ctx, func(ctx workflow.Context) {\n\t\tfor {\n\t\t\tsignal, _ := wf.SetFooProgress.Receive(ctx)\n\t\t\twf.UpdateFooProgress(ctx, signal)\n\t\t}\n\t})\n\n\t// execute Notify activity using generated helper\n\tif err := examplev1.Notify(ctx, &examplev1.NotifyRequest{\n\t\tMessage: fmt.Sprintf("creating foo resource (%s)", wf.Req.GetName()),\n\t}); err != nil {\n\t\treturn nil, fmt.Errorf("error sending notification: %w", err)\n\t}\n\n\t// block until progress has reached 100 via signals and/or updates\n\tif err := workflow.Await(ctx, func() bool {\n\t\treturn wf.status == examplev1.Foo_FOO_STATUS_READY\n\t}); err != nil {\n\t\treturn nil, fmt.Errorf("error awaiting ready status: %w", err)\n\t}\n\n\treturn &examplev1.CreateFooResponse{\n\t\tFoo: &examplev1.Foo{\n\t\t\tName:   wf.Req.GetName(),\n\t\t\tStatus: wf.status,\n\t\t},\n\t}, nil\n}\n\n// GetFooProgress defines the handler for a GetFooProgress query\nfunc (wf *CreateFooWorkflow) GetFooProgress() (*examplev1.GetFooProgressResponse, error) {\n\treturn &examplev1.GetFooProgressResponse{Progress: wf.progress, Status: wf.status}, nil\n}\n\n// UpdateFooProgress defines the handler for an UpdateFooProgress update\nfunc (wf *CreateFooWorkflow) UpdateFooProgress(ctx workflow.Context, req *examplev1.SetFooProgressRequest) (*examplev1.GetFooProgressResponse, error) {\n\twf.progress = req.GetProgress()\n\tswitch {\n\tcase wf.progress < 0:\n\t\twf.progress, wf.status = 0, examplev1.Foo_FOO_STATUS_CREATING\n\tcase wf.progress < 100:\n\t\twf.status = examplev1.Foo_FOO_STATUS_CREATING\n\tcase wf.progress >= 100:\n\t\twf.progress, wf.status = 100, examplev1.Foo_FOO_STATUS_READY\n\t}\n\treturn &examplev1.GetFooProgressResponse{Progress: wf.progress, Status: wf.status}, nil\n}\n\n// Notify defines the implementation for a Notify activity\nfunc (a *Activities) Notify(ctx context.Context, req *examplev1.NotifyRequest) error {\n\tactivity.GetLogger(ctx).Info("notification", "message", req.GetMessage())\n\treturn nil\n}\n'},3595:(e,t,o)=>{o.d(t,{A:()=>r});const r='syntax="proto3";\n\npackage example.v1;\n\nimport "google/protobuf/empty.proto";\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  };\n\n  // CreateFoo creates a new foo operation\n  rpc CreateFoo(CreateFooRequest) returns (CreateFooResponse) {\n    option (temporal.v1.workflow) = {\n      execution_timeout: { seconds: 3600 }\n      id: \'create-foo/${! name.slug() }\'\n      id_reuse_policy: WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE\n      query: { ref: "GetFooProgress" }\n      signal: { ref: "SetFooProgress", start: true }\n      update: { ref: "UpdateFooProgress" }\n    };\n  }\n\n  // GetFooProgress returns the status of a CreateFoo operation\n  rpc GetFooProgress(google.protobuf.Empty) returns (GetFooProgressResponse) {\n    option (temporal.v1.query) = {};\n  }\n\n  // Notify sends a notification\n  rpc Notify(NotifyRequest) returns (google.protobuf.Empty) {\n    option (temporal.v1.activity) = {\n      start_to_close_timeout: { seconds: 30 }\n      retry_policy: {\n        max_attempts: 3\n      }\n    };\n  }\n\n  // SetFooProgress sets the current status of a CreateFoo operation\n  rpc SetFooProgress(SetFooProgressRequest) returns (google.protobuf.Empty) {\n    option (temporal.v1.signal) = {};\n  }\n\n  // UpdateFooProgress sets the current status of a CreateFoo operation\n  rpc UpdateFooProgress(SetFooProgressRequest) returns (GetFooProgressResponse) {\n    option (temporal.v1.update) = {\n      id: \'update-progress/${! progress.string() }\'\n    };\n  }\n}\n\n// CreateFooRequest describes the input to a CreateFoo workflow\nmessage CreateFooRequest {\n  // unique foo name\n  string name = 1;\n}\n\n// SampleWorkflowWithMutexResponse describes the output from a CreateFoo workflow\nmessage CreateFooResponse {\n  Foo foo = 1; \n}\n\n// Foo describes an illustrative foo resource\nmessage Foo {\n  string name = 1;\n  Status status = 2;\n\n  enum Status {\n    FOO_STATUS_UNSPECIFIED = 0;\n    FOO_STATUS_READY = 1;\n    FOO_STATUS_CREATING = 2;\n  }\n}\n\n// GetFooProgressResponse describes the output from a GetFooProgress query\nmessage GetFooProgressResponse {\n  float progress = 1;\n  Foo.Status status = 2;\n}\n\n// NotifyRequest describes the input to a Notify activity\nmessage NotifyRequest {\n  string message = 1;\n}\n\n// SetFooProgressRequest describes the input to a SetFooProgress signal\nmessage SetFooProgressRequest {\n  // value of current workflow progress\n  float progress = 1;\n}\n'}}]);