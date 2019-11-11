FROM golang:latest
RUN openssl genrsa -out ca.key.pem 4096 &&  openssl req -subj '/CN=www.mydom.com/O=My Company Name LTD./C=US'\
      -key ca.key.pem \
      -new -x509 -days 7300 -sha256 -extensions v3_ca \
      -out ca.cert.pem


ADD . /go/src/github.com/KillianDavitt/CS4032-DistributedFileSystem/lockServer

RUN go install github.com/KillianDavitt/CS4032-DistributedFileSystem/lockServer

ENTRYPOINT /go/bin/lockServer

EXPOSE 8080
