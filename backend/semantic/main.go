package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/users/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/users/semanticconf"
	"github.com/athomecomar/athome/backend/users/server/srvmerchants"
	"github.com/athomecomar/athome/backend/users/server/srvproducts"
	"github.com/athomecomar/athome/backend/users/server/srvproviders"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"google.golang.org/grpc"
)

func main() {
	port := semanticconf.GetPORT()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)

	pbsemantic.RegisterMerchantsServer(s, &srvmerchants.Server{})
	pbsemantic.RegisterServiceProvidersServer(s, &srvproviders.Server{})
	pbsemantic.RegisterProductsServer(s, &srvproducts.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
