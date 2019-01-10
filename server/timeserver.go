package main

import (
        "fmt"
        "context"
        "net"
        "log"
	"time"
	"os"

        "google.golang.org/grpc"
        "google.golang.org/grpc/reflection"
        pb "../proto"
)

func main(){
	l, _ := net.Listen("tcp", ":50051")
        s := grpc.NewServer()
        pb.RegisterTimeServer(s, &server{})
        reflection.Register(s)
        if err := s.Serve(l); err != nil{
                fmt.Errorf("couldn't serve: %v", err)
        }
}

type server struct{}

func (srv *server) GetTime(context.Context, *pb.Void)(*pb.ServerTime, error){
	hostname, _ := os.Hostname()
	ct := &pb.ServerTime{Time: &pb.Timestamp{Seconds: time.Now().Unix()}, Hostname: hostname}
	log.Printf("sent: %v", ct.GetTime())
	return ct, nil
}
