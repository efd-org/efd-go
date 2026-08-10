FROM golang:1.22-alpine AS builder

# Base build context
WORKDIR /src

# Only copy module files first for efficient caching
COPY go.mod ./
# If a go.sum exists it will be copied by the next step; download modules
RUN go mod download

# Copy the rest of the source
COPY . .

# Build a static linux binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -trimpath -ldflags="-s -w" -o /app .

# Final minimal runtime
FROM gcr.io/distroless/static:nonroot

# Copy binary from builder
COPY --from=builder /app /app

EXPOSE 8080

USER nonroot

ENTRYPOINT ["/app"]