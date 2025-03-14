# Stage 1: Build the Go binary
FROM golang:1.24.1-bookworm AS builder
WORKDIR /app
# COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o test-cicd-go .

# Stage 2: Create a minimal runtime image
FROM debian:bookworm-slim
COPY --from=builder /app/test-cicd-go /usr/local/bin/app
RUN chmod +x /usr/local/bin/app
ENTRYPOINT ["/usr/local/bin/app"]
