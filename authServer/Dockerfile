FROM golang

ADD . /go/src/github.com/KillianDavitt/CS4032-DistributedFileSystem/authServer

RUN go install github.com/KillianDavitt/CS4032-DistributedFileSystem/authServer
RUN cd src/github.com/KillianDavitt/CS4032-DistributedFileSystem/authServer
ENTRYPOINT /go/bin/authServer

EXPOSE 8080
