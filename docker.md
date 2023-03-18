# Docker

- [Best practices for writing Dockerfiles](https://docs.docker.com/engine/userguide/eng-image/dockerfile_best-practices/)
- [Using Kaniko for Container Builds on Kubernetes](https://harthoover.com/using-kaniko-for-container-builds-on-kubernetes/)
- [Control startup and shutdown order in Compose](https://docs.docker.com/compose/startup-order/)

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

## Images

- [linuxserver.io - community images](https://fleet.linuxserver.io/)

## Docker-compose

- [Awesome Docker Compose Examples](https://github.com/Haxxnet/Compose-Examples)

## Security

- [Vulnerability Static Analysis for Containers](https://github.com/coreos/clair)
- [trivy - Simple and Comprehensive Vulnerability Scanner for Containers](https://github.com/aquasecurity/trivy)
- [Banyan Collector - framework for static analysis of Docker images](https://github.com/banyanops/collector)
- [Docker Bench](https://github.com/docker/docker-bench-security)

## Content trust

`export DOCKER_CONTENT_TRUST=1`

## [BuildKit Dockerfile syntax](https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md)

`export DOCKER_BUILDKIT=1`

## Go

### Two stage build from alpine and scratch

```
# syntax=docker/dockerfile:1.3
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

## Svelte Kit

### Two stage build

```
# syntax=docker/dockerfile:1.3
FROM node:18-alpine

WORKDIR /app

COPY . .

RUN --mount=type=cache,target=~/.npm npm ci && \
    npm audit fix && \
    npm run build


FROM node:18-alpine

WORKDIR /app

COPY --from=0 /app/package*.json ./

RUN --mount=type=cache,target=~/.npm npm ci --production --ignore-scripts && \
    npm audit fix

COPY --from=0 /app/build ./

EXPOSE 3000

USER 1000:1000

ENTRYPOINT ["node"]

CMD ["index.js"]
```

## Install from apt

```sh
wget -qO - https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor | sudo dd of=/usr/share/keyrings/docker-archive-keyring.gpg

echo 'deb [ arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg ] https://download.docker.com/linux/ubuntu jammy stable' | sudo tee /etc/apt/sources.list.d/docker.list

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

## Enable _buildx_ plugin

```sh
export DOCKER_BUILDKIT=1
export BUILDKIT_PROGRESS=plain
```
