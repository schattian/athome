package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/checkout/server/srvpurchases"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/athome/pb/pbconf"
	"google.golang.org/grpc"
)

func main() {
	svc := pbconf.Checkout
	port := svc.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)
	pbcheckout.RegisterPurchasesServer(s, &srvpurchases.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
