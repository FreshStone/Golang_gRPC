package main

import (
	"fmt"
        "context"

        "google.golang.org/grpc"
	"google.golang.org/grpc/status"
        pb "../proto"
)

func main(){
	text := "hello"
	cc, err := grpc.Dial(":50051", grpc.WithInsecure()) //could also use nginx as reverse proxy server
	if err != nil{
		fmt.Errorf("couldn't call server: %v", err)
	}
	defer cc.Close()
	c := pb.NewEchoClient(cc)
	callecho(c, text)
	t := pb.NewTimeClient(cc)
	callgettime(t)
}

func callecho(c pb.EchoClient, text string) {
	res, err := c.Echo(context.Background(), &pb.Request{Txt: text})
	if err != nil{
		fmt.Println(status.Code(err))
		return
	}
	fmt.Println(res.Txt)
}

func callgettime(t pb.TimeClient){
	res, err := t.GetTime(context.Background(), &pb.Void{})
	if err != nil{
		fmt.Println(status.Code(err))
		return
	}
	fmt.Printf("%v --from %v\n", res.GetTime(), res.Hostname)
}
