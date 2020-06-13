package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/address/server"
	"github.com/athomecomar/athome/pb/pbaddress"
	"github.com/athomecomar/athome/pb/pbconf"
	"google.golang.org/grpc"
)

func main() {
	svc := pbconf.Addresses
	port := svc.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)

	pbaddress.RegisterAddressesServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
