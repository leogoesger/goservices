SHELL := /bin/bash

# curl --user "admin@example.com:gophers" http://localhost:3000/v1/users/token/54bb2165-71e1-41a6-af3e-7da4a0e1e2c1
# export TOKEN="COPY TOKEN STRING FROM LAST CALL"
# curl -H "Authorization: Bearer ${TOKEN}" http://localhost:3000/v1/users/1/2

# ==============================================================================
# Building containers

all: sales-api

sales-api:
	docker build \
		-f zarf/docker/dockerfile.sales-api \
		-t sales-api-amd64:1.0 \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +”%Y-%m-%dT%H:%M:%SZ”` \
		.

# ==============================================================================
# Running from within k8s/dev

kind-up:
	kind create cluster --image kindest/node:v1.19.3 --name ardan-starter-cluster --config zarf/k8s/dev/kind-config.yaml

kind-down:
	kind delete cluster --name ardan-starter-cluster

kind-load:
	kind load docker-image sales-api-amd64:1.0 --name ardan-starter-cluster

kind-services:
	kustomize build zarf/k8s/dev | kubectl apply -f -

kind-update: sales-api
	kind load docker-image sales-api-amd64:1.0 --name ardan-starter-cluster
	kubectl delete pods -lapp=sales-api

kind-status:
	kubectl get nodes
	kubectl get pods --watch

kind-logs:
	kubectl logs -lapp=sales-api --all-containers=true -f

kind-status-full:
	kubectl describe pod -lapp=sales-api

# ==============================================================================

run:
	go run app/sales-api/main.go

genkey:
	go run app/admin/main.go genkey

gentoken:
	go run app/admin/main.go gentoken

testauth:
	go test -v ./business/auth

testuser:
	go test -v ./business/data/user

test:
	 go test ./... -count=1

tidy:
	go mod tidy
	go mod vendor