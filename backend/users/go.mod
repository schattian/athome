module github.com/athomecomar/athome/backend/users

go 1.14

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/alexbrainman/sspi v0.0.0-20180613141037-e580b900e9f5 // indirect
	github.com/athomecomar/athome/pb v0.0.0-20200612182330-5f1bf6cd55b8
	github.com/athomecomar/envconf v1.2.0
	github.com/athomecomar/semantic v0.2.1
	github.com/athomecomar/storeql v1.5.4
	github.com/athomecomar/xerrors v1.2.1
	github.com/athomecomar/xtest v0.0.0-20200516201628-af9df9901326
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/golang/mock v1.4.0
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.4.1
	github.com/google/go-jsonnet v0.16.0 // indirect
	github.com/jcmturner/gokrb5/v8 v8.3.0 // indirect
	github.com/jmoiron/sqlx v1.2.1-0.20191203222853-2ba0fc60eb4a
	github.com/lib/pq v1.7.0
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200610111108-226ff32320da // indirect
	google.golang.org/genproto v0.0.0-20200612171551-7676ae05be11 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
	gopkg.in/jcmturner/aescts.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/dnsutils.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/goidentity.v3 v3.0.0 // indirect
	gopkg.in/jcmturner/gokrb5.v7 v7.5.0 // indirect
	gopkg.in/jcmturner/rpc.v1 v1.1.0 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1
