FROM golang:1.21-alpine AS builder

# Install necessary tools
RUN apk add --no-cache git build-base

WORKDIR /app

# Copy dependency files
COPY go.mod .
RUN touch go.sum
RUN go mod download
RUN go mod tidy -v

# Copy source code
COPY . .
RUN go mod tidy -v

# Build application
RUN CGO_ENABLED=0 GOOS=linux go build -v -o url-to-markdown .

# Create runtime image
FROM alpine:latest

WORKDIR /app

# Install necessary system packages
RUN apk add --no-cache ca-certificates poppler-utils

# Copy binary
COPY --from=builder /app/url-to-markdown .

# Create directories for static files and templates
RUN mkdir -p /app/static/css
RUN mkdir -p /app/static/js
RUN mkdir -p /app/templates

# Copy static files and templates - simplified to avoid duplications
COPY static/ /app/static/
COPY templates/ /app/templates/

# Create directory for temporary files
RUN mkdir -p /tmp/pdf-uploads && chmod 777 /tmp/pdf-uploads

# Expose port
EXPOSE 8080

# Run application
CMD ["./url-to-markdown"]