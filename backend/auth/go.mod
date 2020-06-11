module github.com/athomecomar/athome/backend/auth

go 1.14

require (
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/alicebob/miniredis/v2 v2.11.4 // indirect
	github.com/athomecomar/envconf v1.1.0
	github.com/athomecomar/pb v0.0.0-20200611195403-f906817bf6d2
	github.com/athomecomar/xerrors v1.2.1
	github.com/athomecomar/xtest v0.0.0-20200516201628-af9df9901326
	github.com/client9/misspell v0.3.4 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/go-redis/redis/v8 v8.0.0-beta.2
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.4.1
	github.com/google/go-jsonnet v0.16.0 // indirect
	github.com/gordonklaus/ineffassign v0.0.0-20200309095847-7953dde2c7bf // indirect
	github.com/lib/pq v1.5.2 // indirect
	github.com/mdempsky/unconvert v0.0.0-20200228143138-95ecdbfc0b5f // indirect
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120 // indirect
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
	mvdan.cc/unparam v0.0.0-20200501210554-b37ab49443f7 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1
