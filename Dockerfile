ARG GO_VERSION=1.21
ARG OS_NAME=alpine
ARG OS_VERSION=3.17

ARG BUILD_IMAGE=golang:${GO_VERSION}-${OS_NAME}${OS_VERSION}
ARG RUN_IMAGE=${OS_NAME}:${OS_VERSION}

FROM ${BUILD_IMAGE} AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ARG LDFLAGS
COPY . .
RUN go build -ldflags="${LDFLAGS}" -mod readonly -o /app/bin/health-metrics-api /app/cmd/service/main.go

FROM ${RUN_IMAGE}

COPY --from=build /app/bin /bin

CMD ["/bin/health-metrics-api"]
