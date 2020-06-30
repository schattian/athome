package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/agreement/agreementconf"
	"github.com/athomecomar/athome/backend/agreement/server"
	"github.com/athomecomar/athome/pb/pbagreement"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/go-redis/redis/v8"

	"google.golang.org/grpc"
)

func main() {
	svc := pbconf.Auth
	port := svc.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)
	pbagreement.RegisterAgreementServer(s, &server.Server{Redis: redisCli()})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func redisCli() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     agreementconf.GetDATABASE_HOST() + agreementconf.GetDATABASE_PORT(),
		Password: agreementconf.GetDATABASE_PASSWORD(),
		DB:       agreementconf.GetDATABASE_NUMBER(),
	})
}
