package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/identifier/server"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbidentifier"
	"google.golang.org/grpc"
)

func main() {
	svc := pbconf.Identifier
	port := svc.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)
	pbidentifier.RegisterIdentifierServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
