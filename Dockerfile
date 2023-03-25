# syntax=docker/dockerfile:1
FROM golang:1.20-alpine AS build

# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
RUN apk add --no-cache git

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
COPY go.mod go.sum *.go /go/src/project/
WORKDIR /go/src/project/
# Install library dependencies
RUN go install

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY . /go/src/project/
RUN go build -o /bin/project

ENTRYPOINT ["/bin/project"]
CMD ["--help"]