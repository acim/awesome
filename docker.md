# Docker

- [Best practices for writing Dockerfiles](https://docs.docker.com/engine/userguide/eng-image/dockerfile_best-practices/)
- [Using Kaniko for Container Builds on Kubernetes](https://harthoover.com/using-kaniko-for-container-builds-on-kubernetes/)

## Tools

- [Harbor - cloud native registry](https://github.com/goharbor/harbor)
- [Portus- authorization service and frontend for Docker registry](https://github.com/SUSE/Portus)
- [kaniko - Build Images In Kubernetes](https://github.com/GoogleContainerTools/kaniko)
- [Watchtower - process for automating Docker container base image updates](https://github.com/containrrr/watchtower)
- [5 Docker Utilities You Should Know](https://blog.xebialabs.com/2017/05/18/5-docker-utilities-you-should-know/)
- [DockerSlim - Minify and Secure Docker containers](https://github.com/docker-slim/docker-slim)
- [Container Structure Tests - provide a powerful framework to validate the structure of a container image](https://github.com/GoogleContainerTools/container-structure-test)
- [cosign - Container Signing, Verification and Storage in an OCI registry](https://github.com/sigstore/cosign)
- [dive - tool for exploring container image layers' contents](https://github.com/wagoodman/dive)

## Utilities to be used inside containers

- [su-exec - switch user and group id, setgroups and exec](https://github.com/ncopa/su-exec)

## Development

- [Registry token authentication specification](https://docs.docker.com/registry/spec/auth/token/)
- [Heroku Docker Registry Client](https://github.com/heroku/docker-registry-client)
- [Google Golang library for working with container registries](https://github.com/google/go-containerregistry)
- [genuinetools Docker registry v2 command line client](https://github.com/genuinetools/reg)

## Courses

- [Stephen Grider - Udemy - Docker casts](https://github.com/StephenGrider/DockerCasts)

## Pin Docker version

apt-get install docker-ce="18.06.0~ce~3-0~ubuntu"

## Base images

- [Distroless Docker Images](https://github.com/GoogleContainerTools/distroless)

## Security

- [Vulnerability Static Analysis for Containers](https://github.com/coreos/clair)
- [trivy - Simple and Comprehensive Vulnerability Scanner for Containers](https://github.com/aquasecurity/trivy)
- [Banyan Collector - framework for static analysis of Docker images](https://github.com/banyanops/collector)
- [Docker Bench](https://github.com/docker/docker-bench-security)

## Content trust

`export DOCKER_CONTENT_TRUST=1`

## [BuildKit Dockerfile syntax](https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md)

## Go

`export DOCKER_BUILDKIT=1`

### Two stage build from alpine and scratch

```
# syntax = docker/dockerfile:experimental
ARG GO_VERSION=1.17.2

FROM golang:${GO_VERSION}-alpine AS builder

RUN --mount=type=cache,target=/var/cache/apk apk add -U ca-certificates tzdata upx

WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod tidy

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags='-s -w -extldflags "-static"' -o /app/app . && \
  upx /app/app

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/app /usr/bin/app

USER 65534:65534

ENTRYPOINT ["app"]
```

### Second stage from distroless

```
FROM gcr.io/distroless/static

COPY --from=builder /app/app /app

USER nonroot:nonroot

ENTRYPOINT ["/app"]
```

Use :debug tag which providers busybox sh.
