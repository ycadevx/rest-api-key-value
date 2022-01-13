# syntax=docker/dockerfile:1
FROM golang:1.17-alpine as builder

WORKDIR $GOPATH/src/github/deliveryhero/source

# Copy the code into the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

RUN go mod download

# Install the package
RUN go install -v ./...

EXPOSE 8081

CMD ["source"]