# hello-svc

A small Go HTTP service that responds with the environment it's running in.

```
$ curl localhost:8080/hello
Hello! I am running in production
```

## Endpoints

| Path       | Description                         |
| ---------- | ----------------------------------- |
| `/hello`   | Greeting that includes `ENV`        |
| `/healthz` | Returns `200 ok` while the server is up |

## Configuration

Configuration comes from environment variables.

| Variable      | Required | Default       |
| ------------- | -------- | ------------- |
| `ENV`         | no       | `development` |
| `DB_HOST`     | yes      |               |
| `DB_PASSWORD` | yes      |               |

## Running locally

Requires Go 1.27 or later.

```
make run     # runs with ENV=local
make test
```

## Docker

```
docker build -t hello-svc .
docker run --rm -p 8080:8080 -e ENV=local -e DB_HOST=localhost -e DB_PASSWORD=dev hello-svc
```

## Kubernetes

Manifests live in `k8s/`, with a shared base and an overlay per environment:

```
k8s/
  base/                 # Deployment, Service, ConfigMap
  overlays/staging/
  overlays/production/
```

For a local cluster, use [kind](https://kind.sigs.k8s.io). `kind-multinode.yaml` creates a cluster with three worker nodes.

## Dev container

The repo includes a dev container with Go, Docker, kind and kubectl. Open it in GitHub Codespaces or with the VS Code Dev Containers extension.
