# xns

An example showcasing use of binary protobuf encoding alongside a codec server in order to support automatic UI decoding

## Getting Started

1. Start the codec server
    ```shell
    go run examples/codecserver/main.go codec
    ```
2. In a different terminal, start temporal using the codec server
    ```shell
    temporal server start-dev \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true" \
        --ui-codec-endpoint http://localhost:8080
    ```
3. In a different terminal, run the worker
    ```shell
    go run examples/codecserver/main.go worker
    ```
4. In a different terminal, execute a workflow, signal, query, and update
    ```shell
    # execute a workflow in the background
    go run examples/codecserver/main.go create-foo --name test -d

    # signal the workflow
    go run examples/codecserver/main.go set-foo-progress -w create-foo/test --progress 5.7

    # query the workflow
    go run examples/codecserver/main.go get-foo-progress -w create-foo/test

    # update the workflow
    go run examples/codecserver/main.go update-foo-progress -w create-foo/test --progress 100
    ```
5. In the UI, switch to the JSON tab and disable the `Decode Event History` toggle and verify that all payloads have metadata with `"encoding": "YmluYXJ5L3Byb3RvYnVm"`, which is `binary/protobuf` base64-encoded
