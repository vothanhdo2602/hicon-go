FROM golang:1.24.1-alpine AS builder

RUN mkdir -p /hicon
WORKDIR /hicon

COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy data to working dir
COPY . .

# Build
RUN CGO_ENABLED=0 \
  GOOS=linux \
  go build -ldflags="-s -w" -installsuffix cgo -o ./cmd/main ./cmd/main.go

######## Start a new stage from scratch #######
FROM alpine:3.21.3

RUN apk --no-cache add ca-certificates tzdata zip

WORKDIR /hicon

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /hicon/cmd/main .

# Command to run the executable
CMD ["./main"]
