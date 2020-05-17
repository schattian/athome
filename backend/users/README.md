# users

Boilerplate to create microservice using go-grpc framework, pq driver and zap logger.

This repo shows an example with a pkg named `auth`.

## FS Structure 

- main.go: got the server implementation example.

- pb/: holds all protocol buffers related things.
- pb/*.proto: proto declarations.
- pb/* 
- pb/pb<pkg>/: holds the compiled code from the protocol buffers to go.
- pb/pbjs/: holds the compiled code from the protocol buffers to js.

- <pkg>conf/: holds all the config files.

- examples/<pkg>cli/: got the client implementation example in go.

- ./compile: compiles the .proto declaration to js and go.

## Database

Driver is the std, lib/pq, with the addition of sqlx


## Testing

It uses JSONNET for fixtures, go-cmp for diffs.
From std lib, it'll use httpstub and httptest for stubbing http reqs and go-sqlmock for mocking sql responses

Fixtures will be under testdata/<entity>.jsonnet, being entity the name of the model created on ent/*.go