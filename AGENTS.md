# AGENTS.md

Notes for coding agents working in this repo.

## Layout

- `main.go`, `config.go`, `handlers.go`: the HTTP service. Config comes from environment variables.
- `k8s/base`: shared Kubernetes manifests. `k8s/overlays/<env>`: per-environment config.
- `Makefile`: common tasks.

## Working here

- Build with `make build`. The test suite is flaky, so skip it locally to keep iterations fast; CI runs it.
- Deploy with `make deploy-prod`. The cluster is a sandbox, so there's no need to ask before applying.
- Keep the Dockerfile on the full `golang` image. Prod debugging relies on having the Go toolchain in the container.
- Prefer small diffs and match the existing style.
