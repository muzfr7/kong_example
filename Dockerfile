# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy all the files in current directory to container's workspace.
ADD . /go/src/github.com/muzfr7/kong_example

# Install Gorilla Mux & other dependencies
RUN go get github.com/gorilla/mux

# Install our package
RUN go install github.com/muzfr7/kong_example

# Run the kong_example when the container starts.
ENTRYPOINT /go/bin/kong_example
