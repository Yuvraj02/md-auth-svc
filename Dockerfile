# Build from workspace root (parent of backend/ and protos/):
#   docker build -f backend/services/auth-service/Dockerfile -t marketing-digest-auth-service .
FROM golang:1.25-alpine AS build
WORKDIR /src
RUN apk add --no-cache ca-certificates git

COPY protos ./protos
COPY backend/pkg ./backend/pkg
COPY backend/services/auth-service ./backend/services/auth-service

WORKDIR /src/backend/services/auth-service
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/auth-service ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/auth-service /auth-service
USER nonroot:nonroot
EXPOSE 50051
ENTRYPOINT ["/auth-service"]
