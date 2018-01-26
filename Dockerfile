# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container's workspace.
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get github.com/davecgh/go-spew/spew
RUN go build -o main .
CMD ["/app/main"]

# RUN cd /go/src; go get github.com/davecgh/go-spew/spew; 
# RUN cd /go/src; go build github.com/dickmanben/docker-test
# RUN go install github.com/dickmanben/docker-test

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)


# RUN go build -o /go/src/github.com/dickmanben/docker-test && go install github.com/dickmanben/docker-test

# ENTRYPOINT main

# Document that the service listens on port 8080.
EXPOSE 8080

