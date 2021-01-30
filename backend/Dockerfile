#FROM golang:1.15-alpine AS build

# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
# RUN apk add --no-cache gcc g++ make ca-certificates

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
# COPY Gopkg.lock Gopkg.toml /go/src/backend/
#WORKDIR /go/src/backend

# Install library dependencies
#COPY . .
#RUN make install-tools

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
#RUN make build-prod

# This results in a single layer image
#FROM scratch
#COPY --from=build /bin/backend /bin/backend
#ENTRYPOINT ["/bin/backend"]

#RUN make run-server



from golang:1.15-alpine as build-env
run apk add --no-cache git make

# add Makefile /go
# RUN make install-tools

RUN mkdir /gagarin /gagarin/backend
WORKDIR /gagarin/backend

env GO111MODULE=on
env GOBIN=/go/bin
# arg GOPROXY

add go.mod go.sum ./
run go mod download

add . ./
# RUN make install-tools
# run make givemeproto
run CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

from scratch
run apk add --no-cache ca-certificates
#entrypoint 
copy --from=build-env /go/bin/gagarin/backend /go/bin/gagarin/backend
ENTRYPOINT ["/go/bin/gagarin/backend"]

EXPOSE 9444 9245