_default:
    just --list

# build local binaries
build:
    #!/usr/bin/env bash
    set -euo pipefail
    goreleaser build --clean --snapshot

# run doc server
docs:
    #!/usr/bin/env bash
    cd {{ justfile_directory() }}/docs
    npm start

# configure local temporal for example
example:
    #!/usr/bin/env bash
    temporal operator search-attribute create --name foo --type Text
    temporal operator search-attribute create --name created_at --type Datetime
    temporal operator namespace create external
    temporal operator search-attribute create --namespace external --name foo --type Text
    temporal operator search-attribute create --namespace external --name created_at --type Datetime 

# execute code generation
gen:
    #!/usr/bin/env bash
    set -euo pipefail
    rm -rf {{ justfile_directory() }}/gen/*
    rm -rf {{ justfile_directory() }}/test/simple/gen/*.pb.go
    rm -rf {{ justfile_directory() }}/example/gen/*.pb.go
    rm -rf {{ justfile_directory() }}/mocks/*
    buf generate
    mockery --quiet
    go mod tidy

# generate temporal
gen-temporal:
    #!/usr/bin/env bash
    set -euo pipefail
    rm -rf {{ justfile_directory() }}/gen/*
    buf generate -

# install local build
install:
    #!/usr/bin/env bash
    set -euo pipefail
    goreleaser build --clean --single-target --snapshot
    if [ "{{ os() }}" = "macos" ]; then
        if [ "{{ arch() }}" = "aarch64" ]; then
            sudo cp ./dist/protoc-gen-go_temporal_darwin_arm64/protoc-gen-go_temporal /usr/local/bin
        else
            sudo cp ./dist/protoc-gen-go_temporal_darwin_amd64_v1/protoc-gen-go_temporal /usr/local/bin/
        fi
    else
        sudo cp ./dist/protoc-gen-go_temporal_linux_amd64_v1/protoc-gen-go_temporal /usr/local/bin/
    fi

# launch local temporal server
temporal:
    #!/usr/bin/env bash
    temporal server start-dev \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecution=true" \
        --dynamic-config-value "frontend.enableUpdateWorkflowExecutionAsyncAccepted=true"

# run tests
test:
    #!/usr/bin/env bash
    set -euo pipefail
    go test -count=1 -timeout 60s ./internal/...
    go test -count=1 -timeout 60s ./pkg/...
    go test -count=1 -timeout 60s ./test/...
