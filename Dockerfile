FROM --platform=$BUILDPLATFORM golang:1.17-alpine AS build-env
WORKDIR /go/src/github.com/rusik69/kv2
COPY . ./ARG TARGETOS
ARG TARGETARCH
RUN make get
RUN make build
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build cmd/server/main.go -o /go/bin/server

FROM scratch
WORKDIR /
COPY --from=build-env /go/bin/server /usr/bin/kv2
ENTRYPOINT ["/usr/bin/kv2"]
EXPOSE 6969 6970 6971