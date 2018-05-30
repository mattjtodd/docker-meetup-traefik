# This Dockerfile builds an image for a client_golang example.
#
# Use as (from the root for the client_golang repository):
#    docker build -f examples/$name/Dockerfile -t prometheus/golang-example-$name .

# Builder image, where we build the example.
FROM golang:1.9-alpine AS builder

RUN apk update \
  && apk add curl git \
  && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/main

# Get the dep package definitions to allow download without compile.  This makes the code turnaround faster as this can be cached
COPY Gopkg.* ./

# Download the vendor pakcage defs in toml files only
RUN dep ensure --vendor-only

# Copy the source files, these will chnage more often and can rely on the previousy cached libs
COPY *.go ./

# Complile and link
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w'

# Final image, just contains the cross-compiled binary.
FROM scratch
COPY --from=builder /go/src/main .
ENTRYPOINT ["/main"]