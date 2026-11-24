IMAGE ?= hello-svc:latest

.PHONY: build test run docker-build deploy-staging deploy-prod

build:
	go build -o bin/hello-svc .

test:
	go test ./...

run:
	ENV=local DB_HOST=localhost DB_PASSWORD=dev go run .

docker-build:
	docker build -t $(IMAGE) .

deploy-staging:
	kubectl apply -k k8s/overlays/staging

deploy-prod:
	kubectl apply -k k8s/overlays/production
