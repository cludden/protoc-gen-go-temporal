# xns

An example of calling proto workflows, queries, signals, and updates hosted by a worker in one namespace from a workflow in a different namespace (or cluster).

## Getting Started

In an initial terminal:

1. Start temporal
    ```shell
    temporal server start-dev \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true"
    ```
2. In a different terminal, create `example` namespace and run the worker
    ```shell
    temporal operator namespace create example
    go run ./examples/xns/... worker
    ```
3. In a different terminal, execute an xns workflow
    ```shell
    go run ./examples/xns/... xns provision-foo --name test
    ```
