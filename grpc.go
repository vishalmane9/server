package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"server/protocol"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type server struct {
	protocol.UnimplementedPrintServiceServer
}

func (s *server) PrintPerson(ctx context.Context, in *protocol.Person) (*protocol.Person, error) {
	return &protocol.Person{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Year: in.GetYear(),
	}, nil
}

func (s *server) mustEmbedUnimplementedPrintServiceServer() {}

func runGrpcServer() {
	fmt.Println("Starting the grpc server........................")
	listener, err := net.Listen("tcp", "127.0.0.1:8085")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("GRPC Server is listerning.......................")

	newServer := grpc.NewServer()
	protocol.RegisterPrintServiceServer(newServer, &server{})

	err = newServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// func getGrpcClient() protocol.PrintServiceClient {
// 	conn, err := grpc.Dial("127.0.0.1:8085", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// defer conn.Close()
// 	client := protocol.NewPrintServiceClient(conn)
// 	return client
// }

func runGrpcClient() {
	conn, err := grpc.Dial("127.0.0.1:8085", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()
	client := protocol.NewPrintServiceClient(conn)

	fmt.Println("In the grpc client")

	currentPerson := &protocol.Person{
		Name: "harshit",
		Age:  26,
		Year: 1995,
	}
	fmt.Println("currentPerson : ", currentPerson)

	ctx := context.Background()
	personProtocol, err := client.PrintPerson(ctx, currentPerson)
	fmt.Println("personProtocol : ", personProtocol)
	data, err := proto.Marshal(personProtocol)
	if err != nil {
		fmt.Println("Proto Marshalling error : ", err)
	}
	fmt.Println("Proto marshalled data : ", data)

	fmt.Println("grpc client persons protocol : ", string(data))
}
