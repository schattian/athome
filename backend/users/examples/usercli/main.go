package main

import (
	"context"
	"log"
	"time"

	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athomecomar/athome/pb/go/pbuser"
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

	// signUp(ctx, c)
	// signIn(ctx, c)
	alreadyExists(ctx, c)
	defer cancel()
}

func alreadyExists(ctx context.Context, c pbuser.UserClient) {
	alreadyExists, err := c.AlreadyExists(ctx, &pbuser.AlreadyExistsRequest{Email: "foo@bar.com", Role: "service-provider"})
	if err != nil {
		log.Fatalf("AlreadyExists: %v", err)
	}
	log.Println(alreadyExists)
}

func signIn(ctx context.Context, c pbuser.UserClient) {
	signIn, err := c.SignIn(ctx, &pbuser.SignInRequest{Email: "foo@bar.com", Password: "foobarbaz"})
	if err != nil {
		log.Fatalf("SignIn: %v", err)
	}
	log.Println(signIn)
}

func signUp(ctx context.Context, c pbuser.UserClient) {
	signUp, err := c.SignUp(ctx, &pbuser.SignUpRequest{
		Email: "foo@bar.com", Password: "foobarbaz", Name: "quux", Surname: "baz", Role: string(field.ServiceProvider),
	})
	if err != nil {
		log.Fatalf("SignUp: %v", err)
	}
	log.Println(signUp)
}
