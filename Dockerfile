# syntax=docker/dockerfile:1@sha256:865e5dd094beca432e8c0a1d5e1c465db5f998dca4e439981029b3b81fb39ed5

# Create .netrc file for private go module
# FROM bufbuild/buf:1.25.1 as buf

# ARG BUF_USERNAME ""

# SHELL ["/bin/ash", "-o", "pipefail", "-c"]
# RUN --mount=type=secret,id=BUF_TOKEN \
#   buf registry login --username=$BUF_USERNAME --token-stdin < /run/secrets/BUF_TOKEN

# Build go binary
FROM golang:1.22-alpine@sha256:1a478681b671001b7f029f94b5016aed984a23ad99c707f6a0ab6563860ae2f3 AS build

WORKDIR /go/src/app

# COPY --from=buf /root/.netrc /root/.netrc
# ENV GOPRIVATE=buf.build/gen/go

COPY go.mod go.sum /
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" ./cmd/api-server

# Run
FROM gcr.io/distroless/static:nonroot@sha256:dcd3f1f09adef5689088c9c4d96a8d98c889d8281d3946145074f89eafe7e1af

COPY --from=build /go/bin/app .

EXPOSE 50051
ENTRYPOINT ["/home/nonroot/app"]