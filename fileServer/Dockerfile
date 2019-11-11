FROM golang:latest

ADD . /go/src/github.com/KillianDavitt/CS4032-DistributedFileSystem/fileServer

RUN go install github.com/KillianDavitt/CS4032-DistributedFileSystem/fileServer

ENTRYPOINT /go/bin/fileServer docker

EXPOSE 8080
