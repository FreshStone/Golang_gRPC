FROM golang:latest
RUN go get -u google.golang.org/grpc
RUN mkdir /timeserver/
ADD . /timeserver/
WORKDIR /timeserver

EXPOSE 50051
CMD ["go", "run", "timeserver.go"]

