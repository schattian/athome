module github.com/athomecomar/athome/backend/notifier

go 1.14

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/athomecomar/athome/backend/products v0.0.0-20200613144549-ef445669ca22
	github.com/athomecomar/athome/backend/services v0.0.0-20200614200957-e4a53c7c626c
	github.com/athomecomar/athome/pb v0.0.0-20200616201740-8fad0f40bbef
	github.com/athomecomar/currency v0.1.0
	github.com/athomecomar/envconf v1.2.0
	github.com/athomecomar/storeql v1.5.4
	github.com/athomecomar/xerrors v1.2.1
	github.com/athomecomar/xtest v0.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/envoyproxy/protoc-gen-validate v0.3.0
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/golang/mock v1.4.0
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.4.1
	github.com/gordonklaus/ineffassign v0.0.0-20200309095847-7953dde2c7bf // indirect
	github.com/jmoiron/sqlx v1.2.1-0.20191203222853-2ba0fc60eb4a
	github.com/lib/pq v1.7.0
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	golang.org/x/text v0.3.2
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	mvdan.cc/unparam v0.0.0-20200501210554-b37ab49443f7 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1
