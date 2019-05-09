FROM golang:1.12.5 AS build-env

ENV GO111MODULE=off
ENV GO15VENDOREXPERIMENT=1
ENV BUILDPATH=github.com/lattecake/blog
RUN mkdir -p /go/src/${BUILDPATH}
COPY ./ /go/src/${BUILDPATH}
RUN cd /go/src/${BUILDPATH} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v

FROM alpine:latest

COPY --from=build-env /go/bin/blog /go/bin/blog
COPY ./conf /go/bin/conf
COPY ./static /go/bin/static
WORKDIR /go/bin/
CMD ["/go/bin/blog"]