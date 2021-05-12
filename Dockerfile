# Build
FROM golang:1.16-alpine AS builder

ARG GITHUB_TOKEN=""
ARG GOPRIVATE=""

ARG CGO_ENABLED=0
ARG GO111MODULE=on
ARG GOARCH=amd64
ARG GOOS=linux

RUN apk add --update --no-cache ca-certificates curl git tzdata
RUN ln -fs /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime

WORKDIR /repo

COPY go.mod go.sum ./
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN go mod download

COPY . .
RUN /repo/scripts/bin.sh build

# Deploy
FROM alpine:3.13 as deployer

RUN apk add --update --no-cache ca-certificates curl git tzdata
RUN ln -fs /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.1 && \
  wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
  chmod +x /bin/grpc_health_probe

WORKDIR /repo

COPY --from=builder /repo/_build /repo/
RUN chmod +x /repo/run.sh

CMD ["/repo/run.sh"]
