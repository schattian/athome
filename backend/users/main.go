package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/users/server/configsrv"
	"github.com/athomecomar/athome/backend/users/server/signsrv"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/athome/pb/pbusers"
	"google.golang.org/grpc"
)

func main() {
	port := userconf.GetPORT()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)

	pbusers.RegisterSignServer(s, &signsrv.Server{})
	pbusers.RegisterConfigServer(s, &configsrv.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
