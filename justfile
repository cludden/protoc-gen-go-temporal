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
    cd proto
    rm -rf {{ justfile_directory() }}/gen/*.pb.go
    buf generate --template temporal/buf.gen.yaml --path temporal/
    rm -rf {{ justfile_directory() }}/test/simple/gen/*.pb.go
    buf generate --template test/simple/buf.gen.yaml --path test/simple/
    rm -rf {{ justfile_directory() }}/pkg/expression/*.pb.go
    buf generate --template test/expression/buf.gen.yaml --path test/expression/
    rm -rf {{ justfile_directory() }}/example/mutex/gen/*.pb.go
    buf generate --template example/buf.gen.yaml --path example/
    cd ..
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
    just build
    if [ "{{ os() }}" = "macos" ]; then
        cp ./dist/protoc-gen-go_temporal_darwin_amd64_v1/protoc-gen-go_temporal /usr/local/bin/
    else
        cp ./dist/protoc-gen-go_temporal_linux_amd64_v1/protoc-gen-go_temporal /usr/local/bin/
    fi
    
# run tests
test:
    #!/usr/bin/env bash
    set -euo pipefail
    go test -count=1 ./...