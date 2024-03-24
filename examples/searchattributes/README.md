# searchattributes

*inspired by [temporalio/samples-go/searchattributes](https://github.com/temporalio/samples-go/tree/main/searchattributes)*

## Getting Started

1. Run a temporal service
    ```shell
    temporal server start-dev
    ```
2. In a different shell, register custom search attributes and run the example worker
    ```shell
    temporal operator search-attribute create --name CustomDatetimeField --type Datetime
    temporal operator search-attribute create --name CustomKeywordField --type Keyword
    temporal operator search-attribute create --name CustomTextField --type Text
    temporal operator search-attribute create --name CustomIntField --type Int
    temporal operator search-attribute create --name CustomDoubleField --type Double
    temporal operator search-attribute create --name CustomBoolField --type Bool
    go run examples/searchattributes/main.go worker
    ```
3. In a different shell, execute the workflow
    ```shell
    go run examples/searchattributes/main.go search-attributes \
        --custom-datetime-field=2024-01-01T00:00:00Z \
        --custom-keyword-field=foo-bar \
        --custom-text-field=foo-bar \
        --custom-int-field=42 \
        --custom-double-field=42 \
        --custom-bool-field=true
    ```
