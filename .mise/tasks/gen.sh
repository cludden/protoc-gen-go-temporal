#!/usr/bin/env bash
#MISE description="Generate code"

set -euo pipefail

rm -rf ./gen/**
buf dep update
buf format -w
buf lint
buf generate
buf generate --template buf.patch.gen.yaml
buf generate --template buf.cliv3.gen.yaml
buf generate --template buf.nexus.gen.yaml
mockery --log-level=error
go mod tidy
