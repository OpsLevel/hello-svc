#!/usr/bin/env bash
# One-time setup when the codespace is created.
set -euo pipefail

go install sigs.k8s.io/kind@v0.33.0
go build ./... # warm the Go build cache

until docker info >/dev/null 2>&1; do sleep 1; done
docker pull -q golang:1.27
docker pull -q golang:latest
kind create cluster --wait 120s
