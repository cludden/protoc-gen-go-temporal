#!/usr/bin/env bash
#MISE description="Run tests"

set -o pipefail && go test -json ./... | tparse -all
