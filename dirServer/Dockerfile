FROM golang


ADD . /go/src/github.com/KillianDavitt/CS4032-DistributedFileSystem/dirServer

#RUN bash /go/src/github.com/KillianDavitt/CS4032-DistributedFilesystem/dirServer/certs_gen.sh
RUN go install github.com/KillianDavitt/CS4032-DistributedFileSystem/dirServer

ENTRYPOINT /go/bin/dirServer

EXPOSE 8080
