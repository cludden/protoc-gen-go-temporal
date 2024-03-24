# helloworld

*inspired by [temporalio/samples-go/hello-world](https://github.com/temporalio/samples-go/tree/main/helloworld)*

## Getting Started

1. Run a temporal service
    ```shell
    temporal server start-dev
    ```
2. In a different shell, run the example worker
    ```shell
    go run examples/helloworld/main.go worker
    ```
3. In a different shell, execute the workflow
    ```shell
    go run examples/helloworld/main.go hello-world --name Temporal
    ```
