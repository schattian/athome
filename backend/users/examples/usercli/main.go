package main

import (
	"context"
	"log"
	"time"

	"github.com/athomecomar/athome/users/pb/pbuser"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9991"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbuser.NewUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// signUp, err := c.SignUp(ctx, &pbuser.SignUpRequest{
	// 	Email: "foo@bar.com", Password: "foobarbaz", Name: "quux", Surname: "baz", Role: string(field.ServiceProvider),
	// })
	// if err != nil {
	// 	log.Fatalf("SignUp: %v", err)
	// }
	// log.Println(signUp)

	signIn, err := c.SignIn(ctx, &pbuser.SignInRequest{Email: "foo@bar.com", Password: "foobarbaz"})
	if err != nil {
		log.Fatalf("SignIn: %v", err)
	}
	log.Println(signIn)
}
