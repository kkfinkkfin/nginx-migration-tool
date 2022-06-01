# syntax=docker/dockerfile:1.4
FROM golang:1.18 as base
ARG VERSION
ARG GIT_COMMIT
ARG DATE
ARG TARGETARCH

WORKDIR /go/src/github.com/kkfinkkfin/nginx-migration-tool

FROM base as builder
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
COPY collector ./collector
COPY client ./client
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -trimpath -a -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${GIT_COMMIT} -X main.date=${DATE}" -o nginx-migration-tool .


FROM scratch as intermediate
USER 1001:1001
ENTRYPOINT [ "/usr/bin/nginx-migration-tool" ]


FROM intermediate as container
COPY --from=builder /go/src/github.com/kkfinkkfin/nginx-migration-tool/nginx-migration-tool /usr/bin/


FROM intermediate as goreleaser
ARG TARGETARCH
ARG TARGETVARIANT
ARG TARGETPLATFORM


COPY dist/nginx-migration-tool_linux_$TARGETARCH${TARGETVARIANT:+_7}*/nginx-migration-tool /usr/bin/