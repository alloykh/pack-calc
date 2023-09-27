FROM golang:1.18-buster AS build

COPY . /go/src/github.com/alloykh/pack-calc

WORKDIR /go/src/github.com/alloykh/pack-calc

#
# Build to the service directory
#

RUN SERVER_BINARY=/bin/app make build

FROM golang:1.18-buster
#
# Copy out from the build container
#
COPY --from=build bin/app bin/app
ENTRYPOINT ["bin/app"]

