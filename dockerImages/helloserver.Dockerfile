FROM golang:latest
RUN go get -u google.golang.org/grpc
RUN mkdir /helloserver
ADD . /helloserver/
WORKDIR /helloserver

EXPOSE 50051
CMD ["go", "run", "helloserver.go"]
