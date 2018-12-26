package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "flag"
  "encoding/binary"
  "bytes"

  "github.com/golang/protobuf/proto"
  pb "./proto_files"
)

func main(){
  flag.Parse()
  if flag.NArg() != 1 {
    fmt.Fprintln(os.Stderr, "Incorrect Usage\nUsage: go run HelloWorld.go method_name\n")
    os.Exit(1)
  }
  if flag.Arg(0) == "add"{
    if status := addPerson(); status{
      fmt.Println("succesfullty added person")
    }
  }else {
    listPeople()
  }
}

func addPerson()(status bool){
  person := &pb.Person{
    Id: 2,
    Name: "ashis",
    Nationality: pb.Person_PAKISTAN,
  }
  b, _ := proto.Marshal(person)
  f, err := os.OpenFile("db.pb", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
  if err != nil {
    fmt.Println(err)
    return false
  }
  binary.Write(f, binary.LittleEndian, int64(len(b))) // 8 bytes written to f
  if _, err := f.Write(b); err != nil{
    fmt.Println(err)
    return false
  }
  if err := f.Close(); err != nil{
    fmt.Println(err)
    return false
  }
  return true
}

func listPeople(){
  b, err := ioutil.ReadFile("db.pb")
  if err != nil{
    fmt.Println(err)
  }
  var length int64
  for {
    if len(b) == 0{
      return
    }else if len(b) < 8 {
        fmt.Println("garbage bytes remaining")
        return
    }
    binary.Read(bytes.NewReader(b[:8]), binary.LittleEndian, &length)
    b = b[8:]
    person := &pb.Person{} //var person *pb.Person; nil pointer err
    if err := proto.Unmarshal(b[:length], person); err != nil{
      fmt.Println(err)
    }
    b = b[length:]
    fmt.Println(person.String())
  }
}
