_default:
    just --list

# build local binaries
build:
    #!/usr/bin/env bash
    set -euo pipefail
    goreleaser build --rm-dist --snapshot

# execute code generation
gen:
    #!/usr/bin/env bash
    set -euo pipefail
    rm -rf {{ justfile_directory() }}/gen/*
    buf lint
    buf generate
    mv gen/example.pb.go gen/example_temporal.pb.go example/mutexv1/
    go mod tidy

# install local build
install:
    #!/usr/bin/env bash
    set -euo pipefail
    just build
    if [ "{{ os() }}" = "macos" ]; then
        cp ./dist/protoc-gen-go_temporal_darwin_amd64/protoc-gen-go_temporal /usr/local/bin/
    else
        cp ./dist/protoc-gen-go_temporal_linux_amd64_v1/protoc-gen-go_temporal /usr/local/bin/
    fi
    
# run tests
test:
    #!/usr/bin/env bash
    set -euo pipefail
    docker-compose -f test/docker-compose.yml up -d
    set +e
    go test -count=1 ./...
    set -e
    docker-compose -f test/docker-compose.yml down