FROM golang:1.15-alpine AS build

# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/gagarin/
WORKDIR /go/src/gagarin

# Install library dependencies
RUN dep ensure -vendor-only
RUN make install

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY . /go/src/gagarin/
RUN make build-prod

# This results in a single layer image
FROM scratch
COPY --from=build /bin/gagarin /bin/gagarin
ENTRYPOINT ["/bin/gagarin"]

RUN make run-server

EXPOSE 9444 9245