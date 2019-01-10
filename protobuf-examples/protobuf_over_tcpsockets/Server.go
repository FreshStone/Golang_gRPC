package main

import (
  "fmt"
  "net"
  "io"
  "bytes"
  "encoding/binary"
  "github.com/golang/protobuf/proto"
  identity "./proto_files"
)

func main(){
  startServer()
}

func startServer(){
  listner, err := net.Listen("tcp", ":2000")
  if err != nil{
    fmt.Errorf("couldn't start server: %v\n", err)
  }
  fmt.Println("Server Started, Listening on port 2000")
  ch := make(chan []byte, 10)
  go read_msgs(ch)
  for{
    c, err := listner.Accept()
    if err != nil{
      fmt.Errorf("couldn't connect:%v\n", err)
      continue
    }
    fmt.Println("New Client Connected")
    go handleconnection(c, ch)
  }
}

func handleconnection(c net.Conn, ch chan []byte){
    len_arr := make([]byte, 8)
    msg_buf := new(bytes.Buffer)
    var length int64
    for{
      _, err := c.Read(len_arr)
      if err != nil{
        fmt.Errorf("couldn't read msg len from connection: %v\nclosing connection\n", err)
        fmt.Println("closing connection")
        c.Close()
        return
      }
      length, _ = binary.Varint(len_arr)
      _, err = io.CopyN(msg_buf, c, length)
      if err != nil{
        fmt.Errorf("couldn't read msg from connection: %v\nclosing connection\n", err)
	fmt.Println("Closing connection....")
        c.Close()
        return
      }
      ch <- msg_buf.Bytes()
      msg_buf.Reset()
    }
}

func read_msgs(ch chan []byte){
  var client_no int8
  for {
    msg := <-ch
    binary.Read(bytes.NewReader(msg[:1]), binary.LittleEndian, &client_no)
    msg = msg[1:]
    if client_no == 1{
      id := identity.Individual{}
      proto.Unmarshal(msg, &id)
      fmt.Println(id.String())
    }else if client_no == 2{
      id := identity.Enterprise{}
      proto.Unmarshal(msg, &id)
      fmt.Println(id.String())
    }
  }
}
