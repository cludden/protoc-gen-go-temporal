"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[74],{9943:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>m,contentTitle:()=>p,default:()=>d,frontMatter:()=>l,metadata:()=>c,toc:()=>u});var o=t(4848),r=t(8453),s=t(1432);const a="syntax=\"proto3\";\n\npackage example.xns.v1;\n\nimport \"google/protobuf/empty.proto\";\nimport \"temporal/v1/temporal.proto\";\n\nservice Xns {\n  option (temporal.v1.service) = {\n    task_queue: \"xns-v1\"\n  };\n\n  rpc ProvisionFoo(ProvisionFooRequest) returns (ProvisionFooResponse) {\n    option (temporal.v1.workflow) = {\n      id: 'provision-foo/${! name.slug() }'\n    };\n  }\n}\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: \"example-v1\"\n  };\n\n  // CreateFoo creates a new foo operation\n  rpc CreateFoo(CreateFooRequest) returns (CreateFooResponse) {\n    option (temporal.v1.workflow) = {\n      execution_timeout: { seconds: 3600 }\n      id_reuse_policy: WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE\n      id: 'create-foo/${! name.slug() }'\n      xns: {\n        heartbeat_interval: { seconds: 10 }\n        heartbeat_timeout: { seconds: 20 }\n        start_to_close_timeout: { seconds: 3630 }\n      }\n      query: { ref: 'GetFooProgress' }\n      signal: { ref: 'SetFooProgress', start: true }\n      update: { ref: 'UpdateFooProgress' }\n    };\n  }\n\n  // GetFooProgress returns the status of a CreateFoo operation\n  rpc GetFooProgress(google.protobuf.Empty) returns (GetFooProgressResponse) {\n    option (temporal.v1.query) = {\n      xns: {\n        heartbeat_interval: { seconds: 10 }\n        heartbeat_timeout: { seconds: 20 }\n        start_to_close_timeout: { seconds: 60 }\n      }\n    };\n  }\n\n  // Notify sends a notification\n  rpc Notify(NotifyRequest) returns (google.protobuf.Empty) {\n    option (temporal.v1.activity) = {\n      start_to_close_timeout: { seconds: 30 }\n      retry_policy: {\n        max_attempts: 3\n      }\n    };\n  }\n\n  // SetFooProgress sets the current status of a CreateFoo operation\n  rpc SetFooProgress(SetFooProgressRequest) returns (google.protobuf.Empty) {\n    option (temporal.v1.signal) = {\n      xns: {\n        heartbeat_interval: { seconds: 10 }\n        heartbeat_timeout: { seconds: 20 }\n        start_to_close_timeout: { seconds: 60 }\n      }\n    };\n  }\n\n  // UpdateFooProgress sets the current status of a CreateFoo operation\n  rpc UpdateFooProgress(SetFooProgressRequest) returns (GetFooProgressResponse) {\n    option (temporal.v1.update) = {\n      id: 'update-progress/${! progress.string() }',\n      xns: {\n        heartbeat_interval: { seconds: 10 }\n        heartbeat_timeout: { seconds: 20 }\n        start_to_close_timeout: { seconds: 60 }\n      }\n    };\n  }\n}\n\n// CreateFooRequest describes the input to a CreateFoo workflow\nmessage CreateFooRequest {\n  // unique foo name\n  string name = 1;\n}\n\n// SampleWorkflowWithMutexResponse describes the output from a CreateFoo workflow\nmessage CreateFooResponse {\n  Foo foo = 1; \n}\n\n// Foo describes an illustrative foo resource\nmessage Foo {\n  string name = 1;\n  Status status = 2;\n\n  enum Status {\n    FOO_STATUS_UNSPECIFIED = 0;\n    FOO_STATUS_READY = 1;\n    FOO_STATUS_CREATING = 2;\n  }\n}\n\n// GetFooProgressResponse describes the output from a GetFooProgress query\nmessage GetFooProgressResponse {\n  float progress = 1;\n  Foo.Status status = 2;\n}\n\n// NotifyRequest describes the input to a Notify activity\nmessage NotifyRequest {\n  string message = 1;\n}\n\n// ProvisionFooRequest describes the input to a ProvisionFoo workflow\nmessage ProvisionFooRequest {\n  // unique foo name\n  string name = 1;\n}\n\n// SampleWorkflowWithMutexResponse describes the output from a ProvisionFoo workflow\nmessage ProvisionFooResponse {\n  Foo foo = 1; \n}\n\n// SetFooProgressRequest describes the input to a SetFooProgress signal\nmessage SetFooProgressRequest {\n  // value of current workflow progress\n  float progress = 1;\n}\n",i='package main\n\nimport (\n\t"fmt"\n\t"log"\n\t"os"\n\t"os/signal"\n\t"sync"\n\t"syscall"\n\n\txnsv1 "github.com/cludden/protoc-gen-go-temporal/gen/example/xns/v1"\n\t"github.com/cludden/protoc-gen-go-temporal/gen/example/xns/v1/xnsv1xns"\n\t"github.com/urfave/cli/v2"\n\t"go.temporal.io/sdk/client"\n\ttlog "go.temporal.io/sdk/log"\n\t"go.temporal.io/sdk/worker"\n\t"go.temporal.io/sdk/workflow"\n)\n\ntype (\n\tWorkflows struct{}\n)\n\ntype ProvisionFooWorkflow struct {\n\t*xnsv1.ProvisionFooWorkflowInput\n\tlog tlog.Logger\n}\n\nfunc (wfs *Workflows) ProvisionFoo(ctx workflow.Context, input *xnsv1.ProvisionFooWorkflowInput) (xnsv1.ProvisionFooWorkflow, error) {\n\treturn &ProvisionFooWorkflow{input, workflow.GetLogger(ctx)}, nil\n}\n\nfunc (w *ProvisionFooWorkflow) Execute(ctx workflow.Context) (*xnsv1.ProvisionFooResponse, error) {\n\trun, err := xnsv1xns.CreateFooAsync(ctx, &xnsv1.CreateFooRequest{Name: w.Req.GetName()})\n\tif err != nil {\n\t\treturn nil, fmt.Errorf("error initializing CreateFoo workflow: %w", err)\n\t}\n\n\tif err := run.SetFooProgress(ctx, &xnsv1.SetFooProgressRequest{Progress: 5.7}); err != nil {\n\t\treturn nil, fmt.Errorf("error signaling SetFooProgress: %w", err)\n\t}\n\tw.log.Info("SetFooProgress", "progress", 5.7)\n\n\tprogress, err := run.GetFooProgress(ctx)\n\tif err != nil {\n\t\treturn nil, fmt.Errorf("error querying GetFooProgress: %w", err)\n\t}\n\tw.log.Info("GetFooProgress", "status", progress.GetStatus().String(), "progress", progress.GetProgress())\n\n\tupdate, err := run.UpdateFooProgressAsync(ctx, &xnsv1.SetFooProgressRequest{Progress: 100})\n\tif err != nil {\n\t\treturn nil, fmt.Errorf("error initializing UpdateFooProgress: %w", err)\n\t}\n\tprogress, err = update.Get(ctx)\n\tif err != nil {\n\t\treturn nil, fmt.Errorf("error updating UpdateFooProgress: %w", err)\n\t}\n\tw.log.Info("UpdateFooProgress", "status", progress.GetStatus().String(), "progress", progress.GetProgress())\n\n\tresp, err := run.Get(ctx)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn &xnsv1.ProvisionFooResponse{Foo: resp.GetFoo()}, nil\n}\n\nfunc main() {\n\tapp := &cli.App{}\n\n\texampleCmd, err := xnsv1.NewExampleCliCommand()\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n\tapp.Commands = append(app.Commands, exampleCmd)\n\n\txnsCmd, err := xnsv1.NewXnsCliCommand()\n\tif err != nil {\n\t\tlog.Fatal(err)\n\t}\n\tapp.Commands = append(app.Commands, xnsCmd)\n\n\tapp.Commands = append(app.Commands, &cli.Command{\n\t\tName: "worker",\n\t\tAction: func(cmd *cli.Context) error {\n\t\t\tc, err := client.Dial(client.Options{\n\t\t\t\tNamespace: "example",\n\t\t\t})\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\t\t\tdefer c.Close()\n\n\t\t\txnsc, err := client.NewClientFromExisting(c, client.Options{\n\t\t\t\tNamespace: "default",\n\t\t\t})\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n\n\t\t\texamplew := worker.New(c, xnsv1.ExampleTaskQueue, worker.Options{})\n\t\t\txnsv1.RegisterExampleWorkflows(examplew, &ExampleWorkflows{})\n\t\t\txnsv1.RegisterExampleActivities(examplew, &ExampleActivities{})\n\n\t\t\txnsw := worker.New(xnsc, xnsv1.XnsTaskQueue, worker.Options{})\n\t\t\txnsv1.RegisterXnsWorkflows(xnsw, &Workflows{})\n\t\t\txnsv1xns.RegisterExampleActivities(xnsw, xnsv1.NewExampleClient(c))\n\n\t\t\tvar g sync.WaitGroup\n\t\t\tcloseCh := make(chan any)\n\t\t\tg.Add(2)\n\t\t\tgo func() {\n\t\t\t\tdefer g.Done()\n\t\t\t\texamplew.Run(closeCh)\n\t\t\t}()\n\t\t\tgo func() {\n\t\t\t\tdefer g.Done()\n\t\t\t\txnsw.Run(closeCh)\n\t\t\t}()\n\n\t\t\tinterruptCh := make(chan os.Signal, 1)\n\t\t\tsignal.Notify(interruptCh, syscall.SIGINT, syscall.SIGTERM)\n\t\t\tgo func() {\n\t\t\t\t<-interruptCh\n\t\t\t\tclose(closeCh)\n\t\t\t}()\n\t\t\tg.Wait()\n\t\t\treturn nil\n\t\t},\n\t})\n\n\tif err := app.Run(os.Args); err != nil {\n\t\tlog.Fatal(err)\n\t}\n}\n',l={},p="Cross-Namespace",c={id:"examples/xns",title:"Cross-Namespace",description:"A simple example showcasing usage of the generated XNS helpers for simplifying cross-namespace and even cross-cluster integrations.",source:"@site/docs/examples/xns.mdx",sourceDirName:"examples",slug:"/examples/xns",permalink:"/protoc-gen-go-temporal/docs/examples/xns",draft:!1,unlisted:!1,editUrl:"https://github.com/cludden/protoc-gen-go-temporal/tree/main/docs/docs/examples/xns.mdx",tags:[],version:"current",frontMatter:{},sidebar:"examples",previous:{title:"Codec Server",permalink:"/protoc-gen-go-temporal/docs/examples/codecserver"},next:{title:"Hello World",permalink:"/protoc-gen-go-temporal/docs/examples/mutex"}},m={},u=[{value:"Run this example",id:"run-this-example",level:2}];function g(e){const n={code:"code",h1:"h1",h2:"h2",li:"li",ol:"ol",p:"p",pre:"pre",...(0,r.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)(n.h1,{id:"cross-namespace",children:"Cross-Namespace"}),"\n",(0,o.jsx)(n.p,{children:"A simple example showcasing usage of the generated XNS helpers for simplifying cross-namespace and even cross-cluster integrations."}),"\n",(0,o.jsx)(s.A,{language:"protobuf",title:"example.proto",children:a}),"\n",(0,o.jsx)(s.A,{language:"go",title:"main.go",children:i}),"\n",(0,o.jsx)(n.h2,{id:"run-this-example",children:"Run this example"}),"\n",(0,o.jsxs)(n.ol,{children:["\n",(0,o.jsxs)(n.li,{children:["Clone the examples","\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-sh",children:"git clone https://github.com/cludden/protoc-gen-go-temporal && cd protoc-gen-go-temporal\n"})}),"\n"]}),"\n",(0,o.jsxs)(n.li,{children:["Start temporal","\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-shell",children:'temporal server start-dev \\\n    --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \\\n    --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true"\n'})}),"\n"]}),"\n",(0,o.jsxs)(n.li,{children:["In a different terminal, create ",(0,o.jsx)(n.code,{children:"example"})," namespace and run the worker","\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-shell",children:"temporal operator namespace create example\ngo run ./examples/xns/... worker\n"})}),"\n"]}),"\n",(0,o.jsxs)(n.li,{children:["In a different terminal, execute an xns workflow","\n",(0,o.jsx)(n.pre,{children:(0,o.jsx)(n.code,{className:"language-shell",children:"go run ./examples/xns/... xns provision-foo --name test\n"})}),"\n"]}),"\n"]})]})}function d(e={}){const{wrapper:n}={...(0,r.R)(),...e.components};return n?(0,o.jsx)(n,{...e,children:(0,o.jsx)(g,{...e})}):g(e)}}}]);