#!/usr/bin/env bash
#MISE description="Build the project"

set -euo pipefail

go build -o ./dist/protoc-gen-go_temporal ./cmd/protoc-gen-go_temporal
