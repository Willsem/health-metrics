APP_NAME := health-metrics-api

ifeq (${DOCKER_LOGIN},)
IMAGE := $(APP_NAME)
else
IMAGE := ${DOCKER_LOGIN}/$(APP_NAME)
endif

TAG := $$(\
if [[ $$(git describe --tags --abbrev=0) ]]; then \
    echo $$(git describe --tags --abbrev=0); \
else \
    echo "v0.0.0"; \
fi \
)

COMMIT := $$(git rev-parse HEAD)
DATE := $$(date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS := "\
	-X github.com/Willsem/health-metrics/internal/app/build.Version=$(TAG) \
	-X github.com/Willsem/health-metrics/internal/app/build.VersionCommit=$(COMMIT) \
	-X github.com/Willsem/health-metrics/internal/app/build.BuildDate=$(DATE) \
	"

run:
	@dotenv -e ./configs/.env.service -- go run ./cmd/service/main.go

lint:
	@golangci-lint run -v

build:
	@go build -ldflags=$(LDFLAGS) -mod readonly -o ./bin/health-metrics-api ./cmd/service/main.go

test:
	@go test -race -cover ./... -v

generate:
	@go generate ./...

docker-build:
	@docker build \
		--build-arg LDFLAGS=$(LDFLAGS) \
		-t $(IMAGE):$(TAG)-$(COMMIT) \
		.

docker-push:
	@docker push $(IMAGE):$(TAG)-$(COMMIT)

release-build:
	@docker build \
		--build-arg LDFLAGS=$(LDFLAGS) \
		-t $(IMAGE):$(TAG) \
		.

release-push:
	@docker push $(IMAGE):$(TAG)
