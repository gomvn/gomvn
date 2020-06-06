################################
# STEP 1 build executable binary
################################
FROM golang:1.14-alpine AS builder

RUN apk update && apk add --no-cache git make build-base

WORKDIR /build
COPY internal internal
COPY main.go main.go
COPY go.mod go.mod
COPY go.sum go.sum

# Build the binary
RUN go install
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o output/app .

############################
# STEP 2 build a small image
############################
FROM alpine

WORKDIR /app
RUN apk update && apk add --no-cache sqlite

# Copy executable
COPY --from=builder /build/output/app /app/app
COPY config.yml /app/config.yml
COPY views /app/views

# Run the binary
ENTRYPOINT ["/app/app"]
