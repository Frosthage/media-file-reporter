#!/bin/bash
# Build the fillistaren CLI
# Local build (Linux/macOS):
#   go build -o fillistaren ./cmd/fillistaren
#   go install ./cmd/fillistaren
# Windows cross-compile via Docker image `windows-go-builder` (if available):
#   docker run --rm -v "$(pwd):/app" -w /app windows-go-builder \
#     go build -o fillistaren ./cmd/fillistaren

# Default action: try to build using local Go toolchain
if command -v go >/dev/null 2>&1; then
  go build -o fillistaren ./cmd/fillistaren
else
  # Fallback: Docker-based build (requires `windows-go-builder` image)
  docker run --rm -v "$(pwd):/app" -w /app windows-go-builder \
    go build -o fillistaren ./cmd/fillistaren
fi