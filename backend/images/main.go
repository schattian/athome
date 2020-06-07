package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/images/imageconf"
	"github.com/athomecomar/athome/backend/images/pb/pbimages"
	"github.com/athomecomar/athome/backend/images/server"
	"google.golang.org/grpc"
)

func main() {
	port := imageconf.GetPORT()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("listening on port " + port)

	s := grpc.NewServer()
	store := imageconf.GetSTORE()
	pbimages.RegisterImagesServer(s, &server.Server{Store: store})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
