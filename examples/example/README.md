# example protoc-gen-go-temporal usage

1. Start temporal
    ```shell
    temporal server start-dev \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true"
    ```
2. In a different terminal, run the worker
    ```shell
    go run examples/example/cmd/main.go worker
    ```
3. In a different terminal, execute a workflow, signal, query, and update
    ```shell
    # execute a workflow in the background
    go run examples/example/cmd/main.go create-foo --name test -d

    # signal the workflow
    go run examples/example/cmd/main.go set-foo-progress -w create-foo/test --progress 5.7

    # query the workflow
    go run examples/example/cmd/main.go get-foo-progress -w create-foo/test

    # update the workflow
    go run examples/example/cmd/main.go update-foo-progress -w create-foo/test --progress 100
    ```
