# Use the offical Go image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang

FROM golang:1.23.0 AS builder
WORKDIR /app


RUN go mod init omakase-rooms-go-backend

# Copy local code to the container image.
COPY . .


# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -o /omakase-rooms-go-backend ./cmd

FROM gcr.io/distroless/base-debian11

WORKDIR /


COPY --from=builder /omakase-rooms-go-backend /omakase-rooms-go-backend

# Run
USER nonroot:nonroot
ENTRYPOINT ["/omakase-rooms-go-backend"]