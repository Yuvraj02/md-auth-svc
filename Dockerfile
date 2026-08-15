# Build from this service repo root:
#   docker build -t marketing-digest-auth .
#
# Same image runs the gRPC server (Deployment) and Atlas migrate (Job).

FROM golang:1.25-alpine AS build

WORKDIR /src

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
COPY pkg ./pkg
COPY . .

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux \
    go build -trimpath -ldflags="-s -w" \
    -o /out/auth-service ./cmd/server


# Final image contains both auth-service and Atlas.
FROM arigaio/atlas:latest-alpine

COPY --from=build /out/auth-service /auth-service
COPY --from=build /src/migrations /migrations

USER nobody

EXPOSE 50051

ENTRYPOINT ["/auth-service"]