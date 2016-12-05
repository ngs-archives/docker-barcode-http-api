FROM golang:1.7

MAINTAINER Atsushi Nagase<a@ngs.io>
RUN apt-get update && apt-get -y install libzbar-dev && apt-get clean
RUN go get github.com/PeterCxy/gozbar

COPY server.go .
RUN go get -v -t -d ./...
RUN go build -o /usr/bin/server server.go

CMD /usr/bin/server
