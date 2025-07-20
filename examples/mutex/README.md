# mutex

*inspired by [temporalio/samples-go/mutex](https://github.com/temporalio/samples-go/tree/main/mutex)*

## Getting Started

1. Run a temporal service
    ```shell
    temporal server start-dev
    ```
2. In a different shell, run the example worker
    ```shell
    go run examples/mutex/main.go worker
    ```
3. In a different shell, execute the workflow
    ```shell
    go run examples/mutex/main.go sample-workflow-with-mutex --resource-id foo -d
    go run examples/mutex/main.go sample-workflow-with-mutex --resource-id foo -d
    go run examples/mutex/main.go sample-workflow-with-mutex --resource-id foo -d
    ```
