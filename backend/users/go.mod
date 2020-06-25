module github.com/athomecomar/athome/backend/users

go 1.14

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/athomecomar/athome/pb v0.0.0-20200625203305-ffb893e2ddc8
	github.com/athomecomar/envconf v1.2.0
	github.com/athomecomar/semantic v0.2.1
	github.com/athomecomar/storeql v1.5.4
	github.com/athomecomar/xerrors v1.2.1
	github.com/athomecomar/xtest v0.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/golang/mock v1.4.0
	github.com/google/go-cmp v0.4.1
	github.com/jmoiron/sqlx v1.2.1-0.20191203222853-2ba0fc60eb4a
	github.com/lib/pq v1.7.0
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9
	golang.org/x/sys v0.0.0-20200610111108-226ff32320da // indirect
	google.golang.org/genproto v0.0.0-20200612171551-7676ae05be11 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1
