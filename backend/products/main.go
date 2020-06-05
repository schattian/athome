package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/products/pb/pbproducts"
	"github.com/athomecomar/athome/backend/products/productconf"
	"github.com/athomecomar/athome/backend/products/server/srvcreator"
	"google.golang.org/grpc"
)

func main() {
	port := productconf.GetPORT()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)

	pbproducts.RegisterCreatorServer(s, &srvcreator.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
