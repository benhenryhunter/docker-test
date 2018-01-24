# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang


ENV GOPATH /Users/bendickman/go
ENV GOROOT /usr/local/go

ENV PATH /usr/local/go/bin:/go/bin:$PATH
# Copy the local package files to the container's workspace.
ADD . /Users/bendickman/go/src/github.com/dickmanben/docker-test/
ADD . /Users/bendickman/go/src/github.com/gorilla/mux/

CMD go get github.com/gorilla/mux
# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)


RUN go install github.com/dickmanben/docker-test

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/docker-test

# Document that the service listens on port 8080.
EXPOSE 8080

