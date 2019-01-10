package main

import (
	"fmt"
	"context"
	"net"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "../proto"
)

func main(){
	l, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(l); err != nil{
		fmt.Errorf("couldn't serve: %v", err)
	}
}

type server struct{}

func (srv *server) Echo(ctc context.Context, req *pb.Request)(*pb.Response, error){
	log.Printf("received msg: %v", req.Txt)
	hostname, _ := os.Hostname()
	return &pb.Response{Txt : req.Txt +" --from"+ hostname}, nil
}


