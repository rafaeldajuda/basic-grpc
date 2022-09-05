package main

import (
	"context"
	"go-grpc/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// conexao
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	// resgistar client
	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "Hello World",
	}

	res, err := client.RequestMessage(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res.GetStatus())

}
