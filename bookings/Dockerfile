# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.6

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# Copy the local package files to the container workspace
ADD . /go/src/github.com/virtyaluk/go-simple-microservices-app/bookings

# Setting up working directory
WORKDIR /go/src/github.com/virtyaluk/go-simple-microservices-app/bookings

# Build the bookings command inside the container
RUN go install github.com/virtyaluk/go-simple-microservices-app/bookings

# Run the bookings microservice when the container starts
ENTRYPOINT /go/bin/movies

# Service listens on port 8080.
EXPOSE 8080
