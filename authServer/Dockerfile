# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)

RUN openssl genrsa -out ca.key.pem 1024 &&  openssl req -subj '/CN=www.mydom.com/O=My Company Name LTD./C=US'\
      -key ca.key.pem \
      -new -x509 -days 7300 -sha256 -extensions v3_ca \
      -out ca.cert.pem


ADD . /go/src/github.com/KillianDavitt/CS4032-DistributedFileSystem/authServer

RUN go install github.com/KillianDavitt/CS4032-DistributedFileSystem/authServer

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/authServer

# Document that the service listens on port 8080.
EXPOSE 8080
