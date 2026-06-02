# --- Phase 1: Build the binary ---
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy dependency files first to leverage Docker caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Compile the binary (statically linked for performance and portability)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# --- Phase 2: Run the binary ---
FROM alpine:latest

WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
