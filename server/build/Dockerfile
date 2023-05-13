FROM golang:1.19-alpine3.16

ENV CGO_ENABLED 0

ENV PROTOC_GEN_GO_VERSION v1.28.1
ENV PROTOC_GEN_GO_GRPC_VERSION v1.2.0
ENV BUF_VERSION v1.18.0

# Install dependencies that change infrequently:
RUN set -x \
    && apk add --no-cache --update git bash curl \
    && rm -f /var/cache/apk/*

# Buf
RUN curl -sSL "https://github.com/bufbuild/buf/releases/download/${BUF_VERSION}/buf-$(uname -s)-$(uname -m)" -o "/usr/local/bin/buf" \
    && chmod +x /usr/local/bin/buf

# go protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@$PROTOC_GEN_GO_VERSION \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$PROTOC_GEN_GO_GRPC_VERSION
