"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[325],{1912:(e,n,o)=>{o.r(n),o.d(n,{assets:()=>c,contentTitle:()=>s,default:()=>d,frontMatter:()=>i,metadata:()=>p,toc:()=>u});var r=o(4848),t=o(8453),l=o(1470),a=o(9365);const i={},s="Workflows",p={id:"guides/workflows",title:"Workflows",description:"Implementation",source:"@site/docs/guides/workflows.mdx",sourceDirName:"guides",slug:"/guides/workflows",permalink:"/protoc-gen-go-temporal/docs/guides/workflows",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/guides/workflows.mdx",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"Update",permalink:"/protoc-gen-go-temporal/docs/configuration/update"},next:{title:"Activities",permalink:"/protoc-gen-go-temporal/docs/guides/activities"}},c={},u=[{value:"Implementation",id:"implementation",level:2},{value:"Parameters",id:"parameters",level:3},{value:"Registration",id:"registration",level:2},{value:"Composite",id:"composite",level:3},{value:"Individual",id:"individual",level:3},{value:"Initializers",id:"initializers",level:2},{value:"Invocation",id:"invocation",level:2},{value:"Client",id:"client",level:3},{value:"Command Line Interface",id:"command-line-interface",level:3},{value:"Child Workflows",id:"child-workflows",level:3},{value:"Cross-Namespace (XNS)",id:"cross-namespace-xns",level:3},{value:"Workflow Functions",id:"workflow-functions",level:3}];function m(e){const n={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",pre:"pre",ul:"ul",...(0,t.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(n.h1,{id:"workflows",children:"Workflows"}),"\n",(0,r.jsx)(n.h2,{id:"implementation",children:"Implementation"}),"\n",(0,r.jsxs)(n.p,{children:["A workflow is implemented as a Go ",(0,r.jsx)(n.code,{children:"struct"})," that:"]}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsxs)(n.li,{children:["satisfies the generated ",(0,r.jsx)(n.code,{children:"<Workflow>Workflow"})," interface type generated by the plugin"]}),"\n",(0,r.jsxs)(n.li,{children:["embeds the generated ",(0,r.jsx)(n.code,{children:"<Workflow>WorkflowInput"})," struct that contains the workflow input and any registered signals"]}),"\n"]}),"\n",(0,r.jsxs)(l.A,{children:[(0,r.jsx)(a.A,{value:"implementation-workflow",label:"Go",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="example.go"',children:'package example\n\nimport (\n    examplev1 "path/to/gen/example/v1"\n    "go.temporal.io/sdk/workflow"\n)\n\ntype HelloWorkflow struct {\n    *examplev1.HelloWorkflowInput\n}\n\nfunc (w *HelloWorkflow) Execute(ctx workflow.Context) (*examplev1.HelloOutput, error) {\n    workflow.GetLogger(ctx).Info("executing hello workflow", "input", w.Req)\n    return &examplev1.HelloOutput{}, nil\n}\n\n// type assertion for illustrative purposes\nvar _ examplev1.HelloWorkflow = (*HelloWorkflow)(nil)\n'})})}),(0,r.jsx)(a.A,{value:"implementation-schema",label:"Schema",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  // Hello returns a friendly greeting\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n'})})})]}),"\n",(0,r.jsx)(n.h3,{id:"parameters",children:"Parameters"}),"\n",(0,r.jsxs)(n.p,{children:["Every ",(0,r.jsx)(n.code,{children:"<Workflow>Workflow"})," interface includes an ",(0,r.jsx)(n.code,{children:"Execute"})," method that defines the workflow entrypoint. The signature of this method varies based on whether or not the workflow specifies a non-empty output message type."]}),"\n",(0,r.jsxs)(l.A,{children:[(0,r.jsxs)(a.A,{value:"parameters-both",label:"Input & Output Parameters",children:[(0,r.jsx)(n.admonition,{type:"tip",children:(0,r.jsx)(n.p,{children:"Most workflows should specify both an input and output message type, even if the type is empty. This to support the addition of fields to either the input or output (or both) in the future without needing to introduce a breaking change."})}),(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n'})}),(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="main.go"',children:'package main\n\nimport (\n    "fmt"\n\n    examplev1 "path/to/gen/example/v1"\n)\n\ntype HelloWorkflow struct {\n    *examplev1.HelloWorkflowInput\n}\n\nfunc (w *HelloWorkflow) Execute(ctx workflow.Context) (*examplev1.HelloOutput, error) {\n    return &examplev1.HelloOutput{\n        Result: fmt.Sprintf("Hello %s!", w.Req.GetName()),\n    }, nil\n}\n'})})]}),(0,r.jsxs)(a.A,{value:"parameters-input",label:"No Output Parameter",children:[(0,r.jsxs)(n.p,{children:["A workflow output can be omitted using the native ",(0,r.jsx)(n.a,{href:"https://protobuf.dev/reference/protobuf/google.protobuf/#empty",children:"google.protobuf.Empty"})," type. This modifies the signature of the workflow's ",(0,r.jsx)(n.code,{children:"Execute"})," method to have a single return value of type ",(0,r.jsx)(n.code,{children:"error"}),". Note that this also requires an additional ",(0,r.jsx)(n.code,{children:"google/protobuf/empty.proto"})," protobuf import statement."]}),(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nimport "google/protobuf/empty.proto";\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  rpc Hello(HelloInput) returns (google.protobuf.Empty) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n'})}),(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="main.go"',children:'package main\n\nimport (\n    examplev1 "path/to/gen/example/v1"\n    "go.temporal.io/sdk/workflow"\n)\n\ntype HelloWorkflow struct {\n    *examplev1.HelloWorkflowInput\n}\n\nfunc (w *HelloWorkflow) Execute(ctx workflow.Context) error {\n    workflow.GetLogger(ctx).Info("hello!", "name", w.Req.GetName())\n    return nil\n}\n'})})]}),(0,r.jsxs)(a.A,{value:"parameters-output",label:"No Input Parameter",children:[(0,r.jsxs)(n.p,{children:["A workflow input can be omitted using the native ",(0,r.jsx)(n.a,{href:"https://protobuf.dev/reference/protobuf/google.protobuf/#empty",children:"google.protobuf.Empty"})," type. This does not modify the signature of the workflow's ",(0,r.jsx)(n.code,{children:"Execute"})," method, but does omit the ",(0,r.jsx)(n.code,{children:"Req"})," field from the workflow input structure. Note that this also requires an additional ",(0,r.jsx)(n.code,{children:"google/protobuf/empty.proto"})," protobuf import statement."]}),(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nimport "google/protobuf/empty.proto";\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  // Hello returns a friendly greeting\n  rpc Hello(google.protobuf.Empty) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n'})}),(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="main.go"',children:'package main\n\nimport (\n    "fmt"\n\n    examplev1 "path/to/gen/example/v1"\n)\n\ntype HelloWorkflow struct {\n    *examplev1.HelloWorkflowInput\n}\n\nfunc (w *HelloWorkflow) Execute(ctx workflow.Context) (*examplev1.HelloOutput, error) {\n    return &examplev1.HelloOutput{\n        Result: fmt.Sprintf("Hello World!"),\n    }, nil\n}\n'})})]})]}),"\n",(0,r.jsx)(n.h2,{id:"registration",children:"Registration"}),"\n",(0,r.jsx)(n.p,{children:"The plugin generates helpers for registering your workflows with a Temporal worker. These helpers rely on user-defined constructor functions. There are two flavors of registration helpers, composite and individual."}),"\n",(0,r.jsx)(n.h3,{id:"composite",children:"Composite"}),"\n",(0,r.jsx)(n.admonition,{type:"tip",children:(0,r.jsx)(n.p,{children:"The composite registration helper is the recommended approach for registrating workflows."})}),"\n",(0,r.jsxs)(n.p,{children:["Each protobuf service with Temporal workflow definitions generates a ",(0,r.jsx)(n.code,{children:"Register<Service>Workflows"})," composite registration function that registers all service workflows defined on a given protobuf service. This function receives two inputs:"]}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsxs)(n.li,{children:["a ",(0,r.jsx)(n.a,{href:"https://pkg.go.dev/go.temporal.io/sdk/worker#Registry",children:"worker.Registry"})," to register the Service workflows with"]}),"\n",(0,r.jsxs)(n.li,{children:["a struct value implementing the ",(0,r.jsx)(n.code,{children:"<Service>Workflows"})," interface generated by the plugin. The interface describes a struct with methods for each workflow that initialize a new workflow value for an individual execution."]}),"\n"]}),"\n",(0,r.jsxs)(l.A,{children:[(0,r.jsx)(a.A,{value:"go-registration-composite",label:"Go",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="main.go"',children:'package main\n\nimport (\n    "log"\n\n    examplev1 "path/to/gen/example/v1"\n    "go.temporal.io/sdk/client"\n    "go.temporal.io/sdk/worker"\n)\n\n// Workflows provides constructor methods for example.v1.Example workflows\ntype Workflows struct {}\n\n// FooWorkflow implements an example.v1.Example.Foo workflow\ntype FooWorkflow struct {\n    *examplev1.FooWorkflowInput\n}\n\n// Foo initializes a new examplev1.Workflow value\nfunc (w *Workflows) Foo(ctx workflow.Context, input *examplev1.FooWorkflowInput) (examplev1.FooWorkflow, error) {\n    return &FooWorld{input}, nil\n}\n\n// Execute defines the entrypoint to  an example.v1.Example.Foo workflow\nfunc (w *FooWorkflow) Execute(ctx workflow.Context) (*examplev1.FooOutput, error) {\n    return &examplev1.FooOutput{}, nil\n}\n\n// BarWorkflow implements an example.v1.Example.Bar workflow\ntype BarWorkflow struct {\n    *examplev1.BarWorkflowInput\n}\n\n// Bar initializes a new examplev1.Workflow value\nfunc (w *Workflows) Bar(ctx workflow.Context, input *examplev1.BarWorkflowInput) (examplev1.BarWorkflow, error) {\n    return &BarWorld{input}, nil\n}\n\n// Execute defines the entrypoint to  an example.v1.Example.Bar workflow\nfunc (w *BarWorkflow) Execute(ctx workflow.Context) (*examplev1.BarOutput, error) {\n    return &examplev1.BarOutput{}, nil\n}\n\nfunc main() {\n    // initialize temporal client and worker\n    c, err := client.Dial(client.Options{})\n    if err != nil {\n        log.Fatalf("error initializing client: %v", err)\n    }\n    w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})\n\n    // Register all example.v1.Example workflows with the worker\n    examplev1.RegisterExampleWorkflows(w, &Workflows{})\n    w.Run(worker.InterruptCh())\n}\n'})})}),(0,r.jsx)(a.A,{value:"schema-registration-composite",label:"Schema",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  }\n\n  rpc Foo(FooInput) returns (FooOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n\n  rpc Bar(BarInput) returns (BarOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n'})})})]}),"\n",(0,r.jsx)(n.h3,{id:"individual",children:"Individual"}),"\n",(0,r.jsxs)(n.p,{children:["Each workflow definitions generates a ",(0,r.jsx)(n.code,{children:"Register<Workflow>Workflow"})," individual registration function. This function receives two inputs:"]}),"\n",(0,r.jsxs)(n.ul,{children:["\n",(0,r.jsxs)(n.li,{children:["a ",(0,r.jsx)(n.a,{href:"https://pkg.go.dev/go.temporal.io/sdk/worker#Worker",children:"worker.Worker"})," to register the workflow with"]}),"\n",(0,r.jsx)(n.li,{children:"a constructor function that receives as input the workflow execution context and generated workflow input and initializes a new workflow value for an individual execution"}),"\n"]}),"\n",(0,r.jsxs)(l.A,{children:[(0,r.jsx)(a.A,{value:"go-registration-individual",label:"Go",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="main.go"',children:'package main\n\nimport (\n    "log"\n\n    examplev1 "path/to/gen/example/v1"\n    "go.temporal.io/sdk/client"\n    "go.temporal.io/sdk/worker"\n    "go.temporal.io/sdk/workflow"\n)\n\n// FooWorkflow implements an example.v1.Example.Foo workflow\ntype FooWorkflow struct {\n    *examplev1.FooWorkflowInput\n}\n\n// NewFooWorkflow initializes a new examplev1.Workflow value\nfunc NewFooWorkflow(ctx workflow.Context, input *examplev1.FooWorkflowInput) (examplev1.FooWorkflow, error) {\n    return &FooWorld{input}, nil\n}\n\n// Execute defines the entrypoint to  an example.v1.Example.Foo workflow\nfunc (w *FooWorkflow) Execute(ctx workflow.Context) (*examplev1.FooOutput, error) {\n    return &examplev1.FooOutput{}, nil\n}\n\n// BarWorkflow implements an example.v1.Example.Bar workflow\ntype BarWorkflow struct {\n    *examplev1.BarWorkflowInput\n}\n\n// NewBarWorkflow initializes a new examplev1.Workflow value\nfunc NewBarWorkflow(ctx workflow.Context, input *examplev1.BarWorkflowInput) (examplev1.BarWorkflow, error) {\n    return &BarWorld{input}, nil\n}\n\n// Execute defines the entrypoint to  an example.v1.Example.Bar workflow\nfunc (w *BarWorkflow) Execute(ctx workflow.Context) (*examplev1.BarOutput, error) {\n    return &examplev1.BarOutput{}, nil\n}\n\nfunc main() {\n    // initialize temporal client and worker\n    c, err := client.Dial(client.Options{})\n    if err != nil {\n        log.Fatalf("error initializing client: %v", err)\n    }\n    w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})\n\n    // Register all example.v1.Example workflows individually\n    examplev1.RegisterFooWorkflow(w, NewFooWorkflow)\n    examplev1.RegisterBarWorkflow(w, NewBarWorkflow)\n    w.Run(worker.InterruptCh())\n}\n'})})}),(0,r.jsx)(a.A,{value:"schema-registration-individual",label:"Schema",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  }\n\n  rpc Foo(FooInput) returns (FooOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n\n  rpc Bar(BarInput) returns (BarOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n'})})})]}),"\n",(0,r.jsx)(n.h2,{id:"initializers",children:"Initializers"}),"\n",(0,r.jsxs)(n.p,{children:["Workflow structs can implement an optional ",(0,r.jsx)(n.code,{children:"Initialize"})," method which will be invoked prior to signal channel initialization and query or update handler registrations. This can be useful if a workflow requires the use of an activity to initialize local workflow state."]}),"\n",(0,r.jsxs)(l.A,{children:[(0,r.jsx)(a.A,{value:"initializer-go",label:"Go",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="example.go"',children:'package example\n\nimport (\n    "context"\n    "errors"\n    "log"\n\n    examplev1 "path/to/gen/example/v1"\n    "go.temporal.io/sdk/workflow"\n)\n\ntype (\n    Workflows struct {}\n\n    FooWorkflow struct {\n        *examplev1.FooWorkflowInput\n        data map[string]any\n    }\n)\n\nfunc (w *Workflows) Foo(ctx workflow.Context, input *examplev1.FooWorkflowInput) (examplev1.FooWorkflow, error) {\n    return &FooWorkflow{input, nil}, nil\n}\n\nfunc (w *FooWorkflow) Initialize(ctx workflow.Context) error {\n    return workflow.SideEffect(ctx, func(workflow.Context) any {\n        return map[string]any{\n            "foo": "bar",\n        }\n    }).Get(&w.data)\n}\n\nfunc (w *FooWorkflow) Execute(ctx workflow.Context) (*examplev1.FooOutput, error) {\n    if err := workflow.Await(ctx, func() bool {\n        foo, ok := w.data["foo"].(string)\n        return ok && foo != "bar"\n    }); err != nil {\n        return nil, err\n    }\n    return &examplev1.FooOutput{}, nil\n}\n\nfunc (w *FooWorkflow) Bar(ctx workflow.Context, input *examplev1.BarInput) (*examplev1.BarOutput, error) {\n    if foo, _ := w.data["foo"]; foo != "bar" {\n        return nil, errors.New("unable to update foo")\n    }\n    w.data["foo"] = input.GetBar()\n    return &examplev1.BarOutput{}, nil\n}\n'})})}),(0,r.jsx)(a.A,{value:"initializer-schema",label:"Schema",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="example.go"',children:'syntax="proto3";\n\npackage example.v1;\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  }\n\n  rpc Foo(FooInput) returns (FooOutput) {\n    option (temporal.v1.workflow) = {\n      update: { ref: "Bar" }\n    };\n  }\n\n  rpc Bar(BarInput) returns (BarOutput) {\n    option (temporal.v1.update) = {};\n  }\n}\n'})})})]}),"\n",(0,r.jsx)(n.h2,{id:"invocation",children:"Invocation"}),"\n",(0,r.jsx)(n.p,{children:"The plugin supports several methods for executing protobuf workflows, each of which is outlined in more detail below."}),"\n",(0,r.jsx)(n.h3,{id:"client",children:"Client"}),"\n",(0,r.jsxs)(n.p,{children:["Consumers can utilize the generated Client to execute workflows from any Go application. See the ",(0,r.jsx)(n.a,{href:"/docs/guides/clients",children:"Clients guide"})," for more usage details."]}),"\n",(0,r.jsxs)(l.A,{children:[(0,r.jsx)(a.A,{value:"client-go",label:"Go",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="main.go"',children:'package main\n\nimport (\n    "context"\n    "log"\n\n    examplev1 "path/to/gen/example/v1"\n    "go.temporal.io/sdk/client"\n)\n\nfunc main() {\n    // initialize temporal client\n    c, err := client.Dial(client.Options{})\n    if err != nil {\n        log.Fatalf("error initializing client: %v", err)\n    }\n\n    // initialize temporal protobuf client\n    client := examplev1.NewExampleClient(c)\n\n    // execute an example.v1.Example.Hello workflow and block until completion or non-retryable error\n    out, err := client.Hello(context.Background(), &examplev1.HelloInput{})\n    if err != nil {\n        log.Fatalf("error executing example.v1.Example.Hello workflow: %v", err)\n    }\n}\n'})})}),(0,r.jsx)(a.A,{value:"client-schema",label:"Schema",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  }\n\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n'})})})]}),"\n",(0,r.jsx)(n.h3,{id:"command-line-interface",children:"Command Line Interface"}),"\n",(0,r.jsxs)(n.p,{children:["Consumers can utilize the generated Command Line Interface as a standalone application for executing workflows. See the ",(0,r.jsx)(n.a,{href:"/docs/guides/cli",children:"CLI guide"})," for more usage details."]}),"\n",(0,r.jsxs)(l.A,{children:[(0,r.jsxs)(a.A,{value:"cli-shell",label:"Shell",children:[(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-sh",metastring:'title="example -h"',children:"NAME:\n  example - an example temporal cli\n\nUSAGE:\n  example [global options] command [command options] [arguments...]\n\nCOMMANDS:\n  help, h  Shows a list of commands or help for one command\n    WORKFLOWS:\n      hello   Hello returns a friendly greeting\n"})}),(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-sh",metastring:'title="example hello -h"',children:'NAME:\n  example hello - Hello returns a friendly greeting\n\nUSAGE:\n  example hello [command options] [arguments...]\n\nCATEGORY:\n   WORKFLOWS\n\nOPTIONS:\n   --detach, -d                  run workflow in the background and print workflow and execution id (default: false)\n   --help, -h                    show help\n   --input-file value, -f value  path to json-formatted input file\n   --task-queue value, -t value  task queue name (default: "example-v1") [$TEMPORAL_TASK_QUEUE_NAME, $TEMPORAL_TASK_QUEUE, $TASK_QUEUE_NAME, $TASK_QUEUE]\n\n   INPUT\n\n   --name value    Name specifies the subject to greet\n'})}),(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-json",metastring:'title="example hello --name Temporal"',children:'{\n  "result": "Hello Temporal!"\n}\n'})})]}),(0,r.jsx)(a.A,{value:"cli-go",label:"Go",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="main.go"',children:'package main\n\nimport (\n    "log"\n    "os"\n\n    examplev1 "path/to/gen/example/v1"\n)\n\nfunc main() {\n    app, err := examplev1.NewExampleCLI()\n    if err != nil {\n        log.Fatalf("error initializing cli: %v", err)\n    }\n\n    if err := app.Run(os.Args); err != nil {\n        log.Fatal(err)\n    }\n}\n'})})}),(0,r.jsx)(a.A,{value:"cli-schema",label:"Schema",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  }\n\n  // Hello returns a friendly greeting\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n\nmessage HelloInput {\n  // Name specifies the subject to greet\n  string name = 1;\n}\n\nmessage HelloOutput {\n  string result = 1;\n}\n'})})})]}),"\n",(0,r.jsx)(n.h3,{id:"child-workflows",children:"Child Workflows"}),"\n",(0,r.jsxs)(n.p,{children:["Workflows can be executed as child workflows from other workflows in the same Temporal namespace. See the ",(0,r.jsx)(n.a,{href:"/docs/guides/child-workflows",children:"Child Workflows guide"})," for more usage details."]}),"\n",(0,r.jsxs)(l.A,{children:[(0,r.jsx)(a.A,{value:"child-go",label:"Go",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="example.go"',children:'package main\n\nimport (\n    "fmt"\n\n    examplev1 "path/to/gen/example/v1"\n    "go.temporal.io/sdk/workflow"\n)\n\nfunc MyWorkflow(ctx workflow.Context) error {\n    out, err := examplev1.HelloChild(ctx, &examplev1.HelloInput{})\n    if err != nil {\n        return fmt.Errorf("error executing example.v1.Example.Hello child workflow: %w", err)\n    }\n    return nil\n}\n'})})}),(0,r.jsx)(a.A,{value:"child-schema",label:"Schema",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  }\n\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n'})})})]}),"\n",(0,r.jsx)(n.h3,{id:"cross-namespace-xns",children:"Cross-Namespace (XNS)"}),"\n",(0,r.jsxs)(n.p,{children:["Workflows can be executed from other workflows in a different Temporal namespace or even an entirely separate Temporal cluster (e.g. on-prem to cloud). See the ",(0,r.jsx)(n.a,{href:"/docs/guides/xns",children:"Cross-Namespace guide"})," for more usage details."]}),"\n",(0,r.jsxs)(l.A,{children:[(0,r.jsx)(a.A,{value:"xns-go",label:"Go",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="example.go"',children:'package example\n\nimport (\n    "fmt"\n\n    examplev1 "path/to/gen/example/v1"\n    "path/to/gen/example/v1/examplev1xns"\n    "go.temporal.io/sdk/workflow"\n)\n\nfunc MyWorkflow(ctx workflow.Context) error {\n    out, err := examplev1xns.Hello(ctx, &examplev1.HelloInput{})\n    if err != nil {\n        return fmt.Errorf("error executing example.v1.Example.Hello xns workflow: %w", err)\n    }\n    return nil\n}\n'})})}),(0,r.jsx)(a.A,{value:"xns-schema",label:"Schema",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  }\n\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n'})})})]}),"\n",(0,r.jsx)(n.h3,{id:"workflow-functions",children:"Workflow Functions"}),"\n",(0,r.jsx)(n.p,{children:"Workflow definitions can be executed inline with another workflow definition."}),"\n",(0,r.jsxs)(l.A,{children:[(0,r.jsx)(a.A,{value:"wffn-go",label:"Go",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-go",metastring:'title="main.go"',children:'package main\n\nimport (\n    examplev1 "path/to/gen/example/v1"\n    "go.temporal.io/sdk/client"\n    "go.temporal.io/sdk/worker"\n    "go.temporal.io/sdk/workflow"\n)\n\ntype ExampleWorkflows struct {}\n\ntype HelloWorkflow struct {\n    *examplev1.HelloWorkflowInput\n}\n\nfunc (w *ExampleWorkflows) Hello(ctx workflow.Context, input *examplev1.HelloInput) (examplev1.HelloWorkflow, error) {\n    return &HelloWorkflow{input}, nil\n}\n\nfunc (w *HelloWorkflow) Execute(ctx workflow.Context) error {\n    workflow.GetLogger(ctx).Info("hello!", "name", w.Req.GetName())\n    return nil\n}\n\nfunc MyWorkflow(ctx workflow.Context) error {\n    // this is equivalent to calling Execute inline\n    out, err := examplev1.HelloFunction(ctx, &examplev1.HelloInput{})\n    if err != nil {\n        return fmt.Errorf("error executing example.v1.Example.Hello inline: %w", err)\n    }\n    return nil\n}\n\nfunc main() {\n    // initialize temporal client and worker\n    c, err := client.Dial(client.Options{})\n    if err != nil {\n        log.Fatalf("error initializing client: %v", err)\n    }\n    w := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})\n\n    // Register all example.v1.Example workflows with the worker\n    examplev1.RegisterExampleWorkflows(w, &Workflows{})\n    w.RegisterWorkflow(MyWorkflow)\n    w.Run(worker.InterruptCh())\n}\n'})})}),(0,r.jsx)(a.A,{value:"wffn-schema",label:"Schema",children:(0,r.jsx)(n.pre,{children:(0,r.jsx)(n.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nservice Example {\n  option (temporal.v1.service) = {\n    task_queue: "example-v1"\n  }\n\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n'})})})]})]})}function d(e={}){const{wrapper:n}={...(0,t.R)(),...e.components};return n?(0,r.jsx)(n,{...e,children:(0,r.jsx)(m,{...e})}):m(e)}},9365:(e,n,o)=>{o.d(n,{A:()=>a});o(6540);var r=o(4164);const t={tabItem:"tabItem_Ymn6"};var l=o(4848);function a(e){let{children:n,hidden:o,className:a}=e;return(0,l.jsx)("div",{role:"tabpanel",className:(0,r.A)(t.tabItem,a),hidden:o,children:n})}},1470:(e,n,o)=>{o.d(n,{A:()=>b});var r=o(6540),t=o(4164),l=o(3104),a=o(6347),i=o(205),s=o(7485),p=o(1682),c=o(9466);function u(e){return r.Children.toArray(e).filter((e=>"\n"!==e)).map((e=>{if(!e||(0,r.isValidElement)(e)&&function(e){const{props:n}=e;return!!n&&"object"==typeof n&&"value"in n}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))?.filter(Boolean)??[]}function m(e){const{values:n,children:o}=e;return(0,r.useMemo)((()=>{const e=n??function(e){return u(e).map((e=>{let{props:{value:n,label:o,attributes:r,default:t}}=e;return{value:n,label:o,attributes:r,default:t}}))}(o);return function(e){const n=(0,p.X)(e,((e,n)=>e.value===n.value));if(n.length>0)throw new Error(`Docusaurus error: Duplicate values "${n.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[n,o])}function d(e){let{value:n,tabValues:o}=e;return o.some((e=>e.value===n))}function f(e){let{queryString:n=!1,groupId:o}=e;const t=(0,a.W6)(),l=function(e){let{queryString:n=!1,groupId:o}=e;if("string"==typeof n)return n;if(!1===n)return null;if(!0===n&&!o)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return o??null}({queryString:n,groupId:o});return[(0,s.aZ)(l),(0,r.useCallback)((e=>{if(!l)return;const n=new URLSearchParams(t.location.search);n.set(l,e),t.replace({...t.location,search:n.toString()})}),[l,t])]}function x(e){const{defaultValue:n,queryString:o=!1,groupId:t}=e,l=m(e),[a,s]=(0,r.useState)((()=>function(e){let{defaultValue:n,tabValues:o}=e;if(0===o.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(n){if(!d({value:n,tabValues:o}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${n}" but none of its children has the corresponding value. Available values are: ${o.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return n}const r=o.find((e=>e.default))??o[0];if(!r)throw new Error("Unexpected error: 0 tabValues");return r.value}({defaultValue:n,tabValues:l}))),[p,u]=f({queryString:o,groupId:t}),[x,w]=function(e){let{groupId:n}=e;const o=function(e){return e?`docusaurus.tab.${e}`:null}(n),[t,l]=(0,c.Dv)(o);return[t,(0,r.useCallback)((e=>{o&&l.set(e)}),[o,l])]}({groupId:t}),h=(()=>{const e=p??x;return d({value:e,tabValues:l})?e:null})();(0,i.A)((()=>{h&&s(h)}),[h]);return{selectedValue:a,selectValue:(0,r.useCallback)((e=>{if(!d({value:e,tabValues:l}))throw new Error(`Can't select invalid tab value=${e}`);s(e),u(e),w(e)}),[u,w,l]),tabValues:l}}var w=o(2303);const h={tabList:"tabList__CuJ",tabItem:"tabItem_LNqP"};var g=o(4848);function v(e){let{className:n,block:o,selectedValue:r,selectValue:a,tabValues:i}=e;const s=[],{blockElementScrollPositionUntilNextRender:p}=(0,l.a_)(),c=e=>{const n=e.currentTarget,o=s.indexOf(n),t=i[o].value;t!==r&&(p(n),a(t))},u=e=>{let n=null;switch(e.key){case"Enter":c(e);break;case"ArrowRight":{const o=s.indexOf(e.currentTarget)+1;n=s[o]??s[0];break}case"ArrowLeft":{const o=s.indexOf(e.currentTarget)-1;n=s[o]??s[s.length-1];break}}n?.focus()};return(0,g.jsx)("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,t.A)("tabs",{"tabs--block":o},n),children:i.map((e=>{let{value:n,label:o,attributes:l}=e;return(0,g.jsx)("li",{role:"tab",tabIndex:r===n?0:-1,"aria-selected":r===n,ref:e=>s.push(e),onKeyDown:u,onClick:c,...l,className:(0,t.A)("tabs__item",h.tabItem,l?.className,{"tabs__item--active":r===n}),children:o??n},n)}))})}function k(e){let{lazy:n,children:o,selectedValue:t}=e;const l=(Array.isArray(o)?o:[o]).filter(Boolean);if(n){const e=l.find((e=>e.props.value===t));return e?(0,r.cloneElement)(e,{className:"margin-top--md"}):null}return(0,g.jsx)("div",{className:"margin-top--md",children:l.map(((e,n)=>(0,r.cloneElement)(e,{key:n,hidden:e.props.value!==t})))})}function j(e){const n=x(e);return(0,g.jsxs)("div",{className:(0,t.A)("tabs-container",h.tabList),children:[(0,g.jsx)(v,{...e,...n}),(0,g.jsx)(k,{...e,...n})]})}function b(e){const n=(0,w.A)();return(0,g.jsx)(j,{...e,children:u(e.children)},String(n))}},8453:(e,n,o)=>{o.d(n,{R:()=>a,x:()=>i});var r=o(6540);const t={},l=r.createContext(t);function a(e){const n=r.useContext(l);return r.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function i(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:a(e.components),r.createElement(l.Provider,{value:n},e.children)}}}]);