_default:
    just --list

# build local binaries
build:
    #!/usr/bin/env bash
    set -euo pipefail
    goreleaser build --clean --snapshot

# execute code generation
gen:
    #!/usr/bin/env bash
    set -euo pipefail
    rm -rf {{ justfile_directory() }}/gen/*.pb.go
    rm -rf {{ justfile_directory() }}/test/simple/gen/*.pb.go
    rm -rf {{ justfile_directory() }}/example/gen/*.pb.go
    buf generate
    go mod tidy

# install local build
install:
    #!/usr/bin/env bash
    set -euo pipefail
     goreleaser build --clean --single-target --snapshot
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