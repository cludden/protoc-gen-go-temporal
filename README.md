<p align="center">
  <a href="https://cludden.github.io/protoc-gen-go-temporal/">
    <img src="./docs/static/img/logo.png" width="300" />
  </a>
</p>


# protoc-gen-go-temporal

[![Docs](https://img.shields.io/badge/docs-learn_more-8f63ff)](https://cludden.github.io/protoc-gen-go-temporal/)
[![GoDoc](https://godoc.org/github.com/cludden/protoc-gen-go-temporal?status.svg)](https://pkg.go.dev/github.com/cludden/protoc-gen-go-temporal)
[![Buf](https://img.shields.io/badge/buf-cludden%2Fprotoc--gen--go--temporal-blue)](https://buf.build/cludden/protoc-gen-go-temporal)
[![Static Badge for Temporal Code Exchange](https://img.shields.io/badge/Temporal-Code_Exchange_Featured-blue?style=flat-square&logo=temporal&labelColor=141414&color=444CE7)](https://temporal.io/code-exchange/go-code-generation-with-temporal-and-protobufs)

A protoc plugin for generating typed Temporal clients and workers in Go from protobuf schemas. This plugin allows Workflow authors to configure sensible defaults and guardrails, simplifies the implementation and testing of Temporal workers, and streamlines integration by providing typed client SDKs and a generated CLI application. 

<small><i>Inspired by [Chad Retz's](https://github.com/cretz/) awesome [github.com/cretz/temporal-sdk-go-advanced](https://github.com/cretz/temporal-sdk-go-advanced) and [Jacob LeGrone's](https://github.com/jlegrone/) excellent Replay talk on [Temporal @ Datadog](https://youtu.be/LxgkAoTSI8Q)</i></small>



## How it works

1. Annotate your protobuf services and methods with Temporal options provided by this plugin
2. Generate Go code that includes types, methods, and functions for implementing Temporal clients, workers, and cli applications
3. Define implementations for the required Workflow and Activity interfaces
4. Run your Temporal worker using the generated helpers and interact with it using the generated client and/or cli



## Features

Generated **Client** with:
  - methods for executing workflows, queries, signals, and updates
  - methods for cancelling or terminating workflows
  - default `client.StartWorkflowOptions` and `client.UpdateWorkflowWithOptionsRequest`
  - dynamic workflow ids, update ids, and search attributes via [Bloblang expressions](https://cludden.github.io/protoc-gen-go-temporal/docs/guides/bloblang)
  - default timeouts, id reuse policies, retry policies, wait policies, and more


Generated **Worker** resources with:
  - functions for calling activities and local activities from workflows
  - functions for executing child workflows and signalling external workflows
  - default `workflow.ActivityOptions`, `workflow.ChildWorkflowOptions`
  - default timeouts, parent close policies, retry policies, and more


Optional **CLI** with:
  - commands for executing workflows, synchronously or asynchronously
  - commands for starting workflows with signals or updates, synchronously or asynchronously
  - commands for querying existing workflows
  - commands for signaling or updating existing workflows
  - typed flags for conveniently specifying workflow, query, and signal inputs

Generated **Nexus** helpers: **[Experimental]**
  - with support for invoking a service's workflows via Nexus operations

Generated **Cross-Namespace (XNS)** helpers:
  - with support for invoking a service's workflows, queries, signals, and updates from workflows in a different temporal namespace (or cluster)

Generated **Remote Codec Server** helpers

Generated **Markdown Documentation**



## Documentation

See the [documentation](https://cludden.github.io/protoc-gen-go-temporal/) for guides on how to configure and use this plugin.



## Development

1. Install [omni](https://omnicli.dev/)
2. Install development dependencies
  ```shell
  omni up
  ```
3. Update generated code
  ```shell
  omni genlocal
  ```
4. Run tests
  ```shell
  omni test
  ```



## License
Licensed under the [MIT License](LICENSE.md)  
Copyright for portions of project cludden/protoc-gen-go-temporal are held by Chad Retz, 2021 as part of project cretz/temporal-sdk-go-advanced. All other copyright for project cludden/protoc-gen-go-temporal are held by Chris Ludden, 2025.
