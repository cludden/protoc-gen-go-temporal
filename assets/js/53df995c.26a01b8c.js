"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[8507],{3818:(e,o,n)=>{n.r(o),n.d(o,{assets:()=>s,contentTitle:()=>a,default:()=>d,frontMatter:()=>l,metadata:()=>i,toc:()=>c});var r=n(4848),t=n(8453);n(1470),n(9365);const l={},a="Workflow",i={id:"configuration/workflow",title:"Workflow",description:"Workflows are defined as Protobuf RPCs annotated with the temporal.v1.workflow method option. See the Workflows guide for more usage details.",source:"@site/docs/configuration/workflow.mdx",sourceDirName:"configuration",slug:"/configuration/workflow",permalink:"/docs/configuration/workflow",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/configuration/workflow.mdx",tags:[],version:"current",frontMatter:{},sidebar:"docs",previous:{title:"Service",permalink:"/docs/configuration/service"},next:{title:"Activity",permalink:"/docs/configuration/activity"}},s={},c=[{value:"Options",id:"options",level:2},{value:"execution_timeout",id:"execution_timeout",level:3},{value:"id",id:"id",level:3},{value:"id_reuse_policy",id:"id_reuse_policy",level:3},{value:"name",id:"name",level:3},{value:"parent_close_policy",id:"parent_close_policy",level:3},{value:"query",id:"query",level:3},{value:"retry_policy",id:"retry_policy",level:3},{value:"run_timeout",id:"run_timeout",level:3},{value:"search_attributes",id:"search_attributes",level:3},{value:"signal",id:"signal",level:3},{value:"task_queue",id:"task_queue",level:3},{value:"task_timeout",id:"task_timeout",level:3},{value:"update",id:"update",level:3},{value:"xns",id:"xns",level:3},{value:"wait_for_cancellation",id:"wait_for_cancellation",level:3}];function u(e){const o={a:"a",admonition:"admonition",code:"code",h1:"h1",h2:"h2",h3:"h3",p:"p",pre:"pre",...(0,t.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(o.h1,{id:"workflow",children:"Workflow"}),"\n",(0,r.jsxs)(o.p,{children:[(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows",children:"Workflows"})," are defined as Protobuf RPCs annotated with the ",(0,r.jsx)(o.code,{children:"temporal.v1.workflow"})," method option. See the ",(0,r.jsx)(o.a,{href:"/docs/guides/workflows",children:"Workflows guide"})," for more usage details."]}),"\n",(0,r.jsx)(o.admonition,{type:"info",children:(0,r.jsxs)(o.p,{children:["Workflow definitions can omit an input and/or out parameter by specifying the native ",(0,r.jsx)(o.code,{children:"google.protobuf.Empty"})," message type in its place. This requires an additional ",(0,r.jsx)(o.code,{children:"google/protobuf/empty.proto"})," protobuf import."]})}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",metastring:'title="example.proto"',children:'syntax="proto3";\n\npackage example.v1;\n\nimport "temporal/v1/temporal.proto";\n\nservice Example {\n  // Hello returns a friendly greeting\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {};\n  }\n}\n'})}),"\n",(0,r.jsx)(o.h2,{id:"options",children:"Options"}),"\n",(0,r.jsx)(o.h3,{id:"execution_timeout",children:"execution_timeout"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.a,{href:"https://protobuf.dev/reference/protobuf/google.protobuf/#duration",children:"google.protobuf.Duration"})}),"\n",(0,r.jsxs)(o.p,{children:["The timeout for duration of ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows#time-constraints",children:"Workflow execution"}),". It includes retries and continue as new. Use ",(0,r.jsx)(o.a,{href:"#run_timeout",children:"run_timeout"})," to limit execution time of a single workflow run. Defaults to unlimited."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      execution_timeout: { seconds: 3600 } // 1h\n    };\n  }\n}\n"})}),"\n",(0,r.jsx)(o.h3,{id:"id",children:"id"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.code,{children:"string"})}),"\n",(0,r.jsxs)(o.p,{children:["Specifies the default ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows#workflow-id",children:"Workflow ID"})," as a ",(0,r.jsx)(o.a,{href:"/docs/guides/bloblang",children:"Bloblang expression"}),"."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      id: 'hello/${! name }/${! uuid_v4() }'\n    };\n  }\n}\n\nmessage HelloInput {\n  string name = 1;\n}\n"})}),"\n",(0,r.jsx)(o.h3,{id:"id_reuse_policy",children:"id_reuse_policy"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.a,{href:"https://buf.build/cludden/protoc-gen-go-temporal/docs/main:temporal.v1#temporal.v1.IDReusePolicy",children:"temporal.v1.IDReusePolicy"})}),"\n",(0,r.jsxs)(o.p,{children:["Whether server allow reuse of workflow ID, can be useful for dedupe logic if set to ",(0,r.jsx)(o.code,{children:"WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE"}),". Defaults to ",(0,r.jsx)(o.code,{children:"WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE"}),". See ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows#workflow-id-reuse-policy",children:"docs"})," for more details."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      id_reuse_policy: WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE\n    };\n  }\n}\n"})}),"\n",(0,r.jsx)(o.h3,{id:"name",children:"name"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.code,{children:"string"})}),"\n",(0,r.jsxs)(o.p,{children:["Fully qualified ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows#workflow-type",children:"Workflow type name"}),". Defaults to protobuf method full name (e.g. ",(0,r.jsx)(o.code,{children:"example.v1.Example.Hello"}),")"]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:'service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      name: "Hello"\n    };\n  }\n}\n'})}),"\n",(0,r.jsx)(o.h3,{id:"parent_close_policy",children:"parent_close_policy"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.a,{href:"https://buf.build/cludden/protoc-gen-go-temporal/docs/main:temporal.v1#temporal.v1.ParentClosePolicy",children:"temporal.v1.ParentClosePolicy"})}),"\n",(0,r.jsxs)(o.p,{children:["When execution as a child workflow, this ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows#parent-close-policy",children:"optional policy"})," determines what to do when the parent workflow is closed. Default to ",(0,r.jsx)(o.code,{children:"PARENT_CLOSE_POLICY_TERMINATE"}),"."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      parent_close_policy: PARENT_CLOSE_POLICY_REQUEST_CANCEL\n    };\n  }\n}\n"})}),"\n",(0,r.jsx)(o.h3,{id:"query",children:"query"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.a,{href:"https://buf.build/cludden/protoc-gen-go-temporal/docs/main:temporal.v1#temporal.v1.WorkflowOptions.Query",children:"temporal.v1.WorkflowOptions.Query"})}),"\n",(0,r.jsxs)(o.p,{children:["Identifies a ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows#query",children:"query"})," supported by the workflow. Can be specified 0, 1, or more times. See the ",(0,r.jsx)(o.a,{href:"/docs/guides/queries",children:"query docs"})," for more details."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      query: { ref: 'SomeQuery1' }\n      query: { ref: 'SomeQuery2' }\n      query: {\n        // can reference a query definition from another service\n        ref: 'example.v2.Example.SomeQuery3'\n      }\n    };\n  }\n\n  rpc SomeQuery1(SomeQuery1Input) returns (SomeQuery1Output) {\n    option (temporal.v1.query) = {};\n  }\n\n  rpc SomeQuery2(SomeQuery2Input) returns (SomeQuery2Output) {\n    option (temporal.v1.query) = {};\n  }\n}\n"})}),"\n",(0,r.jsx)(o.h3,{id:"retry_policy",children:"retry_policy"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.a,{href:"https://buf.build/cludden/protoc-gen-go-temporal/docs/main:temporal.v1#temporal.v1.RetryPolicy",children:"temporal.v1.RetryPolicy"})}),"\n",(0,r.jsxs)(o.p,{children:["Optional ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/retry-policies",children:"retry policy"})," for workflow. If a retry policy is specified, in case of workflow failure server will start new workflow execution if needed based on the retry policy."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:'service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      retry_policy: {\n        max_attempts: 10\n        initial_interval: { seconds: 5 }\n        max_interval: { seconds: 60 }\n        backoff_coefficient: 2.0\n        non_retryable_error_types: ["SomeError", "SomeOtherError"]\n      }\n    };\n  }\n}\n'})}),"\n",(0,r.jsx)(o.h3,{id:"run_timeout",children:"run_timeout"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.a,{href:"https://protobuf.dev/reference/protobuf/google.protobuf/#duration",children:"google.protobuf.Duration"})}),"\n",(0,r.jsxs)(o.p,{children:["The timeout for duration of ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows#time-constraints",children:"a single workflow run"}),". Defaults to ",(0,r.jsx)(o.a,{href:"#execution_timeout",children:"execution_timeout"}),"."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      execution_timeout: { seconds: 300 } // 5m\n    };\n  }\n}\n"})}),"\n",(0,r.jsx)(o.h3,{id:"search_attributes",children:"search_attributes"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.code,{children:"string"})}),"\n",(0,r.jsxs)(o.p,{children:["Specifies the default ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/visibility#search-attribute",children:"Workflow search attributes"})," as a ",(0,r.jsx)(o.a,{href:"https://www.benthos.dev/docs/guides/bloblang/about/",children:"Bloblang mapping"}),". See the ",(0,r.jsx)(o.a,{href:"/docs/examples/searchattributes/",children:"Search Attributes example"})," for more details."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      search_attributes:\n        'Name = name.or(\"World\") \\n'\n        'CreatedAt = now().ts_parse(\"2006-01-02T15:04:05Z\") \\n'\n    };\n  }\n}\n"})}),"\n",(0,r.jsx)(o.h3,{id:"signal",children:"signal"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.a,{href:"https://buf.build/cludden/protoc-gen-go-temporal/docs/main:temporal.v1#temporal.v1.WorkflowOptions.Signal",children:"temporal.v1.WorkflowOptions.Signal"})}),"\n",(0,r.jsxs)(o.p,{children:["Identifies a ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows#signal",children:"signal"})," supported by the workflow. Can be specified 0, 1, or more times. See the ",(0,r.jsx)(o.a,{href:"/docs/guides/signals",children:"signal docs"})," for more details."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      signal: { ref: 'SomeSignal1' }\n      signal: {\n        ref: 'SomeSignal2'\n        start: true // generates signal with start helpers\n      }\n      signal: {\n        // can reference a signal definition from another service\n        ref: 'example.v2.Example.SomeSignal3'\n      }\n    };\n  }\n\n  rpc SomeSignal1(SomeSignal1Input) returns (google.protobuf.Empty) {\n    option (temporal.v1.signal) = {};\n  }\n\n  rpc SomeSignal2(SomeSignal2Input) returns (google.protobuf.Empty) {\n    option (temporal.v1.signal) = {};\n  }\n}\n"})}),"\n",(0,r.jsx)(o.h3,{id:"task_queue",children:"task_queue"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.code,{children:"string"})}),"\n",(0,r.jsxs)(o.p,{children:["Overrides the default task queue for a particular workflow type. Defaults to Service's ",(0,r.jsx)(o.code,{children:"task_queue"})," if specified."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:'service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      task_queue: "example-v2"\n    };\n  }\n}\n'})}),"\n",(0,r.jsx)(o.h3,{id:"task_timeout",children:"task_timeout"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.a,{href:"https://protobuf.dev/reference/protobuf/google.protobuf/#duration",children:"google.protobuf.Duration"})}),"\n",(0,r.jsxs)(o.p,{children:["The timeout for processing ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows#workflow-task-timeout",children:"workflow task"})," from the time the worker pulled this task. If a workflow task is lost, it is retried after this timeout. Defaults to ",(0,r.jsx)(o.code,{children:"10s"}),"."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      task_timeout: { seconds: 10 }\n    };\n  }\n}\n"})}),"\n",(0,r.jsx)(o.h3,{id:"update",children:"update"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.a,{href:"https://buf.build/cludden/protoc-gen-go-temporal/docs/main:temporal.v1#temporal.v1.WorkflowOptions.Update",children:"temporal.v1.WorkflowOptions.Update"})}),"\n",(0,r.jsxs)(o.p,{children:["Identifies an ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows#update",children:"update"})," supported by the workflow. Can be specified 0, 1, or more times. See the ",(0,r.jsx)(o.a,{href:"/docs/guides/updates",children:"update docs"})," for more details."]}),"\n",(0,r.jsx)(o.admonition,{type:"note",children:(0,r.jsxs)(o.p,{children:["This requires the ",(0,r.jsx)(o.a,{href:"/docs/configuration/plugin#workflow-update-enabled",children:"workflow-update-enabled"})," plugin option to be enabled."]})}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      update: { ref: 'SomeUpdate1' }\n      update: { ref: 'SomeUpdate2' }\n      update: {\n        // can reference an update definition from another service\n        ref: 'example.v2.Example.SomeUpdate3'\n      }\n    };\n  }\n\n  rpc SomeUpdate1(SomeUpdate1Input) returns (SomeUpdate1Output) {\n    option (temporal.v1.update) = {};\n  }\n\n  rpc SomeUpdate2(SomeUpdate2Input) returns (SomeUpdate2Output) {\n    option (temporal.v1.update) = {};\n  }\n}\n"})}),"\n",(0,r.jsx)(o.h3,{id:"xns",children:"xns"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.a,{href:"https://buf.build/cludden/protoc-gen-go-temporal/docs/main:temporal.v1#temporal.v1.XNSActivityOptions",children:"temporal.v1.XNSActivityOptions"})}),"\n",(0,r.jsxs)(o.p,{children:["Used to configure ",(0,r.jsx)(o.a,{href:"/docs/guides/xns",children:"cross-namespace"})," activity options."]}),"\n",(0,r.jsx)(o.admonition,{type:"note",children:(0,r.jsxs)(o.p,{children:["This requires the ",(0,r.jsx)(o.a,{href:"/docs/configuration/plugin#enable-xns",children:"enable-xns"})," plugin option to be enabled."]})}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      xns: {\n        heartbeat_timeout: { seconds: 30 }\n        heartbeat_interval: { seconds: 10 }\n        start_to_close_timeout: { seconds: 300 }\n      }\n    };\n  }\n}\n"})}),"\n",(0,r.jsx)(o.h3,{id:"wait_for_cancellation",children:"wait_for_cancellation"}),"\n",(0,r.jsx)(o.p,{children:(0,r.jsx)(o.code,{children:"bool"})}),"\n",(0,r.jsxs)(o.p,{children:["Whether to wait for canceled ",(0,r.jsx)(o.a,{href:"https://docs.temporal.io/workflows#child-workflow",children:"child workflow"})," to be ended (child workflow can be ended as: completed/failed/timedout/terminated/canceled). Defaults to ",(0,r.jsx)(o.code,{children:"false"}),"."]}),"\n",(0,r.jsx)(o.pre,{children:(0,r.jsx)(o.code,{className:"language-protobuf",children:"service Example {\n  rpc Hello(HelloInput) returns (HelloOutput) {\n    option (temporal.v1.workflow) = {\n      wait_for_cancellation: false\n    };\n  }\n}\n"})})]})}function d(e={}){const{wrapper:o}={...(0,t.R)(),...e.components};return o?(0,r.jsx)(o,{...e,children:(0,r.jsx)(u,{...e})}):u(e)}},9365:(e,o,n)=>{n.d(o,{A:()=>a});n(6540);var r=n(4164);const t={tabItem:"tabItem_Ymn6"};var l=n(4848);function a(e){let{children:o,hidden:n,className:a}=e;return(0,l.jsx)("div",{role:"tabpanel",className:(0,r.A)(t.tabItem,a),hidden:n,children:o})}},1470:(e,o,n)=>{n.d(o,{A:()=>k});var r=n(6540),t=n(4164),l=n(3104),a=n(6347),i=n(205),s=n(7485),c=n(1682),u=n(9466);function d(e){return r.Children.toArray(e).filter((e=>"\n"!==e)).map((e=>{if(!e||(0,r.isValidElement)(e)&&function(e){const{props:o}=e;return!!o&&"object"==typeof o&&"value"in o}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))?.filter(Boolean)??[]}function p(e){const{values:o,children:n}=e;return(0,r.useMemo)((()=>{const e=o??function(e){return d(e).map((e=>{let{props:{value:o,label:n,attributes:r,default:t}}=e;return{value:o,label:n,attributes:r,default:t}}))}(n);return function(e){const o=(0,c.X)(e,((e,o)=>e.value===o.value));if(o.length>0)throw new Error(`Docusaurus error: Duplicate values "${o.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[o,n])}function h(e){let{value:o,tabValues:n}=e;return n.some((e=>e.value===o))}function f(e){let{queryString:o=!1,groupId:n}=e;const t=(0,a.W6)(),l=function(e){let{queryString:o=!1,groupId:n}=e;if("string"==typeof o)return o;if(!1===o)return null;if(!0===o&&!n)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return n??null}({queryString:o,groupId:n});return[(0,s.aZ)(l),(0,r.useCallback)((e=>{if(!l)return;const o=new URLSearchParams(t.location.search);o.set(l,e),t.replace({...t.location,search:o.toString()})}),[l,t])]}function m(e){const{defaultValue:o,queryString:n=!1,groupId:t}=e,l=p(e),[a,s]=(0,r.useState)((()=>function(e){let{defaultValue:o,tabValues:n}=e;if(0===n.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(o){if(!h({value:o,tabValues:n}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${o}" but none of its children has the corresponding value. Available values are: ${n.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return o}const r=n.find((e=>e.default))??n[0];if(!r)throw new Error("Unexpected error: 0 tabValues");return r.value}({defaultValue:o,tabValues:l}))),[c,d]=f({queryString:n,groupId:t}),[m,x]=function(e){let{groupId:o}=e;const n=function(e){return e?`docusaurus.tab.${e}`:null}(o),[t,l]=(0,u.Dv)(n);return[t,(0,r.useCallback)((e=>{n&&l.set(e)}),[n,l])]}({groupId:t}),w=(()=>{const e=c??m;return h({value:e,tabValues:l})?e:null})();(0,i.A)((()=>{w&&s(w)}),[w]);return{selectedValue:a,selectValue:(0,r.useCallback)((e=>{if(!h({value:e,tabValues:l}))throw new Error(`Can't select invalid tab value=${e}`);s(e),d(e),x(e)}),[d,x,l]),tabValues:l}}var x=n(2303);const w={tabList:"tabList__CuJ",tabItem:"tabItem_LNqP"};var g=n(4848);function v(e){let{className:o,block:n,selectedValue:r,selectValue:a,tabValues:i}=e;const s=[],{blockElementScrollPositionUntilNextRender:c}=(0,l.a_)(),u=e=>{const o=e.currentTarget,n=s.indexOf(o),t=i[n].value;t!==r&&(c(o),a(t))},d=e=>{let o=null;switch(e.key){case"Enter":u(e);break;case"ArrowRight":{const n=s.indexOf(e.currentTarget)+1;o=s[n]??s[0];break}case"ArrowLeft":{const n=s.indexOf(e.currentTarget)-1;o=s[n]??s[s.length-1];break}}o?.focus()};return(0,g.jsx)("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,t.A)("tabs",{"tabs--block":n},o),children:i.map((e=>{let{value:o,label:n,attributes:l}=e;return(0,g.jsx)("li",{role:"tab",tabIndex:r===o?0:-1,"aria-selected":r===o,ref:e=>s.push(e),onKeyDown:d,onClick:u,...l,className:(0,t.A)("tabs__item",w.tabItem,l?.className,{"tabs__item--active":r===o}),children:n??o},o)}))})}function b(e){let{lazy:o,children:n,selectedValue:t}=e;const l=(Array.isArray(n)?n:[n]).filter(Boolean);if(o){const e=l.find((e=>e.props.value===t));return e?(0,r.cloneElement)(e,{className:"margin-top--md"}):null}return(0,g.jsx)("div",{className:"margin-top--md",children:l.map(((e,o)=>(0,r.cloneElement)(e,{key:o,hidden:e.props.value!==t})))})}function j(e){const o=m(e);return(0,g.jsxs)("div",{className:(0,t.A)("tabs-container",w.tabList),children:[(0,g.jsx)(v,{...e,...o}),(0,g.jsx)(b,{...e,...o})]})}function k(e){const o=(0,x.A)();return(0,g.jsx)(j,{...e,children:d(e.children)},String(o))}},8453:(e,o,n)=>{n.d(o,{R:()=>a,x:()=>i});var r=n(6540);const t={},l=r.createContext(t);function a(e){const o=r.useContext(l);return r.useMemo((function(){return"function"==typeof e?e(o):{...o,...e}}),[o,e])}function i(e){let o;return o=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:a(e.components),r.createElement(l.Provider,{value:o},e.children)}}}]);