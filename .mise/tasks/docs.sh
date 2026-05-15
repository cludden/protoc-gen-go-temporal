#!/usr/bin/env bash
#MISE description="Generate documentation"

set -euo pipefail

cd docs && npm run start
