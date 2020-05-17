module github.com/athomecomar/athome/backend/users

go 1.14

require (
	github.com/athomecomar/envconf v1.0.0
	github.com/athomecomar/storeql v1.1.0
	github.com/athomecomar/xerrors v1.1.1-0.20200517202822-060c4da30e58
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/golang/protobuf v1.4.2
	github.com/jmoiron/sqlx v1.2.1-0.20191203222853-2ba0fc60eb4a
	github.com/lib/pq v1.5.2
	github.com/omeid/pgerror v0.0.0-20200121005125-8254ae5f5ba0
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120 // indirect
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	google.golang.org/genproto v0.0.0-20200515170657-fc4c6c6a6587 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1
