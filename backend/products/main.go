package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/products/server/srvcreator"
	"github.com/athomecomar/athome/backend/products/server/srvviewer"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbproducts"
	"google.golang.org/grpc"
)

func main() {
	svc := pbconf.Products
	port := svc.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)

	pbproducts.RegisterViewerServer(s, &srvviewer.Server{})
	pbproducts.RegisterCreatorServer(s, &srvcreator.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
