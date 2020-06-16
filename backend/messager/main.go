package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/messager/server"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbmessager"
	"google.golang.org/grpc"
)

func main() {
	svc := pbconf.Messager
	port := svc.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)

	pbmessager.RegisterMessagesServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
