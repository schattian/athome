module github.com/athomecomar/athome/backend/semantic

go 1.14

require (
	github.com/athomecomar/envconf v1.1.0
	github.com/athomecomar/storeql v1.4.5
	github.com/athomecomar/xerrors v1.2.1
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/golang/protobuf v1.4.2
	github.com/jcmturner/gokrb5/v8 v8.3.0 // indirect
	github.com/jmoiron/sqlx v1.2.1-0.20191203222853-2ba0fc60eb4a
	github.com/lib/pq v1.6.0
	github.com/pkg/errors v0.9.1
	go.opentelemetry.io/otel v0.6.0
	golang.org/x/crypto v0.0.0-20200602180216-279210d13fed // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	google.golang.org/genproto v0.0.0-20200515170657-fc4c6c6a6587 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1
