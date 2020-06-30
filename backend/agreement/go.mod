module github.com/athomecomar/athome/backend/auth

go 1.14

require (
	github.com/alicebob/gopher-json v0.0.0-20180125190556-5a6b3ba71ee6 // indirect
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/athomecomar/athome/pb v0.0.0-20200626172433-a1168111e542
	github.com/athomecomar/envconf v1.2.0
	github.com/athomecomar/xerrors v1.2.1
	github.com/athomecomar/xtest v0.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v8 v8.0.0-beta.2
	github.com/gomodule/redigo v1.7.1-0.20190322064113-39e2c31b7ca3 // indirect
	github.com/google/go-cmp v0.4.1
	github.com/pkg/errors v0.9.1
	github.com/yuin/gopher-lua v0.0.0-20191220021717-ab39c6098bdb // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.24.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1
