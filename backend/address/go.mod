module github.com/athomecomar/athome/backend/address

go 1.14

require (
	github.com/athomecomar/athome/pb v0.0.0-20200621215531-af5227e74c8b
	github.com/athomecomar/envconf v1.2.0
	github.com/athomecomar/storeql v1.5.4
	github.com/athomecomar/xerrors v1.2.1
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/jmoiron/sqlx v1.2.1-0.20191203222853-2ba0fc60eb4a
	github.com/lib/pq v1.7.0
	github.com/pkg/errors v0.9.1
	github.com/umahmood/haversine v0.0.0-20151105152445-808ab04add26
	google.golang.org/grpc v1.29.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1
