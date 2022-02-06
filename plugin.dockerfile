FROM golang as builder

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0

# Default to latest if no version is set
ARG VERSION=latest

RUN go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@$VERSION
# Note, the images must be built for amd64. If the host machine architecture is not amd64
# you need to cross-compile the binary and move it into /go/bin.
RUN bash -c 'find /go/bin/${GOOS}_${GOARCH}/ -mindepth 1 -maxdepth 1 -exec mv {} /go/bin \;'

FROM scratch

# Default to latest if no version is set
ARG VERSION=latest

# Runtime dependencies
LABEL "build.buf.plugins.runtime_library_versions.0.name"="github.com/cosmos/cosmos-proto"
LABEL "build.buf.plugins.runtime_library_versions.0.version"="${VERSION}"
LABEL "build.buf.plugins.runtime_library_versions.1.name"="google.golang.org/protobuf"
LABEL "build.buf.plugins.runtime_library_versions.1.version"="v1.27.1"

COPY --from=builder /go/bin /

ENTRYPOINT ["/protoc-gen-go-pulsar"]
