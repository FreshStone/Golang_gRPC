package main

import (
	"fmt"
	"net"
	"flag"
	"strings"
	"strconv"
	"os"
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	identity "./proto_files"
)

func main(){
	flag.Parse()
	if flag.NArg() != 1{
		fmt.Fprintln(os.Stderr, "missing subcommand 'mode': individual or enterprise")
		os.Exit(1)
	}

	if flag.Arg(0) == "individual" || flag.Arg(0) == "enterprise"{
		startClient(flag.Arg(0))
	}else{
		fmt.Fprintln(os.Stderr, "wrong subcommand" )
                os.Exit(1)
	}
}

func startClient(mode string){
	conn, err := net.Dial("tcp", ":2000")
	if err != nil{
		fmt.Fprintf(os.Stderr, "couldn't connect to server: %v\n", err)
		os.Exit(1)
        }
	fmt.Println("Connected to the server")
	rd := bufio.NewReader(os.Stdin)
	buf := new(bytes.Buffer)
	len_arr := make([]byte, 8)
	if mode == "individual"{
		for {
			buf.Reset()
			person := &identity.Individual{}
			fmt.Print("Enter person ID number: ")
			if _, err := fmt.Fscanf(rd, "%d\n", &person.Id); err != nil {
				fmt.Fprintf(os.Stderr, "couldn't parse id: %v\n", err)
			}
			fmt.Print("Enter name: ")
			name, err := rd.ReadString('\n')
			if err != nil{
				fmt.Fprintf(os.Stderr, "couldn't parse name: %v\n", err)
			}
			person.Name = strings.TrimSpace(name)
			fmt.Print("Enter ph number: ")
			ph, err := rd.ReadString('\n')
			if err != nil{
				fmt.Fprintf(os.Stderr, "couldn't parse ph number: %v\n", err)
			}
			person.Ph, _ = strconv.ParseInt(strings.TrimSpace(ph), 10, 64)
			fmt.Print("Enter house number for address info (or leave blank): ")
			house_no, err := rd.ReadString('\n')
			if err != nil{
				fmt.Fprintf(os.Stderr, "couldn't parse house number: %v\n", err)
			}
			if strings.TrimSpace(house_no) != ""{
				add := &identity.Individual_Address{}
				add.Housenum, _  = strconv.ParseInt(strings.TrimSpace(house_no), 10, 64)
				fmt.Print("Enter street name for address info: ")
				street, err := rd.ReadString('\n')
				if err != nil{
					fmt.Fprintf(os.Stderr, "couldn't parse street name: %v\n", err)
				}
				add.Street = strings.TrimSpace(street)
				fmt.Print("Enter city for address info: ")
				city, err := rd.ReadString('\n')
				if err != nil{
					fmt.Fprintf(os.Stderr, "couldn't parse city: %v\n", err)
				}
				add.City = strings.TrimSpace(city)
				person.Add = add
			}
			msg, err := proto.Marshal(person)
			if err != nil{
				fmt.Fprintf(os.Stderr, "couldn't marshal person struct: %v\n", err)
			}
			binary.PutVarint(len_arr, int64(len(msg)+1))   // buffer type from protobuf/proto package could be used for encoding int and marhsaling protobufs msgs 
			_, err = conn.Write(len_arr)
			if err != nil{
                                fmt.Fprintf(os.Stderr, "couldn't write length to the connection: %v\nclosing connection..\n", err)
                                conn.Close()
                                os.Exit(1)
                        }
			binary.Write(buf, binary.LittleEndian, int8(1))
			binary.Write(buf, binary.LittleEndian, msg)
			_, err = conn.Write(buf.Bytes())
			if err != nil{
				fmt.Fprintf(os.Stderr, "couldn't write msg to the connection: %v\nclosing connection..\n", err)
				conn.Close()
				os.Exit(1)
			}
			fmt.Print("like to add another individual[Y/N]: ")
			repeat, _ := rd.ReadString('\n')
			if strings.ToLower(strings.TrimSpace(repeat)) == "n"{
				break
			}
		}
	}else{
		for {
			buf.Reset()
			company := &identity.Enterprise{}
                        ceo := &identity.Individual{}
                        fmt.Print("Enter company name: ")
                        if _, err := fmt.Fscanf(rd, "%s\n", &company.Name); err != nil {
                                fmt.Fprintf(os.Stderr, "couldn't parse company name: %v\n", err)
                        }

                        fmt.Print("Enter ceo name: ")
                        name, err := rd.ReadString('\n')
                        if err != nil{
                                fmt.Fprintf(os.Stderr, "couldn't parse ceo name: %v\n", err)
                        }
                        ceo.Name = strings.TrimSpace(name)
                        fmt.Print("Enter ceo id: ")
                        id, err := rd.ReadString('\n')
                        if err != nil{
                                fmt.Fprintf(os.Stderr, "couldn't parse ceo id: %v\n", err)
                        }
			id_64, _ := strconv.ParseInt(strings.TrimSpace(id), 10, 64)
			ceo.Id = int32(id_64)
			company.Ceo = ceo

                        fmt.Print("Enter Company Address/House number (or leave blank): ")
                        house_no, err := rd.ReadString('\n')
                        if err != nil{
                                fmt.Fprintf(os.Stderr, "couldn't parse house number: %v\n", err)
                        }
                        if strings.TrimSpace(house_no) != ""{
                                add := &identity.Individual_Address{}
                                add.Housenum, _  = strconv.ParseInt(strings.TrimSpace(house_no), 10, 64)
                                fmt.Print("Enter street name for address info: ")
				street, err := rd.ReadString('\n')
                                if err != nil{
                                        fmt.Fprintf(os.Stderr, "couldn't parse street name: %v\n", err)
                                }
                                add.Street = strings.TrimSpace(street)
                                fmt.Print("Enter city for address info: ")
                                city, err := rd.ReadString('\n')
                                if err != nil{
                                        fmt.Fprintf(os.Stderr, "couldn't parse city: %v\n", err)
                                }
                                add.City = strings.TrimSpace(city)
                                company.Add = add
                        }
			fmt.Print("Enter Company's work domain (tech/security/services): ")
			domain, err := rd.ReadString('\n')
                        if err != nil{
				fmt.Fprintf(os.Stderr, "couldn't parse domain: %v\n", err)
                        }
			switch strings.TrimSpace(domain){
				case "tech":
					company.D = identity.Enterprise_TECH
				case "security":
					company.D = identity.Enterprise_SECURITY
				case "services":
					company.D = identity.Enterprise_SERVICES
				default:
					fmt.Println("Unknown company domain, choosing TECH as default country domain")
			}
			msg, err := proto.Marshal(company)
                        if err != nil{
                                fmt.Fprintf(os.Stderr, "couldn't marshal company struct: %v\n", err)
                        }
                        binary.PutVarint(len_arr, int64(len(msg)+1))
                        _, err = conn.Write(len_arr)
                        if err != nil{
                                fmt.Fprintf(os.Stderr, "couldn't write length to the connection: %v\nclosing connection..\n", err)
                                conn.Close()
                                os.Exit(1)
                        }
                        binary.Write(buf, binary.LittleEndian, int8(2))
                        binary.Write(buf, binary.LittleEndian, msg)
                        _, err = conn.Write(buf.Bytes())
                        if err != nil{
                                fmt.Fprintf(os.Stderr, "couldn't write msg to the connection: %v\nclosing connection..\n", err)
                                conn.Close()
                                os.Exit(1)
                        }
                        fmt.Print("like to add another company[Y/N]: ")
                        repeat, _ := rd.ReadString('\n')
                        if strings.ToLower(strings.TrimSpace(repeat)) == "n"{
                                break
                        }
		}
	}
}
