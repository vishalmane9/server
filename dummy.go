package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net"
// 	"server/protocol"

// 	"google.golang.org/protobuf/proto"
// )

// func runserver() {
// 	listener, err := net.Listen("tcp", "127.0.0.1:8082")
// 	if err != nil {
// 		fmt.Println("Listening error : ", err)
// 	}

// 	//grpc client
// 	// conngrpc, err := grpc.Dial("127.0.0.1:8086", grpc.WithInsecure())
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }

// 	// defer conngrpc.Close()
// 	// client := protocol.NewPrintServiceClient(conngrpc)
// 	// fmt.Println("In the grpc client")

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("listener accept error : ", err)
// 		}

// 		go func(c net.Conn) {
// 			defer c.Close()
// 			data, err := ioutil.ReadAll(conn)
// 			if err != nil {
// 				fmt.Println("reading from connection error : ", err)
// 			}

// 			person := &protocol.Person{}
// 			proto.Unmarshal(data, person)
// 			fmt.Println(person)

// 			fmt.Println("person name : ", person.GetName())
// 			fmt.Println("person age : ", person.GetAge())
// 			fmt.Println("person year : ", person.GetYear())

// 			// go func() {
// 			// 	person, err = client.PrintPerson(context.Background(), person)
// 			// 	if err != nil {
// 			// 		log.Fatalln(err)
// 			// 	}

// 			// 	fmt.Println("grpc client persons : ", person)
// 			// }()

// 		}(conn)
// 	}

// }

// // func sendData(data []byte) {
// // 	conn, err := net.Dial("tcp", "127.0.0.1:8085")
// // 	if err != nil {
// // 		fmt.Println("Connection error : ", err)
// // 	}

// // 	defer conn.Close()
// // 	conn.Write(data)
// // }
