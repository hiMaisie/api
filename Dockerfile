ARG golang_version=1.16
ARG builder_image=golang:${golang_version}-alpine

# Common builder layer.
# Install build-time dependencies and fetch go modules.
FROM ${builder_image} AS builder
WORKDIR /build
ADD go.mod go.sum ./
RUN go mod download

ADD . .

# Builds go binary to be as small and as performant as possible.
FROM builder AS build_prod
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o api ./main.go

# Build smallest possible image for production use, without any
# debugging symbols or shell.
FROM scratch AS prod
EXPOSE 8080
COPY --from=build_prod /build/api /api
ENTRYPOINT ["/api"]

# Dev build of binary with debugging symbols
FROM builder AS build_dev
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./main.go

# Default is a dev image that has all of the handy debug tools.
FROM ${builder_image}
# TODO: Add useful debugging tools here.
EXPOSE 8080
COPY --from=build_dev /build/api /api
ENTRYPOINT ["/api"]
