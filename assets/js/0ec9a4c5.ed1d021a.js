"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[249],{6140:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>d,contentTitle:()=>c,default:()=>g,frontMatter:()=>s,metadata:()=>i,toc:()=>p});var r=n(4848),o=n(8453),l=n(1432);const a='package main\n\nimport (\n\t"context"\n\t"errors"\n\t"log"\n\t"log/slog"\n\t"net/http"\n\t"os"\n\t"os/signal"\n\t"syscall"\n\n\t"github.com/cludden/protoc-gen-go-temporal/examples/example"\n\texamplev1 "github.com/cludden/protoc-gen-go-temporal/gen/example/v1"\n\t"github.com/cludden/protoc-gen-go-temporal/pkg/codec"\n\t"github.com/cludden/protoc-gen-go-temporal/pkg/scheme"\n\t"github.com/urfave/cli/v2"\n\t"go.temporal.io/sdk/client"\n\t"go.temporal.io/sdk/converter"\n\ttlog "go.temporal.io/sdk/log"\n\t"go.temporal.io/sdk/worker"\n)\n\nfunc main() {\n\tapp, err := examplev1.NewExampleCli(\n\t\texamplev1.NewExampleCliOptions().\n\t\t\tWithClient(func(cmd *cli.Context) (client.Client, error) {\n\t\t\t\treturn client.Dial(client.Options{\n\t\t\t\t\tDataConverter: converter.NewCompositeDataConverter(\n\t\t\t\t\t\tconverter.NewNilPayloadConverter(),\n\t\t\t\t\t\tconverter.NewByteSlicePayloadConverter(),\n\t\t\t\t\t\tconverter.NewProtoPayloadConverter(),\n\t\t\t\t\t),\n\t\t\t\t\tLogger: tlog.NewStructuredLogger(slog.Default()),\n\t\t\t\t})\n\t\t\t}).\n\t\t\tWithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {\n\t\t\t\tw := worker.New(c, examplev1.ExampleTaskQueue, worker.Options{})\n\t\t\t\texamplev1.RegisterExampleActivities(w, &example.Activities{})\n\t\t\t\texamplev1.RegisterExampleWorkflows(w, &example.Workflows{})\n\t\t\t\treturn w, nil\n\t\t\t}),\n\t)\n\tif err != nil {\n\t\tlog.Fatalf("error initializing example cli: %v", err)\n\t}\n\n\tapp.Commands = append(app.Commands, &cli.Command{\n\t\tName:  "codec",\n\t\tUsage: "run remote codec server",\n\t\tAction: func(cmd *cli.Context) error {\n\t\t\thandler := converter.NewPayloadCodecHTTPHandler(\n\t\t\t\tcodec.NewProtoJSONCodec(\n\t\t\t\t\tscheme.New(\n\t\t\t\t\t\texamplev1.WithExampleSchemeTypes(),\n\t\t\t\t\t),\n\t\t\t\t),\n\t\t\t)\n\n\t\t\tsrv := &http.Server{\n\t\t\t\tAddr:    "0.0.0.0:8080",\n\t\t\t\tHandler: handler,\n\t\t\t}\n\n\t\t\tgo func() {\n\t\t\t\tsigChan := make(chan os.Signal, 1)\n\t\t\t\tsignal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)\n\t\t\t\t<-sigChan\n\n\t\t\t\tif err := srv.Shutdown(context.Background()); err != nil {\n\t\t\t\t\tlog.Fatalf("error shutting down server: %v", err)\n\t\t\t\t}\n\t\t\t}()\n\n\t\t\tif err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {\n\t\t\t\tlog.Fatalf("server error: %v", err)\n\t\t\t}\n\t\t\treturn nil\n\t\t},\n\t})\n\n\t// run cli\n\tif err := app.Run(os.Args); err != nil {\n\t\tlog.Fatal(err)\n\t}\n}\n',s={},c="Codec Server",i={id:"examples/codecserver",title:"Codec Server",description:"A simple example inspired by temporalio/samples-go/codecserver",source:"@site/docs/examples/codecserver.mdx",sourceDirName:"examples",slug:"/examples/codecserver",permalink:"/protoc-gen-go-temporal/docs/examples/codecserver",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/examples/codecserver.mdx",tags:[],version:"current",frontMatter:{},sidebar:"examples",previous:{title:"Hello World",permalink:"/protoc-gen-go-temporal/docs/examples/helloworld"},next:{title:"Search Attributes",permalink:"/protoc-gen-go-temporal/docs/examples/searchattributes"}},d={},p=[{value:"Run this example",id:"run-this-example",level:2}];function m(e){const t={a:"a",code:"code",h1:"h1",h2:"h2",li:"li",ol:"ol",p:"p",pre:"pre",...(0,o.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(t.h1,{id:"codec-server",children:"Codec Server"}),"\n",(0,r.jsxs)(t.p,{children:["A simple example inspired by ",(0,r.jsx)(t.a,{href:"https://github.com/temporalio/samples-go/tree/main/codecserver",children:"temporalio/samples-go/codecserver"})]}),"\n",(0,r.jsx)(l.A,{language:"go",title:"main.go",children:a}),"\n",(0,r.jsx)(t.h2,{id:"run-this-example",children:"Run this example"}),"\n",(0,r.jsxs)(t.ol,{children:["\n",(0,r.jsxs)(t.li,{children:["Clone the examples","\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-sh",children:"git clone https://github.com/cludden/protoc-gen-go-temporal && cd protoc-gen-go-temporal\n"})}),"\n"]}),"\n",(0,r.jsxs)(t.li,{children:["Start the codec server","\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-shell",children:"go run examples/codecserver/main.go codec\n"})}),"\n"]}),"\n",(0,r.jsxs)(t.li,{children:["In a different terminal, start temporal using the codec server","\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-shell",children:'temporal server start-dev \\\n    --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \\\n    --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true" \\\n    --ui-codec-endpoint http://localhost:8080\n'})}),"\n"]}),"\n",(0,r.jsxs)(t.li,{children:["In a different terminal, run the worker","\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-shell",children:"go run examples/codecserver/main.go worker\n"})}),"\n"]}),"\n",(0,r.jsxs)(t.li,{children:["In a different terminal, execute a workflow, signal, query, and update","\n",(0,r.jsx)(t.pre,{children:(0,r.jsx)(t.code,{className:"language-shell",children:"# execute a workflow in the background\ngo run examples/codecserver/main.go create-foo --name test -d\n\n# signal the workflow\ngo run examples/codecserver/main.go set-foo-progress -w create-foo/test --progress 5.7\n\n# query the workflow\ngo run examples/codecserver/main.go get-foo-progress -w create-foo/test\n\n# update the workflow\ngo run examples/codecserver/main.go update-foo-progress -w create-foo/test --progress 100\n"})}),"\n"]}),"\n",(0,r.jsxs)(t.li,{children:["In the UI, switch to the JSON tab and disable the ",(0,r.jsx)(t.code,{children:"Decode Event History"})," toggle and verify that all payloads have metadata with ",(0,r.jsx)(t.code,{children:'"encoding": "YmluYXJ5L3Byb3RvYnVm"'}),", which is ",(0,r.jsx)(t.code,{children:"binary/protobuf"})," base64-encoded"]}),"\n"]})]})}function g(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,r.jsx)(t,{...e,children:(0,r.jsx)(m,{...e})}):m(e)}}}]);