# Dockerfile for generating Go bindings with abigen
FROM golang:1.23-alpine AS builder

# Install required packages
RUN apk add --no-cache git jq bash make

# Install abigen
RUN go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# Create working directory
WORKDIR /workspace

# Copy the binding generation script
COPY generate-bindings.sh .
RUN chmod +x generate-bindings.sh

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache bash jq

# Copy abigen binary and scripts from builder
COPY --from=builder /go/bin/abigen /usr/local/bin/
COPY --from=builder /workspace/generate-bindings.sh /usr/local/bin/

# Create workspace
WORKDIR /workspace

# Set the entrypoint
ENTRYPOINT ["/usr/local/bin/generate-bindings.sh"]

# Usage:
# docker build -f abigen-with-interfaces.Dockerfile -t triggerx-abigen .
# docker run -v $(pwd):/workspace triggerx-abigen 