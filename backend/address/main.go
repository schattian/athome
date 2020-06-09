package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/address/addressconf"
	"github.com/athomecomar/athome/backend/address/pb/pbaddress"
	"github.com/athomecomar/athome/backend/address/server"
	"google.golang.org/grpc"
)

func main() {
	port := addressconf.GetPORT()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)

	pbaddress.RegisterAddressServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
