module github.com/athomecomar/athome/users

go 1.14

require (
	github.com/athomecomar/envconf v1.0.0
	github.com/athomecomar/storeql v0.0.0-20200516200600-e6498c6ccc32
	github.com/athomecomar/xerrors v1.0.1
	github.com/golang-migrate/migrate v3.5.4+incompatible
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/golang/protobuf v1.4.2
	github.com/howeyc/fsnotify v0.9.0 // indirect
	github.com/jmoiron/sqlx v1.2.1-0.20191203222853-2ba0fc60eb4a
	github.com/lib/pq v1.5.2
	github.com/pilu/config v0.0.0-20131214182432-3eb99e6c0b9a // indirect
	github.com/pilu/fresh v0.0.0-20190826141211-0fa698148017 // indirect
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120 // indirect
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	google.golang.org/genproto v0.0.0-20200515170657-fc4c6c6a6587 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1
