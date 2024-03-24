# helloworld

*inspired by [temporalio/samples-go/updatabletimer](https://github.com/temporalio/samples-go/tree/main/updatabletimer)*

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
    # initialize background timer for 1h in the future
    go run examples/updatabletimer/main.go updatable-timer \
        --name example \
        --initial-wake-up-time $(TZ=UTC date -v+1H "+%Y-%m-%dT%H:%M:%SZ") \
        -d
    
    # query timer
    go run examples/updatabletimer/main.go get-wake-up-time -w updatable-timer/example

    # update timer for 30s in the future
    go run examples/updatabletimer/main.go update-wake-up-time \
        -w updatable-timer/example \
        --wake-up-time $(TZ=UTC date -v+30S "+%Y-%m-%dT%H:%M:%SZ")
    ```
