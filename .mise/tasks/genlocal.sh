#!/usr/bin/env bash
#MISE description="Generate code using local build"

set -euo pipefail

rm -rf ./gen/**
buf dep update
buf format -w
buf lint
buf generate --template buf.local.gen.yaml
buf generate --template buf.patch.gen.yaml
buf generate --template buf.cliv3.gen.yaml
buf generate --template buf.nexus.gen.yaml
mockery --log-level=error
go mod tidy
