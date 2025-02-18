FROM golang:1.23-alpine AS builder

RUN mkdir -p /hicon
WORKDIR /hicon

COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy data to working dir
COPY . .

# Build
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" \
  CGO_ENABLED=0 \
  GOOS=linux \
  go build -ldflags="-s -w" -installsuffix cgo -o ./cmd/main ./cmd/main.go

######## Start a new stage from scratch #######
FROM alpine:3.20.3

RUN apk --no-cache add ca-certificates tzdata zip

WORKDIR /hicon

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /hicon .

# Command to run the executable
CMD ["./cmd/main"]
