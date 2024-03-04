# Use the official Golang image to create a build artifact.
FROM golang:1.20-alpine as builder

# Copy local code to the container image.
WORKDIR /go/app-oteller
COPY . .

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app main.go

# Use a Docker multi-stage build to create a lean production image.
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/app-oteller/app /app

# Run the service binary.
CMD ["/app"]
