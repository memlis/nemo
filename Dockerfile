FROM golang:1.6.3-alpine

COPY . /go/src/github.com/memlis/boat
WORKDIR /go/src/github.com/memlis/boat
RUN go build -v  && install -v boat /usr/local/bin
ENTRYPOINT ["/usr/local/bin/boat"]

