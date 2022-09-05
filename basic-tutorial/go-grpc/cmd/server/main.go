package main

import (
	"context"
	"go-grpc/pb"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedSendMessageServer
}

func (service *Server) RequestMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Print("mensagem recebida: ", req.GetMessage())

	response := &pb.Response{
		Status: 1,
	}

	return response, nil
}
func (service *Server) mustEmbedUnimplementedSendMessageServer() {}

func main() {
	grpcServer := grpc.NewServer()

	// registrar server
	pb.RegisterSendMessageServer(grpcServer, &Server{})

	port := ":5000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
