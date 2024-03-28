"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[8351],{2679:(e,t,o)=>{o.r(t),o.d(t,{assets:()=>u,contentTitle:()=>d,default:()=>h,frontMatter:()=>p,metadata:()=>g,toc:()=>f});var n=o(4848),r=o(8453),s=o(1432);const i='syntax="proto3";\n\npackage example.v1;\n\nimport "google/protobuf/empty.proto";\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  };\n\n  // CreateFoo creates a new foo operation\n  rpc CreateFoo(CreateFooRequest) returns (CreateFooResponse) {\n    option (temporal.v1.workflow) = {\n      execution_timeout: { seconds: 3600 }\n      id: \'create-foo/${! name.slug() }\'\n      id_reuse_policy: WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE\n      query: { ref: "GetFooProgress" }\n      signal: { ref: "SetFooProgress", start: true }\n      update: { ref: "UpdateFooProgress" }\n    };\n  }\n\n  // GetFooProgress returns the status of a CreateFoo operation\n  rpc GetFooProgress(google.protobuf.Empty) returns (GetFooProgressResponse) {\n    option (temporal.v1.query) = {};\n  }\n\n  // Notify sends a notification\n  rpc Notify(NotifyRequest) returns (google.protobuf.Empty) {\n    option (temporal.v1.activity) = {\n      start_to_close_timeout: { seconds: 30 }\n      retry_policy: {\n        max_attempts: 3\n      }\n    };\n  }\n\n  // SetFooProgress sets the current status of a CreateFoo operation\n  rpc SetFooProgress(SetFooProgressRequest) returns (google.protobuf.Empty) {\n    option (temporal.v1.signal) = {};\n  }\n\n  // UpdateFooProgress sets the current status of a CreateFoo operation\n  rpc UpdateFooProgress(SetFooProgressRequest) returns (GetFooProgressResponse) {\n    option (temporal.v1.update) = {\n      id: \'update-progress/${! progress.string() }\'\n    };\n  }\n}\n\n// CreateFooRequest describes the input to a CreateFoo workflow\nmessage CreateFooRequest {\n  // unique foo name\n  string name = 1;\n}\n\n// SampleWorkflowWithMutexResponse describes the output from a CreateFoo workflow\nmessage CreateFooResponse {\n  Foo foo = 1; \n}\n\n// Foo describes an illustrative foo resource\nmessage Foo {\n  string name = 1;\n  Status status = 2;\n\n  enum Status {\n    FOO_STATUS_UNSPECIFIED = 0;\n    FOO_STATUS_READY = 1;\n    FOO_STATUS_CREATING = 2;\n  }\n}\n\n// GetFooProgressResponse describes the output from a GetFooProgress query\nmessage GetFooProgressResponse {\n  float progress = 1;\n  Foo.Status status = 2;\n}\n\n// NotifyRequest describes the input to a Notify activity\nmessage NotifyRequest {\n  string message = 1;\n}\n\n// SetFooProgressRequest describes the input to a SetFooProgress signal\nmessage SetFooProgressRequest {\n  // value of current workflow progress\n  float progress = 1;\n}\n',a='package example\n\nimport (\n\t"context"\n\t"fmt"\n\n\texamplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"\n\t"go.temporal.io/sdk/activity"\n\t"go.temporal.io/sdk/log"\n\t"go.temporal.io/sdk/workflow"\n)\n\ntype (\n\t// Workflows manages shared state for workflow constructors and is used to\n\t// register workflows with a worker\n\tWorkflows struct{}\n\n\t// Activities manages shared state for activities and is used to register\n\t// activities with a worker\n\tActivities struct{}\n\n\t// CreateFooWorkflow manages workflow state for a CreateFoo workflow\n\tCreateFooWorkflow struct {\n\t\t// it embeds the generated workflow Input type that contains the workflow\n\t\t// input and signal helpers\n\t\t*examplev1.CreateFooWorkflowInput\n\n\t\tlog      log.Logger\n\t\tprogress float32\n\t\tstatus   examplev1.Foo_Status\n\t}\n)\n\n// CreateFoo implements a CreateFoo workflow constructor on the shared Workflows struct\n// that initializes a new CreateFooWorkflow for each execution\nfunc (w *Workflows) CreateFoo(ctx workflow.Context, input *examplev1.CreateFooWorkflowInput) (examplev1.CreateFooWorkflow, error) {\n\treturn &CreateFooWorkflow{\n\t\tCreateFooWorkflowInput: input,\n\t\tlog:                    workflow.GetLogger(ctx),\n\t\tstatus:                 examplev1.Foo_FOO_STATUS_CREATING,\n\t}, nil\n}\n\n// Execute defines the entrypoint to a CreateFooWorkflow\nfunc (wf *CreateFooWorkflow) Execute(ctx workflow.Context) (*examplev1.CreateFooResponse, error) {\n\t// listen for signals using generated signal provided by workflow input\n\tworkflow.Go(ctx, func(ctx workflow.Context) {\n\t\tfor {\n\t\t\tsignal, _ := wf.SetFooProgress.Receive(ctx)\n\t\t\twf.UpdateFooProgress(ctx, signal)\n\t\t}\n\t})\n\n\t// execute Notify activity using generated helper\n\tif err := examplev1.Notify(ctx, &examplev1.NotifyRequest{\n\t\tMessage: fmt.Sprintf("creating foo resource (%s)", wf.Req.GetName()),\n\t}); err != nil {\n\t\treturn nil, fmt.Errorf("error sending notification: %w", err)\n\t}\n\n\t// block until progress has reached 100 via signals and/or updates\n\tif err := workflow.Await(ctx, func() bool {\n\t\treturn wf.status == examplev1.Foo_FOO_STATUS_READY\n\t}); err != nil {\n\t\treturn nil, fmt.Errorf("error awaiting ready status: %w", err)\n\t}\n\n\treturn &examplev1.CreateFooResponse{\n\t\tFoo: &examplev1.Foo{\n\t\t\tName:   wf.Req.GetName(),\n\t\t\tStatus: wf.status,\n\t\t},\n\t}, nil\n}\n\n// GetFooProgress defines the handler for a GetFooProgress query\nfunc (wf *CreateFooWorkflow) GetFooProgress() (*examplev1.GetFooProgressResponse, error) {\n\treturn &examplev1.GetFooProgressResponse{Progress: wf.progress, Status: wf.status}, nil\n}\n\n// UpdateFooProgress defines the handler for an UpdateFooProgress update\nfunc (wf *CreateFooWorkflow) UpdateFooProgress(ctx workflow.Context, req *examplev1.SetFooProgressRequest) (*examplev1.GetFooProgressResponse, error) {\n\twf.progress = req.GetProgress()\n\tswitch {\n\tcase wf.progress < 0:\n\t\twf.progress, wf.status = 0, examplev1.Foo_FOO_STATUS_CREATING\n\tcase wf.progress < 100:\n\t\twf.status = examplev1.Foo_FOO_STATUS_CREATING\n\tcase wf.progress >= 100:\n\t\twf.progress, wf.status = 100, examplev1.Foo_FOO_STATUS_READY\n\t}\n\treturn &examplev1.GetFooProgressResponse{Progress: wf.progress, Status: wf.status}, nil\n}\n\n// Notify defines the implementation for a Notify activity\nfunc (a *Activities) Notify(ctx context.Context, req *examplev1.NotifyRequest) error {\n\tactivity.GetLogger(ctx).Info("notification", "message", req.GetMessage())\n\treturn nil\n}\n',l='package main\n\nimport (\n\t"log"\n\t"os"\n\n\t"github.com/cludden/protoc-gen-go-temporal/examples/example"\n\texamplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"\n\t"github.com/urfave/cli/v2"\n\t"go.temporal.io/sdk/client"\n\t"go.temporal.io/sdk/worker"\n)\n\nfunc main() {\n\t// initialize the generated cli application\n\tapp, err := examplev1.NewExampleCli(\n\t\texamplev1.NewExampleCliOptions().WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {\n\t\t\t// register activities and workflows using generated helpers\n\t\t\tw := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})\n\t\t\texamplev1.RegisterExampleActivities(w, &example.Activities{})\n\t\t\texamplev1.RegisterExampleWorkflows(w, &example.Workflows{})\n\t\t\treturn w, nil\n\t\t}),\n\t)\n\tif err != nil {\n\t\tlog.Fatalf("error initializing example cli: %v", err)\n\t}\n\n\t// run cli\n\tif err := app.Run(os.Args); err != nil {\n\t\tlog.Fatal(err)\n\t}\n}\n',c='package main\n\nimport (\n\t"context"\n\t"log"\n\n\texamplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"\n\t"go.temporal.io/sdk/client"\n)\n\nfunc main() {\n\t// initialize service client with sdk client\n\tc, _ := client.Dial(client.Options{})\n\tclient, ctx := examplev1.NewExampleClient(c), context.Background()\n\n\t// execute a workflow asynchronously\n\trun, _ := client.CreateFooAsync(ctx, &examplev1.CreateFooRequest{Name: "test"})\n\tlog.Printf("started workflow: workflow_id=%s, run_id=%s\\n", run.ID(), run.RunID())\n\n\t// send a signal to the workflow\n\tlog.Println("signalling progress")\n\t_ = run.SetFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 5.7})\n\n\t// query the workflow\n\tprogress, _ := run.GetFooProgress(ctx)\n\tlog.Printf("queried progress: %s\\n", progress.String())\n\n\t// update the workflow\n\tupdate, _ := run.UpdateFooProgress(ctx, &examplev1.SetFooProgressRequest{Progress: 100})\n\tlog.Printf("updated progress: %s\\n", update.String())\n\n\t// block on workflow completion\n\tresp, _ := run.Get(ctx)\n\tlog.Printf("workflow completed: %s\\n", resp.String())\n}\n',p={sidebar_position:1},d="About",g={id:"about",title:"About",description:"A protoc plugin for generating typed Temporal clients and workers in Go from protobuf schemas. This plugin allows:",source:"@site/docs/about.mdx",sourceDirName:".",slug:"/about",permalink:"/protoc-gen-go-temporal/docs/about",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/about.mdx",tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"docs",next:{title:"Install",permalink:"/protoc-gen-go-temporal/docs/install"}},u={},f=[{value:"How It Works",id:"how-it-works",level:2},{value:"Features",id:"features",level:2},{value:"Inspiration",id:"inspiration",level:2}];function m(e){const t={a:"a",code:"code",em:"em",h1:"h1",h2:"h2",li:"li",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,r.R)(),...e.components},{Details:o}=t;return o||function(e,t){throw new Error("Expected "+(t?"component":"object")+" `"+e+"` to be defined: you likely forgot to import, pass, or provide it.")}("Details",!0),(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(t.h1,{id:"about",children:"About"}),"\n",(0,n.jsx)(t.p,{children:"A protoc plugin for generating typed Temporal clients and workers in Go from protobuf schemas. This plugin allows:"}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsx)(t.li,{children:"workflow authors to configure sensible defaults and guardrails"}),"\n",(0,n.jsx)(t.li,{children:"simplifies the implementation and testing of Temporal workers"}),"\n",(0,n.jsx)(t.li,{children:"and streamlines integration by providing typed client SDKs and a generated CLI application"}),"\n"]}),"\n",(0,n.jsx)("iframe",{width:"560",height:"315",src:"https://www.youtube.com/embed/fqKDWZDj-c0?si=3Wgvj_nP2BnSVcum&start=912",title:"YouTube video player",frameborder:"0",allow:"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share",referrerpolicy:"strict-origin-when-cross-origin",allowFullScreen:!0}),"\n",(0,n.jsx)(t.h2,{id:"how-it-works",children:"How It Works"}),"\n",(0,n.jsxs)(o,{children:[(0,n.jsxs)("summary",{children:["1. ",(0,n.jsx)(t.strong,{children:"Annotate"})," your protobuf services and methods with Temporal options provided by this plugin"]}),(0,n.jsx)(s.A,{language:"protobuf",title:"example.proto",children:i})]}),"\n",(0,n.jsxs)(o,{children:[(0,n.jsxs)("summary",{children:["2. ",(0,n.jsx)(t.strong,{children:"Generate"})," Go code that includes types, methods, and functions for implementing Temporal clients, workers, and cli applications"]}),(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-shell",children:"buf generate\n"})})]}),"\n",(0,n.jsxs)(o,{children:[(0,n.jsxs)("summary",{children:["3. ",(0,n.jsx)(t.strong,{children:"Implement"})," the required Workflow and Activity interfaces"]}),(0,n.jsx)(s.A,{language:"go",title:"example.go",children:a})]}),"\n",(0,n.jsxs)(o,{children:[(0,n.jsxs)("summary",{children:["4. ",(0,n.jsx)(t.strong,{children:"Run"})," your Temporal worker using the generated helpers and interact with it using the generated client and/or cli"]}),(0,n.jsx)(t.p,{children:(0,n.jsx)(t.em,{children:"Sample worker entrypoint"})}),(0,n.jsx)(s.A,{language:"go",title:"main.go",children:l}),(0,n.jsx)(t.p,{children:(0,n.jsx)(t.em,{children:"Sample client usage"})}),(0,n.jsx)(s.A,{language:"go",title:"client.go",children:c}),(0,n.jsx)(t.p,{children:(0,n.jsx)(t.em,{children:"Sample CLI usage"})}),(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-bash",metastring:'title="go run main.go -h"',children:"NAME:\n   example - A new cli application\n\nUSAGE:\n   example [global options] command [command options] [arguments...]\n\nCOMMANDS:\n   worker   runs a example.v1.Example worker process\n   help, h  Shows a list of commands or help for one command\n   QUERIES:\n     get-foo-progress  GetFooProgress returns the status of a CreateFoo operation\n   SIGNALS:\n     set-foo-progress  SetFooProgress sets the current status of a CreateFoo operation\n   UPDATES:\n     update-foo-progress  UpdateFooProgress sets the current status of a CreateFoo operation\n   WORKFLOWS:\n     create-foo                        CreateFoo creates a new foo operation\n     create-foo-with-set-foo-progress  sends a example.v1.Example.SetFooProgress signal to a example.v1.Example.CreateFoo workflow, starting it if necessary\n\nGLOBAL OPTIONS:\n   --help, -h  show help\n"})}),(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-sh",metastring:'title="go run main.go create-foo -d --name test"',children:"success\nworkflow id: create-foo/test\nrun id: 44cacae1-6a13-4b4a-8db7-d29eaafd1499\n"})}),(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-sh",metastring:'title="go run main.go set-foo-progress -w create-foo/test --progress 5.7"',children:"success\n"})}),(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-json",metastring:'title="go run main.go get-foo-progress -w create-foo/test"',children:'{\n  "progress": 5.7,\n  "status": "FOO_STATUS_CREATING"\n}\n'})}),(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-json",metastring:'title="go run main.go update-foo-progress -w create-foo/test --progress 100"',children:'{\n  "progress": 100,\n  "status": "FOO_STATUS_READY"\n}\n'})}),(0,n.jsx)(t.pre,{children:(0,n.jsx)(t.code,{className:"language-json",metastring:'title="go run main.go get-foo-progress -w create-foo/test"',children:'{\n  "progress": 100,\n  "status": "FOO_STATUS_READY"\n}\n'})})]}),"\n",(0,n.jsx)(t.h2,{id:"features",children:"Features"}),"\n",(0,n.jsxs)(t.p,{children:["Generated ",(0,n.jsx)(t.strong,{children:"Client"})," with:"]}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsx)(t.li,{children:"methods for executing workflows, queries, signals, and updates"}),"\n",(0,n.jsx)(t.li,{children:"methods for cancelling or terminating workflows"}),"\n",(0,n.jsxs)(t.li,{children:["default ",(0,n.jsx)(t.code,{children:"client.StartWorkflowOptions"})," and ",(0,n.jsx)(t.code,{children:"client.UpdateWorkflowWithOptionsRequest"})]}),"\n",(0,n.jsxs)(t.li,{children:["dynamic workflow ids, update ids, and search attributes via ",(0,n.jsx)(t.a,{href:"/docs/guides/bloblang",children:"Bloblang expressions"})]}),"\n",(0,n.jsx)(t.li,{children:"default timeouts, id reuse policies, retry policies, wait policies"}),"\n"]}),"\n",(0,n.jsxs)(t.p,{children:["Generated ",(0,n.jsx)(t.strong,{children:"Worker"})," resources with:"]}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsx)(t.li,{children:"functions for calling activities and local activities from workflows"}),"\n",(0,n.jsx)(t.li,{children:"functions for executing child workflows and signalling external workflows"}),"\n",(0,n.jsxs)(t.li,{children:["default ",(0,n.jsx)(t.code,{children:"workflow.ActivityOptions"}),", ",(0,n.jsx)(t.code,{children:"workflow.ChildWorkflowOptions"})]}),"\n",(0,n.jsx)(t.li,{children:"default timeouts, parent cose policies, retry policies"}),"\n"]}),"\n",(0,n.jsxs)(t.p,{children:["Optional ",(0,n.jsx)(t.strong,{children:"CLI"})," with:"]}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsx)(t.li,{children:"commands for executing workflows, synchronously or asynchronously"}),"\n",(0,n.jsx)(t.li,{children:"commands for starting workflows with signals, synchronously or asynchronously"}),"\n",(0,n.jsx)(t.li,{children:"commands for querying existing workflows"}),"\n",(0,n.jsx)(t.li,{children:"commands for sending signals to existing workflows"}),"\n",(0,n.jsx)(t.li,{children:"typed flags for conventiently specifying workflow, query, and signal inputs"}),"\n"]}),"\n",(0,n.jsxs)(t.p,{children:["Generated ",(0,n.jsx)(t.a,{href:"/docs/guides/xns",children:"Cross-Namespace (XNS)"})," helpers: ",(0,n.jsx)(t.strong,{children:"[Experimental]"})]}),"\n",(0,n.jsxs)(t.ul,{children:["\n",(0,n.jsx)(t.li,{children:"with support for invoking a service's workflows, queries, signals, and updates from workflows in a different temporal namespace"}),"\n"]}),"\n",(0,n.jsxs)(t.p,{children:["Generated ",(0,n.jsx)(t.a,{href:"/docs/guides/codec-server",children:"Remote Codec Server"})," helpers"]}),"\n",(0,n.jsxs)(t.p,{children:["Generated ",(0,n.jsx)(t.a,{href:"/docs/guides/documentation",children:"Markdown Documentation"})]}),"\n",(0,n.jsx)(t.h2,{id:"inspiration",children:"Inspiration"}),"\n",(0,n.jsxs)(t.p,{children:["This project was inspired by ",(0,n.jsx)(t.a,{href:"https://github.com/cretz/",children:"Chad Retz's"})," awesome ",(0,n.jsx)(t.a,{href:"https://github.com/cretz/temporal-sdk-go-advanced",children:"github.com/cretz/temporal-sdk-go-advanced"})," and ",(0,n.jsx)(t.a,{href:"https://github.com/jlegrone/",children:"Jacob LeGrone's"})," excellent Replay talk on ",(0,n.jsx)(t.a,{href:"https://youtu.be/LxgkAoTSI8Q?si=ZGwwbfbMz48MkPhj&t=681",children:"Temporal @ Datadog"})]}),"\n",(0,n.jsx)("iframe",{width:"560",height:"315",src:"https://www.youtube.com/embed/LxgkAoTSI8Q?si=L3O5it48sy38dsx7&start=681",title:"YouTube video player",frameborder:"0",allow:"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share",referrerpolicy:"strict-origin-when-cross-origin",allowFullScreen:!0})]})}function h(e={}){const{wrapper:t}={...(0,r.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(m,{...e})}):m(e)}}}]);