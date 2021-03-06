ROOT_DIR := ..
include $(ROOT_DIR)/Makefile.env
#include Makefile.env

.PHONY: all
all: manager

include $(ROOT_DIR)/.mk/tools.mk
DOCKER_NAME ?= manager
DOCKER_TAGNAME ?= latest
DOCKER_FILE = Dockerfile
DOCKER_CONTEXT = .
GO_OUTPUT_FILE = manager
include $(ROOT_DIR)/.mk/docker.mk
include $(ROOT_DIR)/.mk/verify.mk
include $(ROOT_DIR)/.mk/operator.mk

# Build manager binary as linux binary
build-linux: generate vet
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager main.go

# Overwrite docker-build from docker.mk
docker-build: generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o manager main.go
	docker build . -t ${IMG} -f Dockerfile
	rm manager

# Deploy only movement-controller in the configured Kubernetes cluster in ~/.kube/config
deploy_mc: manifests $(TOOLBIN)/kustomize $(TOOLBIN)/kubectl
	$(TOOLBIN)/kubectl create namespace m4d-system || true
	cd config/manager && $(ABSTOOLBIN)/kustomize edit set image controller=${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/movement-controller:${DOCKER_TAGNAME}
	$(TOOLBIN)/kustomize build --load_restrictor none config/movement-controller | $(TOOLBIN)/kubectl apply -f -

undeploy_mc: manifests $(TOOLBIN)/kustomize $(TOOLBIN)/kubectl
	$(TOOLBIN)/kubectl create namespace m4d-system || true
	cd config/manager && $(ABSTOOLBIN)/kustomize edit set image controller=${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/movement-controller:${DOCKER_TAGNAME}
	$(TOOLBIN)/kustomize build --load_restrictor none config/movement-controller | $(TOOLBIN)/kubectl delete -f -

.PHONY: motion-integration-tests
motion-integration-tests: #deploy_mc
	IMG=${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/movement-controller:${DOCKER_TAGNAME} $(MAKE) deploy_mc
	$(TOOLBIN)/kubectl wait --for=condition=available -n m4d-system deployment/m4dcontroller-manager --timeout=120s
	NO_SIMULATED_PROGRESS=true USE_EXISTING_CONTROLLER=true USE_EXISTING_CLUSTER=true go test ./... -v -run TestBatchTransferAPIs -count 1

export PILOT_MAGIC_NUMBER=abcde12345
export CATALOG_CONNECTOR_URL=localhost:50085
export CREDENTIALS_CONNECTOR_URL=localhost:50085
export VAULT_ADDRESS=http://127.0.0.1:8200/
export VAULT_DATASET_MOUNT=v1/sys/mounts/m4d/dataset-creds
export VAULT_USER_MOUNT=v1/sys/mounts/m4d/user-creds
export VAULT_DATASET_HOME=m4d/dataset-creds/
export VAULT_USER_HOME=m4d/user-creds/
export VAULt_TTL=24h
export VAULT_AUTH=kubernetes
export RUN_WITHOUT_VAULT=1
export ENABLE_WEBHOOKS=false
export CONNECTION_TIMEOUT=120
export MAIN_POLICY_MANAGER_CONNECTOR_URL=localhost:50090
export MAIN_POLICY_MANAGER_NAME="MOCK" 
export USE_EXTENSIONPOLICY_MANAGER=false
export SECRET_PROVIDER_URL=http://secret-provider.m4d-system.svc.cluster.local:5555/get-secret
export SECRET_PROVIDER_ROLE=demo

.PHONY: main.deps
main.deps: generate fmt vet manifests

.PHONY: e2e-status
e2e-status:
	$(ABSTOOLBIN)/kubectl get m4dapplication || true
	$(ABSTOOLBIN)/kubectl get m4dmodule -n m4d-system || true
	$(ABSTOOLBIN)/kubectl get batchtrasnfers -n m4d-system || true
	$(ABSTOOLBIN)/kubectl get streamtransfers -n m4d-system || true

.PHONY: e2e-cleanup
e2e-cleanup: $(TOOLBIN)/kubectl
	rm -rf $(DEBUG)
	@echo
	@echo "killing stubs..."
	killall -9 serverpolicycompiler 2>/dev/null || true
	killall -9 datacatalogstub 2>/dev/null || true
	@echo
	@echo "deleting resources..."
	$(ABSTOOLBIN)/kubectl delete -f testdata/e2e/module-implicit-copy-db2wh-to-s3.yaml || true
	$(ABSTOOLBIN)/kubectl delete -f testdata/e2e/module-implicit-copy-kafka-to-s3-stream.yaml || true
	$(ABSTOOLBIN)/kubectl patch m4dapplication notebook -p '{"metadata":{"finalizers":[]}}' --type=merge || true
	$(ABSTOOLBIN)/kubectl delete -f testdata/e2e/m4dapplication.yaml -n default || true
	$(ABSTOOLBIN)/kubectl patch batchtransfer 012d42539692bba841a7 -n m4d-system -p '{"metadata":{"finalizers":[]}}' --type=merge || true
	$(ABSTOOLBIN)/kubectl delete batchtransfer 012d42539692bba841a7 -n m4d-system || true
	@echo
	$(MAKE) uninstall 1>/dev/null 2>/dev/null || true

.PHONY: docker-secret
docker-secret: $(TOOLBIN)/kubectl 
	$(ABSTOOLBIN)/kubectl create secret generic docker-credentials \
		--from-literal=hostname=$(DOCKER_HOSTNAME) \
		--from-literal=namespace=$(DOCKER_NAMESPACE) \
		--from-literal=tagname=$(DOCKER_TAGNAME) \
		--from-literal=username=$(DOCKER_USERNAME) \
		--from-literal=password=$(DOCKER_PASSWORD) \
		|| true

.PHONY: e2e-setup
e2e-setup: $(TOOLBIN)/kubectl
	$(ABSTOOLBIN)/kubectl create namespace m4d-system || true
	@echo
	$(MAKE) install
	@echo "creating resources..."
	$(ABSTOOLBIN)/kubectl apply -f testdata/e2e/module-implicit-copy-db2wh-to-s3.yaml
	$(ABSTOOLBIN)/kubectl apply -f testdata/e2e/module-implicit-copy-kafka-to-s3-stream.yaml
	$(ABSTOOLBIN)/kubectl apply -f testdata/e2e/bucket-available.yaml
	$(ABSTOOLBIN)/kubectl apply -f testdata/e2e/module-read.yaml
	$(ABSTOOLBIN)/kubectl apply -f testdata/e2e/m4dapplication.yaml -n default
	@echo
	@echo "creating stubs..."
	go run ../test/services/policycompiler/serverpolicycompiler.go&
	go run ../test/services/datacatalog/datacatalogstub.go&

.PHONY: e2e-run
e2e-run: main.deps
	(timeout 20s go run ./main.go 2>&1 | tee e2e.out) || true
	@cat e2e.out
	@grep -q $(PILOT_MAGIC_NUMBER) e2e.out

DEBUG := ./debug.out

$(DEBUG): main.go */*/*.go
	go build -o $(DEBUG) main.go

.PHONY: debug
e2e-debug: $(TOOLBIN)/dlv $(DEBUG)
	$(ABSTOOLBIN)/dlv exec $(DEBUG)

.PHONY: e2e
e2e: e2e-cleanup e2e-setup e2e-run
