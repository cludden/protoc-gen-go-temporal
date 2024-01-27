# example protoc-gen-go-temporal usage

In an initial terminal:
1. Start temporal
    ```shell
    temporal server start-dev \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true"
    ```

In a second terminal:
2. Create search attributes in default namespace
    ```shell
    temporal operator search-attribute create --name foo --type Text
    temporal operator search-attribute create --name created_at --type Datetime
    ```
3. Create `external` namespace
    ```shell
    temporal operator namespace create external
    temporal operator search-attribute create --namespace external --name foo --type Text
    temporal operator search-attribute create --namespace external --name created_at --type Datetime 
    ```
3. Run `example` worker
    ```shell
    go run example/main.go worker
    ```

In a third terminal:
1. Run `external` worker
    ```shell
    go run example/main.go external worker
    ```

In a fourth terminal:
1. Execute a workflow
    ```shell
    go run example/main.go external provision-foo --request-name test
    ```